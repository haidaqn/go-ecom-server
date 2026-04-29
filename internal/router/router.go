package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/controller"
)

func NewRoute() *gin.Engine {
	r := gin.Default()

	api_v1 := r.Group("/api/v1")

	api_v1.GET("/user/:id", controller.NewUserController().GetInfoUser)

	return r
}
