package route

import (
	"github.com/labstack/echo/v4"
)

// Init ...
func Init(e *echo.Echo) {
	g := e.Group("/fram")

	// Components
	user(g)
	employee(g)
}
