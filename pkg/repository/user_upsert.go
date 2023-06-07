package repository

import (
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/internal/pg"

	"gorm.io/gorm"
)

func (userImpl) Insert(user dto.User, db *gorm.DB) (int, error) {
	if db == nil {
		db = pg.GetDB()
	}

	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}
