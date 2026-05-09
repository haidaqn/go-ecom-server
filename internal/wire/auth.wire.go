package wire

import (
	"github.com/google/wire"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/controller"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/repo"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/service"
)

func InitAuthController() (*controller.AuthController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		service.NewAuthService,
		controller.NewAuthController,
	)
	return new(controller.AuthController), nil
}
