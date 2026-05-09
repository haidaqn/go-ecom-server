package service

type IAuthService interface {
	Register(email string, purpose string) bool
}

type authService struct {
	userService IUservices
}

func NewAuthService(userService IUservices) IAuthService {
	return &authService{
		userService: userService,
	}
}

func (a *authService) Register(email string, purpose string) bool {
	return true
}
