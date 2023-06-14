package service

import (
	"employee-hierarchy-api/pkg/repository"
)

type Service struct {
	UserService     UserInterface
	EmployeeService EmployeeInterface
}

func Init(repository *repository.Repository) *Service {
	return &Service{
		UserService:     User(repository.UserRepository),
		EmployeeService: Employee(repository.EmployeeRepository),
	}
}
