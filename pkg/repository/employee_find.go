package repository

import (
	"context"
	"employee-hierarchy-api/internal/dto"
)

func (r *EmployeeImpl) Find(ctx context.Context, name string) (*dto.Employee, error) {
	ctx, cancel := r.withCancellation(ctx)
	defer cancel()

	var employee dto.Employee
	if err := r.dbConnector.GetDB().WithContext(ctx).Preload("Subordinates").Where("name = ?", name).First(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *EmployeeImpl) FindByID(ctx context.Context, id int) (*dto.Employee, error) {
	ctx, cancel := r.withCancellation(ctx)
	defer cancel()

	var employee dto.Employee
	if err := r.dbConnector.GetDB().WithContext(ctx).Preload("Subordinates").Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeImpl) Exist(ctx context.Context, name string) (bool, error) {
	ctx, cancel := r.withCancellation(ctx)
	defer cancel()

	var employee dto.Employee
	result := r.dbConnector.GetDB().WithContext(ctx).Where("name = ?", name).First(&employee)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
