package service

import (
	"context"
	"employee-hierarchy-api/internal/config/errorcode"
	"employee-hierarchy-api/internal/dto"
	middleware2 "employee-hierarchy-api/internal/middleware"
	"employee-hierarchy-api/internal/utils/hash"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"errors"
	"fmt"
)

func (s *UserImpl) Login(ctx context.Context, request requestmodel.UserLogin) (string, error) {
	// find user
	loginData, err := s.userRepository.Find(ctx, dto.User{
		Username: request.Username,
	})
	if err != nil {
		return "", errors.New(errorcode.UserDoesNotExist)
	}

	// compare password
	if err = hash.ComparePassword(loginData.Password, request.Password); err != nil {
		fmt.Println(err)
		return "", errors.New("invalid username or password")
	}

	accessToken, err := middleware2.GenerateJWT(loginData)
	if err != nil {
		return "", errors.New(errorcode.FailedToGenerateAccessToken)
	}

	return accessToken, nil
}

func (s *UserImpl) Logout(ctx context.Context, accessToken string) error {
	if err := middleware2.Revoke(accessToken); err != nil {
		return err
	}
	return nil
}
