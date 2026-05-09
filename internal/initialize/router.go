package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/middlewares"
	routers "github.com/haidaqn/go-ecommerce-backend-api/internal/router"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.Use(
		middlewares.ValidatorMiddleware(),
		// middlewares.ApiKeyMiddleware(),
		// middlewares.LoggerMiddleware(),
		// middlewares.ErrorHandlerMiddleware(),
		// middlewares.CorsMiddleware(),
		// middlewares.AuthenticateMiddleware(),
	)

	authRouter := routers.RouterGroupApp.Auth

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}
	{
		authRouter.InitAuthRouter(MainGroup)
	}

	return r
}
