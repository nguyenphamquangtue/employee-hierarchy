package service

import (
	"context"
	"employee-hierarchy-api/internal/pg"
	requestmodel "employee-hierarchy-api/pkg/model/request"
	"testing"
)

func Test_employeeImpl_Update(t *testing.T) {
	err := pg.ConnectDBTest()
	if err != nil {
		t.Errorf("Can not connect database to test")
	}

	tx := pg.GetDB().Begin()
	defer tx.Rollback()
	type args struct {
		ctx  context.Context
		eID  int
		data requestmodel.EmployeeUpdate
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				ctx: context.TODO(),
				eID: 5,
				data: requestmodel.EmployeeUpdate{
					SupervisorID: 3,
				},
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "Case 2",
			args: args{
				ctx: context.TODO(),
				eID: 5,
				data: requestmodel.EmployeeUpdate{
					SupervisorID: 5,
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := employeeImpl{}
			got, err := s.Update(tt.args.ctx, tt.args.eID, tt.args.data, tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
