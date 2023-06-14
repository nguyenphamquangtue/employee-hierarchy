package handler

import (
	"employee-hierarchy-api/internal/response"
	echocontext2 "employee-hierarchy-api/internal/utils/echocontext"
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

type UserImpl struct {
	userService service.UserInterface
}

func User(userService service.UserInterface) *UserImpl {
	return &UserImpl{
		userService: userService,
	}
}

// Login ...
func (h *UserImpl) Login(c echo.Context) error {
	var (
		ctx     = echocontext2.GetContext(c)
		payload = echocontext2.GetPayload(c).(requestmodel.UserLogin)
	)

	// Login
	accessToken, err := h.userService.Login(ctx, payload)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}

	return response.R200(c, responsemodel.UserLogin{AccessToken: accessToken}, "")

}

// Logout ...
func (h *UserImpl) Logout(c echo.Context) error {
	var (
		ctx         = echocontext2.GetContext(c)
		accessToken = echocontext2.GetPayload(c).(string)
	)

	// Logout
	err := h.userService.Logout(ctx, accessToken)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, nil, "Token revoked successfully")

}

// Register ...
func (h *UserImpl) Register(c echo.Context) error {
	var (
		ctx     = echocontext2.GetContext(c)
		payload = echocontext2.GetPayload(c).(requestmodel.UserRegister)
	)
	// register
	id, err := h.userService.Register(ctx, payload)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}

	return response.R200(c, responsemodel.Upsert{ID: id}, "")

}
