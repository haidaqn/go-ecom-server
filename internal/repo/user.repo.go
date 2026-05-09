package repo

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/po"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
	CreateUser(email string, hashedPassword string) bool
}

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUserByEmail(email string) bool {
	if global.MySQL == nil {
		return false
	}

	var user po.User
	err := global.MySQL.Where("username = ?", email).First(&user).Error
	if err == nil {
		return true
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	// Fallback an toàn: nếu DB lỗi, coi như đã tồn tại để tránh tạo luồng register sai.
	return true
}

func (ur *userRepository) CreateUser(email string, hashedPassword string) bool {
	if global.MySQL == nil {
		return false
	}

	user := po.User{
		ID:       uuid.New(),
		Username: strings.ToLower(strings.TrimSpace(email)),
		Password: hashedPassword,
		IsActive: false,
	}

	if err := global.MySQL.Create(&user).Error; err != nil {
		return false
	}

	return true
}
