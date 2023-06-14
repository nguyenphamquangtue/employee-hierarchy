package repository

import (
	"context"
	"employee-hierarchy-api/internal/dto"
	"employee-hierarchy-api/internal/pg"
)

type EmployeeInterface interface {
	Exist(ctx context.Context, name string) (bool, error)
	Find(ctx context.Context, name string) (*dto.Employee, error)
	FindByID(ctx context.Context, id int) (*dto.Employee, error)
	Insert(ctx context.Context, employee dto.Employee) (int, error)
	Update(ctx context.Context, employee dto.Employee) (int, error)
}

type EmployeeImpl struct {
	dbConnector pg.DBConnector
}

func Employee(dbConnector pg.DBConnector) *EmployeeImpl {
	return &EmployeeImpl{
		dbConnector: dbConnector,
	}
}

func (r *EmployeeImpl) withCancellation(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}
