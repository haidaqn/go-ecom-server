//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/controller"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/repo"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/service"
)

func ProvideRedisService() service.IRedisService {
	return service.NewRedisService(global.Redis)
}

func ProvideKafkaService() service.IKafkaService {
	return service.NewKafkaService(global.KafkaProducer)
}

func InitAuthRouteHandler() (*controller.AuthController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		ProvideRedisService,
		ProvideKafkaService,
		service.NewAuthService,
		controller.NewAuthController,
	)
	return new(controller.AuthController), nil
}
