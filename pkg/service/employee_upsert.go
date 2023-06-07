package service

import (
	"context"
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/internal/config/errorcode"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"employee-hierarchy-api/pkg/repository"
	"errors"
	"gorm.io/gorm"
)

func (s employeeImpl) Create(ctx context.Context, data requestmodel.EmployeeCreate) (int, error) {
	var (
		repo = repository.Employee()
	)

	exist, err := repo.Exist(data.Name, nil)
	if exist {
		return 0, errors.New(errorcode.EmployeeExisted)
	}

	id, err := repo.Insert(dto.Employee{
		Name: data.Name,
	}, nil)

	return id, err
}

func (s employeeImpl) Update(ctx context.Context, eID int, data requestmodel.EmployeeUpdate, db *gorm.DB) (int, error) {
	var (
		repo = repository.Employee()
	)

	if eID == data.SupervisorID {
		return 0, errors.New(errorcode.SupervisorCannotBeTheSameAsTheEmployee)
	}

	// find user
	employee, err := repo.FindByID(eID, db)
	if err != nil {
		return 0, errors.New(errorcode.EmployeeNotFound)
	}

	// find supervisor
	supervisor, err := repo.FindByID(data.SupervisorID, db)
	if err != nil {
		return 0, errors.New(errorcode.SupervisorNotFound)
	}

	// Check if there is a cycle in the supervisor hierarchy
	visited := make(map[int]bool)
	if HasCycle(supervisor, employee, visited) {
		return 0, errors.New(errorcode.CycleDetectedInSupervisorHierarchy)
	}

	//for _, e := range employee.Subordinates {
	//	if e.ID == supervisor.ID {
	//		return 0, errors.New(errorcode.SubordinateCannotBeASupervisor)
	//	}
	//}

	supervisorID := data.SupervisorID
	id, err := repo.Update(dto.Employee{
		ID:           eID,
		SupervisorID: &supervisorID,
	}, db)
	return id, err
}

func HasCycle(employee *dto.Employee, supervisor *dto.Employee, visited map[int]bool) bool {
	if employee.ID == supervisor.ID {
		return true
	}
	visited[int(employee.ID)] = true
	if employee.SupervisorID != nil {
		nextSupervisor, err := repository.Employee().FindByID(int(*employee.SupervisorID), nil)
		if err != nil {
			return false
		}
		if nextSupervisor != nil {
			if visited[int(nextSupervisor.ID)] {
				return false
			}
			return HasCycle(nextSupervisor, supervisor, visited)
		}
	}
	return false
}
