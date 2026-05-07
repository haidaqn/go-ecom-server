package auth

import "github.com/gin-gonic/gin"

type AuthRouter struct{}

func (a *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {

	authRouterPublic := Router.Group("auth")
	// authRouterPublic.Use(Limiter())
	{
		authRouterPublic.POST("/register")
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
