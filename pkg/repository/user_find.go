package repository

import (
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/internal/pg"

	"gorm.io/gorm"
)

func (r userImpl) Find(user dto.User, db *gorm.DB) (*dto.User, error) {
	if db == nil {
		db = pg.GetDB()
	}

	result := db.Where("username = ?", user.Username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
