package entity

import (
	"reflect"
	"testing"
)

func TestAccount_Withdraw(t *testing.T) {
	type fields struct {
		Accountid int
		Balance   float32
	}
	type args struct {
		value int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Bill
		wantErr bool
	}{
		{
			name:    "Testing for withdraw 10",
			fields:  fields{Accountid: 12345, Balance: 12},
			args:    args{value: 10},
			want:    []Bill{{Value: 10}},
			wantErr: false,
		},
		{
			name:    "Testing for withdraw 20",
			fields:  fields{Accountid: 12345, Balance: 40},
			args:    args{value: 20},
			want:    []Bill{{Value: 10}, {Value: 10}},
			wantErr: false,
		},
		{
			name:    "Testing for withdraw 70",
			fields:  fields{Accountid: 12345, Balance: 70},
			args:    args{value: 70},
			want:    []Bill{{Value: 50}, {Value: 10}, {Value: 10}},
			wantErr: false,
		},
		{
			name:    "Testing for withdraw 100",
			fields:  fields{Accountid: 12345, Balance: 100},
			args:    args{value: 100},
			want:    []Bill{{Value: 50}, {Value: 50}},
			wantErr: false,
		},
		{
			name:    "Testing for withdraw 3",
			fields:  fields{Accountid: 12345, Balance: 14},
			args:    args{value: 3},
			want:    []Bill{{Value: 1}, {Value: 1}, {Value: 1}},
			wantErr: false,
		},
		{
			name:    "Test for withdraw 0",
			fields:  fields{Accountid: 12345, Balance: 14},
			args:    args{value: 0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &Account{
				ID:      tt.fields.Accountid,
				Balance: tt.fields.Balance,
			}
			got, err := cm.Withdraw(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CashMachine.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CashMachine.Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}
