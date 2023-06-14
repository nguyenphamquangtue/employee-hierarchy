package repository

import (
	"context"
	"employee-hierarchy-api/internal/dto"
	"employee-hierarchy-api/internal/pg"
	"reflect"
	"testing"
)

func TestEmployeeImpl_Find(t *testing.T) {
	// pg connect
	dbConnector := &pg.PostgreSQLConnector{}
	err := dbConnector.ConnectDBTest()
	if err != nil {
		t.Errorf("Failed to connect to the database test: %s", err)
	}

	supervisorID1 := 1
	supervisorID3 := 3

	type fields struct {
		dbConnector pg.DBConnector
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Employee
		wantErr bool
	}{
		{
			name: "Case 1",
			fields: fields{
				dbConnector: dbConnector,
			},
			args: args{
				ctx:  context.TODO(),
				name: "Nick",
			},
			want: &dto.Employee{
				ID:           3,
				Name:         "Nick",
				SupervisorID: &supervisorID1,
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
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				dbConnector: dbConnector,
			},
			args: args{
				ctx:  context.TODO(),
				name: "John",
			},
			want: &dto.Employee{
				ID:           3,
				Name:         "Nick",
				SupervisorID: &supervisorID1,
				Subordinates: []*dto.Employee{
					{
						ID:           2,
						Name:         "Sophie",
						SupervisorID: &supervisorID3,
						Subordinates: nil,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EmployeeImpl{
				dbConnector: tt.fields.dbConnector,
			}
			got, err := r.Find(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				return
			}
			if !compareEmployees(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
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
