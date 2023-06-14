package service

import (
	"context"
	"employee-hierarchy-api/internal/config/errorcode"
	"employee-hierarchy-api/internal/dto"
	"employee-hierarchy-api/internal/utils/hash"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"errors"
	"log"
)

func (s *UserImpl) Register(ctx context.Context, request requestmodel.UserRegister) (id int, err error) {
	// hashPassword
	hashedPassword, err := hash.HashPassword(request.Password)
	if err != nil {
		log.Printf("error: %s", err)
		return 0, errors.New(errorcode.FailedToHashPassword)
	}

	id, err = s.userRepository.Insert(ctx, dto.User{
		Username: request.Username,
		Password: hashedPassword,
	})
	if err != nil {
		log.Printf("error: %s", err)
		return 0, errors.New(errorcode.UserInsertFailed)
	}
	return id, err
}
