package repository

import (
	"employee-hierarchy-api/internal/pg"
)

type Repository struct {
	dbConnector        pg.DBConnector
	UserRepository     UserInterface
	EmployeeRepository EmployeeInterface
}

func Init(dbConnector pg.DBConnector) *Repository {
	return &Repository{
		dbConnector:        dbConnector,
		UserRepository:     User(dbConnector),
		EmployeeRepository: Employee(dbConnector),
	}
}
