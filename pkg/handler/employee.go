package handler

import (
	"employee-hierarchy-api/external/utils/echocontext"
	"employee-hierarchy-api/internal/response"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	responsemodel "employee-hierarchy-api/pkg/model/response"
	"employee-hierarchy-api/pkg/service"
	"github.com/labstack/echo/v4"
	"strconv"
)

type EmployeeInterface interface {
	Find(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
}

type employeeImpl struct{}

func Employee() EmployeeInterface {
	return employeeImpl{}
}

// Find ...
func (employeeImpl) Find(c echo.Context) error {
	var (
		ctx  = echocontext.GetContext(c)
		s    = service.Employee()
		name = c.QueryParam("name")
	)

	employee, err := s.Find(ctx, name)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, employee, "")
}

// Create ...
func (employeeImpl) Create(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.EmployeeCreate)
		s       = service.Employee()
	)

	id, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}

	return response.R200(c, responsemodel.Upsert{ID: id}, "")

}

// Update ...
func (employeeImpl) Update(c echo.Context) error {
	var (
		ctx        = echocontext.GetContext(c)
		payload    = echocontext.GetPayload(c).(requestmodel.EmployeeUpdate)
		s          = service.Employee()
		employeeID = c.Param("id")
	)

	eID, err := strconv.Atoi(employeeID)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}

	id, err := s.Update(ctx, eID, payload, nil)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}

	return response.R200(c, responsemodel.Upsert{ID: id}, "")

}
