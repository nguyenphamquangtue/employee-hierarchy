package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// R404 not found
func R404(c echo.Context, data interface{}, key string) error {
	// Get lang from echo context, if handle multilingualism

	if key == "" {
		key = CommonNotFound
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	return sendResponse(c, http.StatusNotFound, false, data, localeData.Message, localeData.Code)
}
