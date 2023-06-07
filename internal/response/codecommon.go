package response

// Constant
const (
	CommonSuccess      = "common_success"
	CommonBadRequest   = "common_bad_request"
	CommonUnauthorized = "common_unauthorized"
	CommonNotFound     = "common_not_found"
	CommonForbidden    = "common_forbidden"
)

const (
	commonSuccessCode = iota + 1
	commonBadRequestCode
	commonUnauthorizedCode
	commonNotFoundCode
	commonForbiddenCode
)

// 1-99
var common = []Code{
	{
		Key:     CommonSuccess,
		Code:    commonSuccessCode,
		Message: "Success",
	},
	{
		Key:     CommonBadRequest,
		Code:    commonBadRequestCode,
		Message: "Invalid data",
	},
	{
		Key:     CommonUnauthorized,
		Code:    commonUnauthorizedCode,
		Message: "You are unauthorized to perform this action",
	},
	{
		Key:     CommonNotFound,
		Code:    commonNotFoundCode,
		Message: "Data not found",
	},
	{
		Key:     CommonForbidden,
		Code:    commonForbiddenCode,
		Message: "Access denied",
	},
}
