package service

import (
	"context"
	"employee-hierarchy-api/internal/config/errorcode"
	"employee-hierarchy-api/internal/dto"
	"errors"
	"log"
)

func (s *EmployeeImpl) Find(ctx context.Context, name string) (result *dto.Employee, err error) {
	// find user
	employee, err := s.employeeRepository.Find(ctx, name)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return nil, errors.New(errorcode.EmployeeNotFound)
	}

	if employee.SupervisorID != nil {
		// find supervisor
		supervisor, err := s.employeeRepository.FindByID(ctx, *employee.SupervisorID)
		if err != nil {
			log.Printf("error: %s", err.Error())
			return nil, errors.New(errorcode.SupervisorNotFound)
		}
		employee.SupervisorName = supervisor.Name
	}

	return employee, nil
}
