package service

import (
	"context"
	"employee-hierarchy-api/internal/dto"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"employee-hierarchy-api/pkg/repository"
	"gorm.io/gorm"
)

type EmployeeInterface interface {
	Find(ctx context.Context, name string) (result *dto.Employee, err error)
	Create(ctx context.Context, data requestmodel.EmployeeCreate) (int, error)
	Update(ctx context.Context, id int, data requestmodel.EmployeeUpdate, db *gorm.DB) (int, error)
}

type EmployeeImpl struct {
	employeeRepository repository.EmployeeInterface
}

func Employee(employeeRepository repository.EmployeeInterface) *EmployeeImpl {
	return &EmployeeImpl{
		employeeRepository: employeeRepository,
	}
}
