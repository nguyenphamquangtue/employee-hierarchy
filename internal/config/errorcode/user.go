package errorcode

import "employee-hierarchy-api/internal/response"

const (
	UsernameInvalid             = "username_invalid"
	UserDoesNotExist            = "user_does_not_existed"
	PasswordInvalid             = "password_invalid"
	FailedToHashPassword        = "failed_to_hash_password"
	FailedToGenerateAccessToken = "failed_to_generate_access_token"
)

const (
	userNameInvalidCode = iota + 600
	userDoesNotExistCode
	passwordInvalidCode
	failedToHashPassword
	failedToGenerateAccessToken
)

var user = []response.Code{
	{
		Key:     UsernameInvalid,
		Code:    userNameInvalidCode,
		Message: "Username Invalid",
	},
	{
		Key:     UserDoesNotExist,
		Code:    userDoesNotExistCode,
		Message: "Username Does Not Existed",
	},
	{
		Key:     PasswordInvalid,
		Code:    passwordInvalidCode,
		Message: "Password Invalid",
	},
	{
		Key:     FailedToHashPassword,
		Code:    failedToHashPassword,
		Message: "Failed to Generate Access Token",
	},
	{
		Key:     FailedToGenerateAccessToken,
		Code:    failedToGenerateAccessToken,
		Message: "Failed to Hash Password",
	},
}
