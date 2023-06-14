package service

import (
	"context"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"employee-hierarchy-api/pkg/repository"
)

type UserInterface interface {
	Login(ctx context.Context, request requestmodel.UserLogin) (string, error)
	Logout(ctx context.Context, accessToken string) error
	Register(ctx context.Context, request requestmodel.UserRegister) (int, error)
}

type UserImpl struct {
	userRepository repository.UserInterface
}

func User(userRepository repository.UserInterface) *UserImpl {
	return &UserImpl{
		userRepository: userRepository,
	}
}
