package global

import (
	"database/sql"

	"github.com/haidaqn/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *zap.Logger
	MySQL  *gorm.DB
	Redis  *redis.Client
	MySQLC *sql.DB
)
