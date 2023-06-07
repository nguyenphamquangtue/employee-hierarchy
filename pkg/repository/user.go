package repository

import (
	"employee-hierarchy-api/external/dto"

	"gorm.io/gorm"
)

type UserInterface interface {
	Find(user dto.User, db *gorm.DB) (*dto.User, error)
	Insert(user dto.User, db *gorm.DB) (int, error)
}

type userImpl struct{}

func User() UserInterface {
	return userImpl{}
}
