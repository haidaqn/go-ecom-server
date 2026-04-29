package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/service"
	"github.com/haidaqn/go-ecommerce-backend-api/pkg/response"
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
		response.ErrorResponse(c, response.CodeInternalServer, err.Error())
		return
	}
	response.SuccessReponse(c, response.CodeSuccess, data, nil)
}
