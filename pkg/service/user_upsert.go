package service

import (
	"context"
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/external/utils/hash"
	"employee-hierarchy-api/internal/config/errorcode"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"employee-hierarchy-api/pkg/repository"
	"errors"
)

func (s userImpl) Register(ctx context.Context, request requestmodel.UserRegister) (id int, err error) {
	var (
		repo = repository.User()
	)

	// hashPassword
	hashedPassword, err := hash.HashPassword(request.Password)
	if err != nil {
		return 0, errors.New(errorcode.FailedToHashPassword)
	}

	id, err = repo.Insert(dto.User{
		Username: request.Username,
		Password: hashedPassword,
	}, nil)

	return id, nil
}
