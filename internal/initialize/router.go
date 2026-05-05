package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/controller"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})

	go middleware.CleanupClient()

	r.Use(
		middleware.LoggerMiddleware(),
		middleware.RateLimitingMiddleware(),
		middleware.CorsMiddleware(),
		middleware.ApiKeyMiddleware(),
		middleware.AuthenticateMiddleware(),
		middleware.ErrorHandlerMiddleware(),
	)

	api_v1 := r.Group("/api/v1")

	api_v1.GET("/user/:id", controller.NewUserController().GetInfoUser)

	return r
}
