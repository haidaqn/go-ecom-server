//go:build wireinject

package wire

import (
	"os"

	"github.com/google/wire"
	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/controller"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/repo"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/service"
	"github.com/haidaqn/go-ecommerce-backend-api/thrid_party/sendmail"
)

func ProvideRedisService() service.IRedisService {
	return service.NewRedisService(global.Redis)
}

func ProvideMailService() sendmail.IMailService {
	return sendmail.NewMailService(os.Getenv("RESEND_API_KEY"), os.Getenv("RESEND_FROM"))
}

func InitAuthRouteHandler() (*controller.AuthController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		ProvideRedisService,
		ProvideMailService,
		service.NewAuthService,
		controller.NewAuthController,
	)
	return new(controller.AuthController), nil
}
