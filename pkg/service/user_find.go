package service

import (
	"context"
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/external/middleware"
	"employee-hierarchy-api/external/utils/hash"
	"employee-hierarchy-api/internal/config/errorcode"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"employee-hierarchy-api/pkg/repository"
	"errors"
	"fmt"
)

func (s userImpl) Login(ctx context.Context, request requestmodel.UserLogin) (string, error) {
	var (
		repo = repository.User()
	)

	// find user
	loginData, err := repo.Find(dto.User{
		Username: request.Username,
	}, nil)
	if err != nil {
		return "", errors.New(errorcode.UserDoesNotExist)
	}

	// compare password
	if err = hash.ComparePassword(loginData.Password, request.Password); err != nil {
		fmt.Println(err)
		return "", errors.New("invalid username or password")
	}

	accessToken, err := middleware.GenerateJWT(loginData)
	if err != nil {
		return "", errors.New(errorcode.FailedToGenerateAccessToken)
	}

	return accessToken, nil
}

func (s userImpl) Logout(ctx context.Context, accessToken string) error {
	if err := middleware.Revoke(accessToken); err != nil {
		return err
	}
	return nil
}
