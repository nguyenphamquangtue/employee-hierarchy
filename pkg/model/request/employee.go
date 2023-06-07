package requestmodel

import (
	"employee-hierarchy-api/internal/config/errorcode"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type EmployeeCreate struct {
	Name string `json:"name"`
}

func (m EmployeeCreate) Validate() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(
			&m.Name,
			validation.Required.Error(errorcode.EmployeeNameRequired),
		),
	)
}

type EmployeeUpdate struct {
	SupervisorID int `json:"supervisor_id"`
}

func (m EmployeeUpdate) Validate() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(
			&m.SupervisorID,
			validation.Required.Error(errorcode.SupervisorIDRequired),
		),
	)
}

type Employee struct {
	Name       string `json:"name"`
	Supervisor string `json:"supervisor"`
}

func (m Employee) Validate() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(
			&m.Name,
			validation.Required.Error(errorcode.UsernameInvalid),
		),
		validation.Field(
			&m.Supervisor,
			validation.Required.Error(errorcode.PasswordInvalid),
		),
	)
}
