package route

import (
	"employee-hierarchy-api/pkg/service"
	"github.com/labstack/echo/v4"
)

// Init ...
func Init(e *echo.Echo, service *service.Service) {
	g := e.Group("/fram")

	// cancellation effect context

	// Components
	user(g, service.UserService)
	employee(g, service.EmployeeService)
}
