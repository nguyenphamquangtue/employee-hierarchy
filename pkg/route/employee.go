package route

import (
	"employee-hierarchy-api/external/middleware"
	"employee-hierarchy-api/pkg/handler"
	routevalidation "employee-hierarchy-api/pkg/route/validation"
	"github.com/labstack/echo/v4"
)

func employee(e *echo.Group) {
	g := e.Group("/employees", middleware.RequireLogin)
	h := handler.Employee()
	v := routevalidation.Employee()

	g.GET("", h.Find)
	g.POST("", h.Create, v.Create)
	g.PUT("/:id", h.Update, v.Update)
}
