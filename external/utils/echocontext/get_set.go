package echocontext

import (
	"github.com/labstack/echo/v4"
)

// GetQuery ...
func GetQuery(c echo.Context) interface{} {
	return c.Get(KeyQuery)
}

// SetQuery ...
func SetQuery(c echo.Context, value interface{}) {
	c.Set(KeyQuery, value)
}

// GetPayload ...
func GetPayload(c echo.Context) interface{} {
	return c.Get(KeyPayload)
}

// SetPayload ...
func SetPayload(c echo.Context, value interface{}) {
	c.Set(KeyPayload, value)
}
