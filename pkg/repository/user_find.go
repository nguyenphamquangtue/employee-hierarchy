package repository

import (
	"context"
	"employee-hierarchy-api/internal/dto"
)

func (r *UserImpl) Find(ctx context.Context, user dto.User) (*dto.User, error) {
	ctx, cancel := r.withCancellation(ctx)
	defer cancel()

	if err := r.dbConnector.GetDB().WithContext(ctx).Where("username = ?", user.Username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
