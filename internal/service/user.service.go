package service

import "github.com/haidaqn/go-ecommerce-backend-api/internal/repo"

type IUservices interface {
	GetUserByEmail(email string) bool
	CreateUser(email string, hashedPassword string) bool
}

type userService struct {
	userRepository repo.IUserRepository
}

func NewUserService(userRepository repo.IUserRepository) IUservices {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) GetUserByEmail(email string) bool {
	return u.userRepository.GetUserByEmail(email)
}

func (u *userService) CreateUser(email string, hashedPassword string) bool {
	return u.userRepository.CreateUser(email, hashedPassword)
}
