package admin

import "github.com/gin-gonic/gin"

type UserRouterAdmin struct{}

func (u *UserRouterAdmin) InitUserRouter(Router *gin.RouterGroup) {

	userRouter := Router.Group("admin")
	{
		userRouter.GET("/user/:id/info")
		userRouter.POST("/user/create")
		userRouter.PUT("/user/:id/update")
		userRouter.DELETE("/user/:id/delete")
		userRouter.GET("/user")
	}
}
