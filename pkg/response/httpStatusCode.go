package response

const (
	CodeSuccess        = 20001
	CodeParamInvalid   = 20003
	CodeOTPError       = 30001 // OTP is invalid
	CodeInvalidParams  = 30002 // Invalid parameters
	CodeHashEmailError = 40001 // Hash email error
	CodeSendOTPErr     = 40002 // Send OTP error
	CodeEmailExist     = 50001 // Email already exists
	CodeInternalServer = 50002 // Internal server error
)

var msg = map[int]string{
	CodeSuccess:        "Success",
	CodeParamInvalid:   "Email is invalid",
	CodeOTPError:       "OTP error",
	CodeHashEmailError: "Hash email error",
	CodeEmailExist:     "Email already exists",
	CodeInternalServer: "Internal server error",
	CodeSendOTPErr:     "Send OTP error",
	CodeInvalidParams:  "Invalid parameters",
}
