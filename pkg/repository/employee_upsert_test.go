package repository

import (
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/internal/pg"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func Test_employeeImpl_Update(t *testing.T) {
	err := pg.ConnectDBTest()
	if err != nil {
		t.Errorf("Can not connect database to test")
	}

	supervisorID3 := 3

	tx := pg.GetDB().Begin()
	defer tx.Rollback()

	txDb := tx
	type args struct {
		employee dto.Employee
		db       *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "Case 1",
			args: args{
				employee: dto.Employee{
					ID:           5,
					SupervisorID: &supervisorID3,
				},
				db: txDb,
			},
			want:    5,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := employeeImpl{}
			got, err := r.Update(tt.args.employee, tt.args.db)
			if err != nil && err != tt.wantErr {
				fmt.Println(err)
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
