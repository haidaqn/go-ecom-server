package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userRouterPublic := Router.Group("user")
	// userRouterPublic.Use(Limiter())
	{
		userRouterPublic.GET("/:id/info", func(c *gin.Context) {
			c.String(200, "hello")
		})
	}

	userRouterPrivate := Router.Group("user")
	{
		userRouterPrivate.PUT("/")
		userRouterPrivate.GET("/")
	}
}
