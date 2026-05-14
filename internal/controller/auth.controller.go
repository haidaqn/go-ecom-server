package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/service"
	"github.com/haidaqn/go-ecommerce-backend-api/pkg/response"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (a *AuthController) Register(c *gin.Context) {
	type request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, response.CodeInvalidParams, err.Error())
		return
	}

	res := a.authService.Register(req.Email, req.Password)
	c.JSON(http.StatusOK, res)
}
