package service

import (
	"context"
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/internal/config/errorcode"
	"employee-hierarchy-api/pkg/repository"
	"errors"
)

func (s employeeImpl) Find(ctx context.Context, name string) (result *dto.Employee, err error) {
	var (
		repo = repository.Employee()
	)

	// find user
	employee, err := repo.Find(name, nil)
	if err != nil {
		return nil, errors.New(errorcode.EmployeeNotFound)
	}
	if employee.SupervisorID != nil {
		// find supervisor
		supervisor, err := repo.FindByID(*employee.SupervisorID, nil)
		if err != nil {
			return nil, errors.New(errorcode.SupervisorNotFound)
		}
		employee.SupervisorName = supervisor.Name
	}

	return employee, nil
}
