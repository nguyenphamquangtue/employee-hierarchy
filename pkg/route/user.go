package route

import (
	"employee-hierarchy-api/pkg/handler"
	"employee-hierarchy-api/pkg/service"

	routevalidation "employee-hierarchy-api/pkg/route/validation"

	"github.com/labstack/echo/v4"
)

func user(e *echo.Group, userService service.UserInterface) {
	g := e.Group("/users")
	h := handler.User(userService)
	v := routevalidation.User()

	g.GET("/login", h.Login, v.Login)
	g.POST("/logout", h.Logout, v.Logout)
	g.POST("", h.Register, v.Register)
}
