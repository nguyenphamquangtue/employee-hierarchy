package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// R401 unauthorized
func R401(c echo.Context, data interface{}, key string) error {
	// Get lang from echo context, if handle multilingualism

	if key == "" {
		key = CommonUnauthorized
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	return sendResponse(c, http.StatusUnauthorized, false, data, localeData.Message, localeData.Code)
}
