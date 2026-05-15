package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	otpKafkaAction       = "send_otp_email"
	otpKafkaWriteTimeout = 10 * time.Second
)

var (
	ErrKafkaProducerNil = errors.New("kafka: producer not initialized")
	ErrKafkaEmptyEmail  = errors.New("kafka: email is empty")
	ErrKafkaEmptyOTP    = errors.New("kafka: otp is empty")
)

type OTPEmailMessage struct {
	Action string `json:"action"`
	Email  string `json:"email"`
	OTP    string `json:"otp"`
}

type IKafkaService interface {
	PublishOTPEmail(email, otp string) error
}

type kafkaService struct {
	producer *kafka.Writer
}

func NewKafkaService(producer *kafka.Writer) IKafkaService {
	return &kafkaService{producer: producer}
}

func (k *kafkaService) PublishOTPEmail(email, otp string) error {
	if k == nil || k.producer == nil {
		return ErrKafkaProducerNil
	}

	email = strings.TrimSpace(email)
	otp = strings.TrimSpace(otp)
	if email == "" {
		return ErrKafkaEmptyEmail
	}
	if otp == "" {
		return ErrKafkaEmptyOTP
	}

	payload, err := json.Marshal(OTPEmailMessage{
		Action: otpKafkaAction,
		Email:  email,
		OTP:    otp,
	})
	if err != nil {
		return fmt.Errorf("kafka: marshal otp message: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), otpKafkaWriteTimeout)
	defer cancel()

	if err := k.producer.WriteMessages(ctx, kafka.Message{Value: payload}); err != nil {
		return fmt.Errorf("kafka: write otp message: %w", err)
	}

	return nil
}
