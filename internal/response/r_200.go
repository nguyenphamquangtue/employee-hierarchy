package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// R200 response success
func R200(c echo.Context, data interface{}, key string) error {
	// Get lang from echo context, if handle multilingualism
	if key == "" {
		key = CommonSuccess
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	return sendResponse(c, http.StatusOK, true, data, localeData.Message, localeData.Code)
}
