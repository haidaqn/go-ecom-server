package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/service"
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	success := a.authService.Register(req.Email, req.Password)
	if success {
		c.JSON(200, gin.H{"message": "Registration successful"})
	} else {
		c.JSON(500, gin.H{"message": "Registration failed"})
	}
}
