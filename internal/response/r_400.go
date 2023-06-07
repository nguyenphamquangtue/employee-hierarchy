package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// R400 bad request
func R400(c echo.Context, data interface{}, key string) error {
	// Get lang from echo context, if handle multilingualism

	if key == "" {
		key = CommonBadRequest
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	return sendResponse(c, http.StatusBadRequest, false, data, localeData.Message, localeData.Code)
}
