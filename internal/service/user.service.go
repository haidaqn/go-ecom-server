package service

import "github.com/haidaqn/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUser(id string) (string, error) {
	return us.userRepo.GetInfoUser(id)
}
