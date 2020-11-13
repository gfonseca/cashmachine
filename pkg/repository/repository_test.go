package repository

import (
	"cashmachine/pkg/entity"
	"cashmachine/pkg/mocks"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
)

func TestNewRepository(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		dbdriver DBInterface
	}
	tests := []struct {
		name string
		want func(DBInterface) *Repository
	}{
		{
			name: "Testing new repository",
			want: func(dbd DBInterface) *Repository {
				return &Repository{dbd}
			},
		},
	}
	for _, tt := range tests {
		dbd := mocks.NewMockDBInterface(mockCtrl)
		want := tt.want(dbd)
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(dbd); !reflect.DeepEqual(got, want) {
				t.Errorf("NewRepository() = %v, want %v", got, want)
			}
		})
	}
}

func TestRepository_NewAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		dBDriver DBInterface
	}
	type args struct {
		value float32
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		makeMock func(mockCtrl *gomock.Controller) *fields
		wantAcc  *entity.Account
		wantErr  bool
	}{
		{
			name: "Testing for new Account",
			args: args{value: 10.0},
			makeMock: func(mockCtrl *gomock.Controller) *fields {
				dbMock := mocks.NewMockDBInterface(mockCtrl)
				dbMock.EXPECT().Create(float32(10.0)).Return(1, nil).Times(1)
				return &fields{dbMock}
			},
			wantAcc: &entity.Account{ID: 1, Balance: 10.0},
			wantErr: false,
		},
		{
			name: "Testing for fail new Account",
			args: args{value: 10.0},
			makeMock: func(mockCtrl *gomock.Controller) *fields {
				dbMock := mocks.NewMockDBInterface(mockCtrl)
				dbMock.EXPECT().Create(float32(10.0)).Return(0, errors.New("Failed account")).Times(1)
				return &fields{dbMock}
			},
			wantAcc: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.makeMock(mockCtrl)

			r := Repository{
				dBDriver: fields.dBDriver,
			}
			gotAcc, err := r.NewAccount(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.NewAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAcc, tt.wantAcc) {
				t.Errorf("Repository.NewAccount() = %v, want %v", gotAcc, tt.wantAcc)
			}
		})
	}
}

func TestRepository_UpdateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		dBDriver DBInterface
	}
	type args struct {
		acc entity.Account
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		makeMock func(mockCtrl *gomock.Controller) *fields
		wantErr  bool
	}{
		{
			name: "Testing for success update Account",
			args: args{entity.Account{ID: 1, Balance: 20}},
			makeMock: func(mockCtrl *gomock.Controller) *fields {
				dbMock := mocks.NewMockDBInterface(mockCtrl)
				dbMock.EXPECT().Update(1, float32(20.0)).Return(nil).Times(1)
				return &fields{dbMock}
			},
			wantErr: false,
		},
		{
			name: "Testing for success update Account",
			args: args{entity.Account{ID: 1, Balance: 20}},
			makeMock: func(mockCtrl *gomock.Controller) *fields {
				dbMock := mocks.NewMockDBInterface(mockCtrl)
				dbMock.EXPECT().Update(1, float32(20.0)).Return(errors.New("Failed to update")).Times(1)
				return &fields{dbMock}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.makeMock(mockCtrl)

			r := Repository{
				dBDriver: fields.dBDriver,
			}

			if err := r.UpdateAccount(tt.args.acc); (err != nil) != tt.wantErr {
				t.Errorf("Repository.UdpdateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		dBDriver DBInterface
	}
	type args struct {
		id int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *entity.Account
		makeMock func(mockCtrl *gomock.Controller) *fields
		wantErr  bool
	}{
		{
			name: "Testing for success get Account",
			args: args{1},
			makeMock: func(mockCtrl *gomock.Controller) *fields {
				dbMock := mocks.NewMockDBInterface(mockCtrl)
				dbMock.EXPECT().Get(1).Return(float32(30), nil).Times(1)
				return &fields{dbMock}
			},
			want:    &entity.Account{ID: 1, Balance: 30},
			wantErr: false,
		},
		{
			name: "Testing for fail to get Account",
			args: args{1},
			makeMock: func(mockCtrl *gomock.Controller) *fields {
				dbMock := mocks.NewMockDBInterface(mockCtrl)
				dbMock.EXPECT().Get(1).Return(float32(0), errors.New("Wrong account")).Times(1)
				return &fields{dbMock}
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing for inexistent Account",
			args: args{1},
			makeMock: func(mockCtrl *gomock.Controller) *fields {
				dbMock := mocks.NewMockDBInterface(mockCtrl)
				dbMock.EXPECT().Get(1).Return(float32(0), pgx.ErrNoRows).Times(1)
				return &fields{dbMock}
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.makeMock(mockCtrl)

			r := Repository{
				dBDriver: fields.dBDriver,
			}
			got, err := r.GetAccount(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
