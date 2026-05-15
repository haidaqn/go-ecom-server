package initialize

import (
	"log"
	"strconv"

	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"github.com/segmentio/kafka-go"
)

func InitKafka() {
	broker := global.Config.Kafka.Host + ":" + strconv.Itoa(global.Config.Kafka.Port)
	global.KafkaProducer = &kafka.Writer{
		Addr: kafka.TCP(broker),
		// topic mặc định
		Topic: "otp-auth-topic",
		// cơ chế phân phối message
		Balancer: &kafka.LeastBytes{},
		// optional
		RequiredAcks: kafka.RequireAll, // đảm bảo tất cả replica đã nhận được message
		Async:        false,            // gửi đồng bộ để đảm bảo message đã được gửi thành công
	}

}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Failed to close kafka producer: %v", err)
	}
}
