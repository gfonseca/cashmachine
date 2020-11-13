package entity

import (
	"reflect"
	"testing"
)

func TestBill_String(t *testing.T) {
	type fields struct {
		Value int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Test String output for $50 bill",
			fields: fields{Value: 50},
			want:   "<Bill $50>",
		},
		{
			name:   "Test String output for $10 bill",
			fields: fields{Value: 10},
			want:   "<Bill $10>",
		},
		{
			name:   "Test String output for $5 bill",
			fields: fields{Value: 5},
			want:   "<Bill $5>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bill{
				Value: tt.fields.Value,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("Bill.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBill(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name    string
		args    args
		want    *Bill
		wantErr bool
	}{
		{
			name:    "Building a valid Bill value 10",
			args:    args{value: 10},
			want:    &Bill{Value: 10},
			wantErr: false,
		},
		{
			name:    "Building a valid Bill value 50",
			args:    args{value: 50},
			want:    &Bill{Value: 50},
			wantErr: false,
		},
		{
			name:    "Testing for invalid Bill value",
			args:    args{value: 13},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBill(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBill() = %v, want %v", got, tt.want)
			}
		})
	}
}
