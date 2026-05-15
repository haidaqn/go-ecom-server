package global

import (
	"database/sql"

	"github.com/haidaqn/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *zap.Logger
	MySQL         *gorm.DB
	Redis         *redis.Client
	MySQLC        *sql.DB
	KafkaProducer *kafka.Writer
)
