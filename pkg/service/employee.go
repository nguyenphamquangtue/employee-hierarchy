package service

import (
	"context"
	"employee-hierarchy-api/external/dto"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"gorm.io/gorm"
)

type EmployeeInterface interface {
	Find(ctx context.Context, name string) (result *dto.Employee, err error)
	Create(ctx context.Context, data requestmodel.EmployeeCreate) (int, error)
	Update(ctx context.Context, id int, data requestmodel.EmployeeUpdate, db *gorm.DB) (int, error)
}

type employeeImpl struct{}

func Employee() EmployeeInterface {
	return employeeImpl{}
}
