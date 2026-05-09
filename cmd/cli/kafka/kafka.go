package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:9092"
	kafkaTopic = "user_topic_test"
)

type KafkaMessage struct {
	Action string    `json:"action"`
	Info   StockInfo `json:"info"`
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func getKafkaReader(
	kafkaURL,
	topic,
	groupID string,
) *kafka.Reader {

	brokers := strings.Split(kafkaURL, ",")

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		GroupID:     groupID,
		Topic:       topic,
		StartOffset: kafka.LastOffset, // không đọc msg cũ
	})
}

func getKafkaWriter(
	kafkaURL,
	topic string,
) *kafka.Writer {

	return &kafka.Writer{
		Addr:         kafka.TCP(kafkaURL),
		Topic:        topic,               // tên topic để gửi message
		Balancer:     &kafka.RoundRobin{}, // phân phối message đến partition có ít dữ liệu nhất
		RequiredAcks: kafka.RequireAll,    // đảm bảo message được ghi vào tất cả replica trước khi xác nhận
		BatchSize:    1,                   // số lượng message gửi cùng lúc
		BatchTimeout: 0,                   // thời gian chờ gửi message
		Async:        false,               // true: gửi message không đợi response, false: đợi response
	}
}

func NewStockInfo(
	message,
	type_ string,
) *StockInfo {

	return &StockInfo{
		Message: message,
		Type:    type_,
	}
}

func actionStock(c *gin.Context) {

	s := NewStockInfo(
		c.Query("msg"),
		c.Query("type"),
	)

	body := KafkaMessage{
		Action: "stock_action",
		Info:   *s,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	msg := kafka.Message{
		Value: jsonBody,
	}

	err = kafkaProducer.WriteMessages(
		context.Background(),
		msg,
	)

	if err != nil {

		c.JSON(500, gin.H{
			"error": "Failed to write message",
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Message sent successfully",
	})
}

func processStockMessage(
	consumerID int,
	data KafkaMessage,
) {

	fmt.Printf(
		"[Consumer %d] Processing stock: %+v\n",
		consumerID,
		data,
	)

	switch data.Info.Type {

	case "MUA":

		fmt.Printf(
			"[Consumer %d] BUY stock %s\n",
			consumerID,
			data.Info.Message,
		)

		fmt.Printf(
			"[Consumer %d] BUY DONE %s\n",
			consumerID,
			data.Info.Message,
		)

	case "BAN":

		fmt.Printf(
			"[Consumer %d] SELL stock %s\n",
			consumerID,
			data.Info.Message,
		)

	default:

		fmt.Printf(
			"[Consumer %d] UNKNOWN TYPE %s\n",
			consumerID,
			data.Info.Type,
		)
	}
}

func RegisterConsumerATC(id int) {

	reader := getKafkaReader(
		kafkaURL,
		kafkaTopic,
		"stock_group",
	)

	defer reader.Close()

	fmt.Printf(
		"Consumer %d started\n",
		id,
	)

	for {

		m, err := reader.ReadMessage(
			context.Background(),
		)

		if err != nil {

			fmt.Printf(
				"[Consumer %d] Read Error: %v\n",
				id,
				err,
			)

			continue
		}

		fmt.Printf(
			"[Consumer %d] RAW MESSAGE: %s\n",
			id,
			string(m.Value),
		)

		var data KafkaMessage

		err = json.Unmarshal(
			m.Value,
			&data,
		)

		if err != nil {

			fmt.Printf(
				"[Consumer %d] JSON Parse Error: %v\n",
				id,
				err,
			)

			continue
		}

		switch data.Action {

		case "stock_action":
			processStockMessage(id, data)

		default:

			fmt.Printf(
				"[Consumer %d] Unknown action: %s\n",
				id,
				data.Action,
			)
		}
	}
}

func main() {

	r := gin.Default()

	r.SetTrustedProxies(nil)

	kafkaProducer = getKafkaWriter(
		kafkaURL,
		kafkaTopic,
	)

	defer kafkaProducer.Close()

	r.POST(
		"/action/stock",
		actionStock,
	)

	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)
	go RegisterConsumerATC(4)

	r.Run(":8089")
}
