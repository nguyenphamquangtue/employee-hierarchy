package routevalidation

import (
	"employee-hierarchy-api/external/utils/echocontext"
	"employee-hierarchy-api/internal/response"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserInterface interface {
	Login(next echo.HandlerFunc) echo.HandlerFunc
	Logout(next echo.HandlerFunc) echo.HandlerFunc
	Register(next echo.HandlerFunc) echo.HandlerFunc
}

type userImpl struct{}

func User() UserInterface {
	return userImpl{}
}

func (userImpl) Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload requestmodel.UserLogin
		)
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetPayload(c, payload)
		return next(c)
	}
}

func (userImpl) Logout(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the Authorization header value
		authHeader := c.Request().Header.Get("Authorization")
		// Ensure the Authorization header is not empty
		if authHeader == "" {
			return fmt.Errorf("Authorization header is missing")
		}

		// Split the Authorization header value to retrieve the token
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			return fmt.Errorf("Invalid Authorization header format")
		}
		tokenString := authParts[1]
		echocontext.SetPayload(c, tokenString)
		return next(c)
	}
}

func (userImpl) Register(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload requestmodel.UserRegister
		)
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetPayload(c, payload)
		return next(c)
	}
}
