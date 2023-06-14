package echocontext

import (
	"context"

	"github.com/labstack/echo/v4"
)

// GetContext ...
func GetContext(c echo.Context) context.Context {
	return c.Request().Context()
}
