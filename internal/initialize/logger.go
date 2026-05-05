package initialize

import (
	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"github.com/haidaqn/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	logger := logger.NewLogger(global.Config.Logger)
	global.Logger = logger.Logger
}
