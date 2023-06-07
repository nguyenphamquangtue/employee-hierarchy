package repository

import (
	"employee-hierarchy-api/external/dto"
	"gorm.io/gorm"
)

type EmployeeInterface interface {
	Exist(name string, db *gorm.DB) (bool, error)
	Find(name string, db *gorm.DB) (*dto.Employee, error)
	FindByID(id int, db *gorm.DB) (*dto.Employee, error)
	Insert(employee dto.Employee, db *gorm.DB) (int, error)
	Update(employee dto.Employee, db *gorm.DB) (int, error)
}

type employeeImpl struct{}

func Employee() EmployeeInterface {
	return employeeImpl{}
}
