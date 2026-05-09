package service

import (
	"strings"

	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"golang.org/x/crypto/bcrypt"
	"go.uber.org/zap"
)

type IAuthService interface {
	Register(email string, password string) bool
}

type authService struct {
	userService IUservices
}

func NewAuthService(userService IUservices) IAuthService {
	return &authService{
		userService: userService,
	}
}

func (a *authService) Register(email string, password string) bool {
	if a.userService == nil {
		return false
	}

	normalizedEmail := strings.ToLower(strings.TrimSpace(email))
	normalizedPassword := strings.TrimSpace(password)
	if normalizedEmail == "" || normalizedPassword == "" {
		return false
	}

	if a.userService.GetUserByEmail(normalizedEmail) {
		if global.Logger != nil {
			global.Logger.Info("register failed: email already exists", zap.String("email", normalizedEmail))
		}
		return false
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(normalizedPassword), bcrypt.DefaultCost)
	if err != nil {
		if global.Logger != nil {
			global.Logger.Error("failed to hash password", zap.String("email", normalizedEmail), zap.Error(err))
		}
		return false
	}

	if !a.userService.CreateUser(normalizedEmail, string(hashedPassword)) {
		if global.Logger != nil {
			global.Logger.Error("failed to create user", zap.String("email", normalizedEmail))
		}
		return false
	}

	if global.Logger != nil {
		global.Logger.Info("user registered", zap.String("email", normalizedEmail))
	}

	return true
}
