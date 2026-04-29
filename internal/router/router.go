package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/controller"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/middleware"
)

func NewRoute() *gin.Engine {
	r := gin.Default()

	go middleware.CleanupClient()

	r.Use(middleware.LoggerMiddleware(), middleware.RateLimitingMiddleware())
	// r.Use(middleware.LoggerMiddleware(), middleware.ApiKeyMiddleware(), middleware.RateLimitingMiddleware())

	api_v1 := r.Group("/api/v1")

	api_v1.GET("/user/:id", controller.NewUserController().GetInfoUser)

	return r
}
