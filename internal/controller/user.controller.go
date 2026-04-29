package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc - user controller
// us - user service

func (uc *UserController) GetInfoUser(c *gin.Context) {

	data, err := uc.userService.GetInfoUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
