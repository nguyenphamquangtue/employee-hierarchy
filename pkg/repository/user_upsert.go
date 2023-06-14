package repository

import (
	"context"
	"employee-hierarchy-api/internal/dto"
)

func (r *UserImpl) Insert(ctx context.Context, user dto.User) (int, error) {
	ctx, cancel := r.withCancellation(ctx)
	defer cancel()

	if err := r.dbConnector.GetDB().WithContext(ctx).Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}
