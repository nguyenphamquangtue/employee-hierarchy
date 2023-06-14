package service

import (
	"context"
	"employee-hierarchy-api/internal/config/errorcode"
	"employee-hierarchy-api/internal/dto"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"errors"
	"gorm.io/gorm"
	"log"
)

func (s *EmployeeImpl) Create(ctx context.Context, data requestmodel.EmployeeCreate) (int, error) {

	exist, err := s.employeeRepository.Exist(ctx, data.Name)
	if exist || err != nil {
		log.Printf("error: %s", err.Error())
		return 0, errors.New(errorcode.EmployeeExisted)
	}

	id, err := s.employeeRepository.Insert(ctx, dto.Employee{
		Name: data.Name,
	})
	if err != nil {
		log.Printf("error: %s", err.Error())
		return 0, errors.New(errorcode.EmployeeInsertFailed)
	}

	return id, err
}

func (s *EmployeeImpl) Update(ctx context.Context, eID int, data requestmodel.EmployeeUpdate, db *gorm.DB) (int, error) {
	if eID == data.SupervisorID {
		return 0, errors.New(errorcode.SupervisorCannotBeTheSameAsTheEmployee)
	}

	// find user
	employee, err := s.employeeRepository.FindByID(ctx, eID)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return 0, errors.New(errorcode.EmployeeNotFound)
	}

	// find supervisor
	supervisor, err := s.employeeRepository.FindByID(ctx, data.SupervisorID)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return 0, errors.New(errorcode.SupervisorNotFound)
	}

	// Check if there is a cycle in the supervisor hierarchy
	visited := make(map[int]bool)
	if s.HasCycle(ctx, supervisor, employee, visited) {
		return 0, errors.New(errorcode.CycleDetectedInSupervisorHierarchy)
	}

	supervisorID := data.SupervisorID
	id, err := s.employeeRepository.Update(ctx, dto.Employee{
		ID:           eID,
		SupervisorID: &supervisorID,
	})
	return id, err
}

func (s *EmployeeImpl) HasCycle(ctx context.Context, employee *dto.Employee, supervisor *dto.Employee, visited map[int]bool) bool {
	if employee.ID == supervisor.ID {
		return true
	}
	visited[employee.ID] = true
	if employee.SupervisorID != nil {
		nextSupervisor, err := s.employeeRepository.FindByID(ctx, *employee.SupervisorID)
		if err != nil {
			return false
		}
		if nextSupervisor != nil {
			if visited[nextSupervisor.ID] {
				return false
			}
			return s.HasCycle(ctx, nextSupervisor, supervisor, visited)
		}
	}
	return false
}
