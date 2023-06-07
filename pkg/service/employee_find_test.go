package service

import (
	"context"
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/internal/pg"
	"reflect"
	"testing"
)

func Test_employeeImpl_Find(t *testing.T) {
	err := pg.ConnectDBTest()
	if err != nil {
		t.Errorf("Can not connect database to test")
	}

	supervisorID1 := 1
	supervisorID3 := 3

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantResult *dto.Employee
		wantErr    error
	}{
		{
			name: "Case 1",
			args: args{
				ctx:  context.TODO(),
				name: "Jonas",
			},
			wantResult: &dto.Employee{
				ID:             1,
				Name:           "Jonas",
				SupervisorID:   nil,
				SupervisorName: "",
				Subordinates:   nil,
			},
			wantErr: nil,
		},
		{
			name: "Case 2",
			args: args{
				ctx:  context.TODO(),
				name: "Nick",
			},
			wantResult: &dto.Employee{
				ID:             3,
				Name:           "Nick",
				SupervisorID:   &supervisorID1,
				SupervisorName: "Jonas",
				Subordinates: []*dto.Employee{
					{
						ID:           2,
						Name:         "Sophie",
						SupervisorID: &supervisorID3,
						Subordinates: nil,
					},
					{
						ID:           4,
						Name:         "Pete",
						SupervisorID: &supervisorID3,
						Subordinates: nil,
					},
					{
						ID:           5,
						Name:         "Barbara",
						SupervisorID: &supervisorID3,
						Subordinates: nil,
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := employeeImpl{}
			gotResult, err := s.Find(tt.args.ctx, tt.args.name)
			if err != nil && err != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !compareEmployees(gotResult, tt.wantResult) {
				t.Errorf("Find() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func compareEmployees(a, b *dto.Employee) bool {
	if a.Subordinates == nil || b.Subordinates == nil {
		return a.ID == b.ID &&
			a.Name == b.Name &&
			reflect.DeepEqual(a.SupervisorID, b.SupervisorID)
	}
	return a.ID == b.ID &&
		a.Name == b.Name &&
		reflect.DeepEqual(a.SupervisorID, b.SupervisorID) &&
		a.SupervisorName == b.SupervisorName &&
		reflect.DeepEqual(a.Subordinates, b.Subordinates)
}
