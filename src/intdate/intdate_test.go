package intdate

import (
	"testing"
)

func TestIntDate__NewIntDate(t *testing.T) {
	type args struct {
		date int
	}
	tests := []struct {
		name    string
		args    args
		want    IntDate
		wantErr bool
	}{
		{"", args{20200101}, 20200101, false},
		{"", args{20201231}, 20201231, false},
		{"", args{20200230}, -1, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIntDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIntDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewIntDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntDate__IntDate__ToPath(t *testing.T) {
	tests := []struct {
		name string
		date IntDate
		want string
	}{
		{"", IntDate(20200101), "2020/01/01"},
		{"", IntDate(20201231), "2020/12/31"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.ToPath(); got != tt.want {
				t.Errorf("IntDate.toPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
