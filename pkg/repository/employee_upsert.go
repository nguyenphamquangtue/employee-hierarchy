package repository

import (
	"context"
	"employee-hierarchy-api/internal/dto"
)

func (r *EmployeeImpl) Insert(ctx context.Context, employee dto.Employee) (int, error) {
	ctx, cancel := r.withCancellation(ctx)
	defer cancel()

	if err := r.dbConnector.GetDB().WithContext(ctx).Create(&employee).Error; err != nil {
		return 0, err
	}
	return employee.ID, nil
}

func (r *EmployeeImpl) Update(ctx context.Context, employee dto.Employee) (int, error) {
	ctx, cancel := r.withCancellation(ctx)
	defer cancel()

	if err := r.dbConnector.GetDB().WithContext(ctx).Model(&dto.Employee{}).Where("id = ?", employee.ID).Updates(employee).Error; err != nil {
		return 0, err
	}
	return employee.ID, nil
}
