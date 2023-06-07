package handler

import (
	"employee-hierarchy-api/external/utils/echocontext"
	"employee-hierarchy-api/internal/response"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	responsemodel "employee-hierarchy-api/pkg/model/response"
	"employee-hierarchy-api/pkg/service"
	"github.com/labstack/echo/v4"
)

type UserInterface interface {
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Register(c echo.Context) error
}

type userImpl struct{}

func User() UserInterface {
	return userImpl{}
}

// Login ...
func (userImpl) Login(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.UserLogin)
		s       = service.User()
	)

	// Login
	accessToken, err := s.Login(ctx, payload)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}

	return response.R200(c, responsemodel.UserLogin{AccessToken: accessToken}, "")

}

// Logout ...
func (userImpl) Logout(c echo.Context) error {
	var (
		ctx         = echocontext.GetContext(c)
		accessToken = echocontext.GetPayload(c).(string)
		s           = service.User()
	)

	// Logout
	err := s.Logout(ctx, accessToken)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, nil, "Token revoked successfully")

}

// Register ...
func (userImpl) Register(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.UserRegister)
		s       = service.User()
	)

	// register
	id, err := s.Register(ctx, payload)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}

	return response.R200(c, responsemodel.Upsert{ID: id}, "")

}
