package route

import (
	"employee-hierarchy-api/internal/middleware"
	"employee-hierarchy-api/pkg/handler"
	routevalidation "employee-hierarchy-api/pkg/route/validation"
	"employee-hierarchy-api/pkg/service"
	"github.com/labstack/echo/v4"
)

func employee(e *echo.Group, employeeService service.EmployeeInterface) {
	g := e.Group("/employees", middleware.RequireLogin)
	h := handler.Employee(employeeService)
	v := routevalidation.Employee()
	g.GET("", h.Find)
	g.POST("", h.Create, v.Create)
	g.PUT("/:id", h.Update, v.Update)
}
