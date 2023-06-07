package repository

import (
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/internal/pg"
	"gorm.io/gorm"
)

func (r employeeImpl) Find(name string, db *gorm.DB) (*dto.Employee, error) {
	if db == nil {
		db = pg.GetDB()
	}
	var employee dto.Employee
	if err := db.Preload("Subordinates").Where("name = ?", name).First(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r employeeImpl) FindByID(id int, db *gorm.DB) (*dto.Employee, error) {
	if db == nil {
		db = pg.GetDB()
	}
	var employee dto.Employee
	if err := db.Preload("Subordinates").Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r employeeImpl) Exist(name string, db *gorm.DB) (bool, error) {
	if db == nil {
		db = pg.GetDB()
	}
	var employee dto.Employee
	result := db.Where("name = ?", name).First(&employee)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
