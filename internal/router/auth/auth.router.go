package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/controller"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/repo"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/service"
)

type AuthRouter struct {
	authController controller.IAuthController
}

func (a *AuthRouter) ensureDependencies() {
	if a.authController != nil {
		return
	}

	userRepo := repo.NewUserRepository()
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService)
	a.authController = controller.NewAuthController(authService)
}

func notImplementedHandler(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}

func (a *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	a.ensureDependencies()

	authRouterPublic := Router.Group("auth")
	// authRouterPublic.Use(Limiter())
	{
		authRouterPublic.POST("/register", a.authController.Register)
		authRouterPublic.POST("/login", notImplementedHandler)
		authRouterPublic.POST("/resend-otp", notImplementedHandler)
		authRouterPublic.POST("/otp", notImplementedHandler)
	}

	authRouterPrivate := Router.Group("auth")
	// authRouterPrivate.Use(Limiter())
	// authRouterPrivate.Use(Authe())
	// authRouterPrivate.Use(Permission())
	{
		authRouterPrivate.POST("/refresh", notImplementedHandler)
	}
}
