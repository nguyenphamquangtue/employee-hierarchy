package handler

import (
	"employee-hierarchy-api/internal/response"
	echocontext2 "employee-hierarchy-api/internal/utils/echocontext"
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

type EmployeeImpl struct {
	employeeService service.EmployeeInterface
}

func Employee(employeeService service.EmployeeInterface) *EmployeeImpl {
	return &EmployeeImpl{
		employeeService: employeeService,
	}
}

// Find ...
func (h *EmployeeImpl) Find(c echo.Context) error {
	var (
		ctx  = echocontext2.GetContext(c)
		name = c.QueryParam("name")
	)

	employee, err := h.employeeService.Find(ctx, name)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, employee, "")
}

// Create ...
func (h *EmployeeImpl) Create(c echo.Context) error {
	var (
		ctx     = echocontext2.GetContext(c)
		payload = echocontext2.GetPayload(c).(requestmodel.EmployeeCreate)
	)

	id, err := h.employeeService.Create(ctx, payload)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}

	return response.R200(c, responsemodel.Upsert{ID: id}, "")

}

// Update ...
func (h *EmployeeImpl) Update(c echo.Context) error {
	var (
		ctx        = echocontext2.GetContext(c)
		payload    = echocontext2.GetPayload(c).(requestmodel.EmployeeUpdate)
		employeeID = c.Param("id")
	)

	eID, err := strconv.Atoi(employeeID)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}

	id, err := h.employeeService.Update(ctx, eID, payload, nil)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}

	return response.R200(c, responsemodel.Upsert{ID: id}, "")

}
