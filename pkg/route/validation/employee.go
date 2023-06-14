package routevalidation

import (
	"employee-hierarchy-api/internal/response"
	"employee-hierarchy-api/internal/utils/echocontext"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"github.com/labstack/echo/v4"
)

type EmployeeInterface interface {
	Create(next echo.HandlerFunc) echo.HandlerFunc
	Update(next echo.HandlerFunc) echo.HandlerFunc
}

type employeeImpl struct{}

func Employee() EmployeeInterface {
	return employeeImpl{}
}

func (employeeImpl) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload requestmodel.EmployeeCreate
		)
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetPayload(c, payload)
		return next(c)
	}
}

func (employeeImpl) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload requestmodel.EmployeeUpdate
		)
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetPayload(c, payload)
		return next(c)
	}
}
