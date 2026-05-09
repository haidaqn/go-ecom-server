package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/global"
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

	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User
	authRouter := routers.RouterGroupApp.Auth

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}
	{
		userRouter.InitUserRouter(MainGroup)
	}
	{
		adminRouter.InitUserRouter(MainGroup)
	}
	{
		authRouter.InitAuthRouter(MainGroup)
	}

	return r
}
