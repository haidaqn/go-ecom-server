package service

import (
	"strconv"
	"strings"

	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/utils"
	"github.com/haidaqn/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const registerOTPTTLSeconds = 300

func registerOTPRedisKey(email string) string {
	return "usr:" + email
}

// RegisterResult đồng bộ với pkg/response.ResponseData (HTTP luôn 200, phân biệt bằng Code).
type RegisterResult = response.ResponseData

type IAuthService interface {
	Register(email string, password string) RegisterResult
}

type authService struct {
	userService  IUservices
	redisService IRedisService
	kafkaService IKafkaService
}

func NewAuthService(userService IUservices, redisService IRedisService, kafkaService IKafkaService) IAuthService {
	return &authService{
		userService:  userService,
		redisService: redisService,
		kafkaService: kafkaService,
	}
}

func (a *authService) Register(email string, password string) RegisterResult {
	if a.userService == nil || a.redisService == nil || a.kafkaService == nil {
		return RegisterResult{
			Code:    response.CodeInternalServer,
			Message: "Hệ thống chưa sẵn sàng. Vui lòng thử lại sau.",
			Data:    nil,
		}
	}

	normalizedEmail := strings.ToLower(strings.TrimSpace(email))
	normalizedPassword := strings.TrimSpace(password)
	if normalizedEmail == "" || normalizedPassword == "" {
		return RegisterResult{
			Code:    response.CodeInvalidParams,
			Message: "Email hoặc mật khẩu không hợp lệ.",
			Data:    nil,
		}
	}

	if a.userService.GetUserByEmail(normalizedEmail) {
		if global.Logger != nil {
			global.Logger.Info("register failed: email already exists", zap.String("email", normalizedEmail))
		}
		return RegisterResult{
			Code:    response.CodeEmailExist,
			Message: "Email đã được sử dụng.",
			Data:    nil,
		}
	}

	otp := utils.GenerateSixDigitCode()
	otpKey := registerOTPRedisKey(normalizedEmail)
	otpStr := strconv.Itoa(otp)

	if !a.redisService.Set(otpKey, otpStr, registerOTPTTLSeconds) {
		if global.Logger != nil {
			global.Logger.Error("failed to set OTP", zap.String("email", normalizedEmail))
		}
		return RegisterResult{
			Code:    response.CodeInternalServer,
			Message: "Không thể tạo mã xác thực. Vui lòng thử lại sau.",
			Data:    nil,
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(normalizedPassword), bcrypt.DefaultCost)
	if err != nil {
		a.redisService.Delete(otpKey)
		if global.Logger != nil {
			global.Logger.Error("failed to hash password", zap.String("email", normalizedEmail), zap.Error(err))
		}
		return RegisterResult{
			Code:    response.CodeHashEmailError,
			Message: "Không thể xử lý mật khẩu. Vui lòng thử lại.",
			Data:    nil,
		}
	}

	if !a.userService.CreateUser(normalizedEmail, string(hashedPassword)) {
		a.redisService.Delete(otpKey)
		if global.Logger != nil {
			global.Logger.Error("failed to create user", zap.String("email", normalizedEmail))
		}
		return RegisterResult{
			Code:    response.CodeInternalServer,
			Message: "Không thể tạo tài khoản. Vui lòng thử lại sau.",
			Data:    nil,
		}
	}

	if global.Logger != nil {
		global.Logger.Info("user registered", zap.String("email", normalizedEmail))
	}

	// if err := a.mailService.SendMail([]string{normalizedEmail}, otpStr); err != nil {
	// 	a.redisService.Delete(otpKey)
	// 	if global.Logger != nil {
	// 		global.Logger.Error("failed to send OTP email", zap.String("email", normalizedEmail), zap.Error(err))
	// 	}
	// 	return RegisterResult{
	// 		Code:    response.CodeSendOTPErr,
	// 		Message: "Không thể gửi email xác thực. Vui lòng thử lại sau.",
	// 		Data:    nil,
	// 	}
	// }

	if err := a.kafkaService.PublishOTPEmail(normalizedEmail, otpStr); err != nil {
		a.redisService.Delete(otpKey)
		if global.Logger != nil {
			global.Logger.Error("failed to publish OTP email event", zap.String("email", normalizedEmail), zap.Error(err))
		}
		return RegisterResult{
			Code:    response.CodeSendOTPErr,
			Message: "Không thể gửi email xác thực. Vui lòng thử lại sau.",
			Data:    nil,
		}
	}

	return RegisterResult{
		Code:    response.CodeSuccess,
		Message: "Đăng ký thành công. Kiểm tra email để lấy mã OTP.",
		Data:    nil,
	}
}
