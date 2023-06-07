package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// R403 forbidden
func R403(c echo.Context, data interface{}, key string) error {
	// Get lang from echo context, if handle multilingualism

	if key == "" {
		key = CommonForbidden
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	return sendResponse(c, http.StatusForbidden, false, data, localeData.Message, localeData.Code)
}
