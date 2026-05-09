package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/wire"
)

type AuthRouter struct {
}

func (a *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {

	authController, _ := wire.InitAuthController()

	authRouterPublic := Router.Group("auth")
	// authRouterPublic.Use(Limiter())
	{
		authRouterPublic.POST("/register", authController.Register)
		authRouterPublic.POST("/login")
		authRouterPublic.POST("/resend-otp")
		authRouterPublic.POST("/otp")
	}

	authRouterPrivate := Router.Group("auth")
	// authRouterPrivate.Use(Limiter())
	// authRouterPrivate.Use(Authe())
	// authRouterPrivate.Use(Permission())
	{
		authRouterPrivate.POST("/refresh")
	}
}
