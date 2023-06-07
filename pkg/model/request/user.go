package requestmodel

import (
	"employee-hierarchy-api/internal/config/errorcode"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m UserLogin) Validate() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(
			&m.Username,
			validation.Required.Error(errorcode.UsernameInvalid),
		),
		validation.Field(
			&m.Password,
			validation.Required.Error(errorcode.PasswordInvalid),
		),
	)
}

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m UserRegister) Validate() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(
			&m.Username,
			validation.Required.Error(errorcode.UsernameInvalid),
		),
		validation.Field(
			&m.Password,
			validation.Required.Error(errorcode.PasswordInvalid),
		),
	)
}
