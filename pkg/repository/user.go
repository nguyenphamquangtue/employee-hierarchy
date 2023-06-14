package repository

import (
	"context"
	"employee-hierarchy-api/internal/dto"
	"employee-hierarchy-api/internal/pg"
)

type UserInterface interface {
	Find(ctx context.Context, user dto.User) (*dto.User, error)
	Insert(ctx context.Context, user dto.User) (int, error)
}

type UserImpl struct {
	dbConnector pg.DBConnector
}

func User(dbConnector pg.DBConnector) *UserImpl {
	return &UserImpl{
		dbConnector: dbConnector,
	}
}

func (r *UserImpl) withCancellation(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}
