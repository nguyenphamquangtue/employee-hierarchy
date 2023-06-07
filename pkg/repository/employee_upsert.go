package repository

import (
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/internal/pg"
	"gorm.io/gorm"
)

func (employeeImpl) Insert(employee dto.Employee, db *gorm.DB) (int, error) {
	if db == nil {
		db = pg.GetDB()
	}

	result := db.Create(&employee)
	if result.Error != nil {
		return 0, result.Error
	}

	return employee.ID, nil
}

func (r employeeImpl) Update(employee dto.Employee, db *gorm.DB) (int, error) {
	if db == nil {
		db = pg.GetDB()
	}

	result := db.Model(&dto.Employee{}).Where("id = ?", employee.ID).Updates(employee)
	if result.Error != nil {
		return 0, result.Error
	}

	return employee.ID, nil
}
