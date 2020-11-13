package usecase

import (
	"cashmachine/pkg/entity"
	"cashmachine/pkg/mocks"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestWithdrawUsecase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		post RequestWithdraw
		ri   RepositoryInterface
	}
	tests := []struct {
		name     string
		makeArgs func(mockCtrl *gomock.Controller) args
		wantRw   ResponseWithdraw
		wantErr  bool
	}{
		{
			name: "Test for $30 Withdraw",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(22).Return(&entity.Account{ID: 22, Balance: 70}, nil).Times(1)
				repositoryMock.EXPECT().UpdateAccount(entity.Account{ID: 22, Balance: 40}).Return(nil).Times(1)
				rw := RequestWithdraw{AccID: 22, Value: 30}
				return args{ri: repositoryMock, post: rw}
			},
			wantRw:  ResponseWithdraw{Bills: []entity.Bill{{Value: 10}, {Value: 10}, {Value: 10}}},
			wantErr: false,
		},
		{
			name: "Test for $80 Withdraw",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(22).Return(&entity.Account{ID: 22, Balance: 100}, nil).Times(1)
				repositoryMock.EXPECT().UpdateAccount(entity.Account{ID: 22, Balance: 20}).Return(nil).Times(1)
				rw := RequestWithdraw{AccID: 22, Value: 80}
				return args{ri: repositoryMock, post: rw}
			},
			wantRw:  ResponseWithdraw{Bills: []entity.Bill{{Value: 50}, {Value: 10}, {Value: 10}, {Value: 10}}},
			wantErr: false,
		},
		{
			name: "Test for invalid account",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(22).Return(nil, errors.New("Invalid account")).Times(1)
				rw := RequestWithdraw{AccID: 22, Value: 80}
				return args{ri: repositoryMock, post: rw}
			},
			wantRw:  ResponseWithdraw{Error: errors.New("Invalid account")},
			wantErr: true,
		},
		{
			name: "Test for invalid balance",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(33).Return(&entity.Account{ID: 33, Balance: 10}, nil).Times(1)
				rw := RequestWithdraw{AccID: 33, Value: 80}
				return args{ri: repositoryMock, post: rw}
			},
			wantRw:  ResponseWithdraw{Error: errors.New("Insufficient funds")},
			wantErr: true,
		},
		{
			name: "Test for invalid Update",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(33).Return(&entity.Account{ID: 33, Balance: 60}, nil).Times(1)
				repositoryMock.EXPECT().UpdateAccount(entity.Account{ID: 33, Balance: 30}).Return(errors.New("Failed to execute withdraw")).Times(1)
				rw := RequestWithdraw{AccID: 33, Value: 30}
				return args{ri: repositoryMock, post: rw}
			},
			wantRw:  ResponseWithdraw{Error: errors.New("Failed to execute withdraw")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.makeArgs(mockCtrl)
			gotRw, err := WithdrawUsecase(args.post, args.ri)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithdrawUsecase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRw, tt.wantRw) {
				t.Errorf("WithdrawUsecase() = %v, want %v", gotRw, tt.wantRw)
			}
		})
	}
}

func TestDepositUsecase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		request RequestDeposit
		r       RepositoryInterface
	}
	tests := []struct {
		name     string
		args     args
		makeArgs func(mockCtrl *gomock.Controller) args
		want     ResponseGeneral
		wantErr  bool
	}{
		{
			name: "Test for $30 Deposit",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(22).Return(&entity.Account{ID: 22, Balance: 70}, nil).Times(1)
				repositoryMock.EXPECT().UpdateAccount(entity.Account{ID: 22, Balance: 100}).Return(nil).Times(1)
				rd := RequestDeposit{AccID: 22, Value: 30}
				return args{r: repositoryMock, request: rd}
			},
			want:    ResponseGeneral{Msg: "ok"},
			wantErr: false,
		},
		{
			name: "Test for $100 Deposit",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(22).Return(&entity.Account{ID: 22, Balance: 70}, nil).Times(1)
				repositoryMock.EXPECT().UpdateAccount(entity.Account{ID: 22, Balance: 170}).Return(nil).Times(1)
				rd := RequestDeposit{AccID: 22, Value: 100}
				return args{r: repositoryMock, request: rd}
			},
			want:    ResponseGeneral{Msg: "ok"},
			wantErr: false,
		},
		{
			name: "Test for invalid account",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(22).Return(nil, errors.New("invalid account")).Times(1)
				rd := RequestDeposit{AccID: 22, Value: 100}
				return args{r: repositoryMock, request: rd}
			},
			want:    ResponseGeneral{Error: errors.New("invalid account")},
			wantErr: true,
		},
		{
			name: "Test for fail Update",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(22).Return(&entity.Account{ID: 22, Balance: 70}, nil).Times(1)
				repositoryMock.EXPECT().UpdateAccount(entity.Account{ID: 22, Balance: 170}).Return(errors.New("Failed to execute deposit")).Times(1)
				rd := RequestDeposit{AccID: 22, Value: 100}
				return args{r: repositoryMock, request: rd}
			},
			want:    ResponseGeneral{Error: errors.New("Failed to execute deposit")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.makeArgs(mockCtrl)
			got, err := DepositUsecase(args.request, args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("DepositUsecase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DepositUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBalanceUsecase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		request RequestBalance
		r       RepositoryInterface
	}
	tests := []struct {
		name     string
		args     args
		makeArgs func(mockCtrl *gomock.Controller) args
		want     ResponseBalance
		wantErr  bool
	}{
		{
			name: "Test for valid account",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(22).Return(&entity.Account{ID: 22, Balance: 100}, nil).Times(1)
				rd := RequestBalance{AccID: 22}
				return args{r: repositoryMock, request: rd}
			},
			want:    ResponseBalance{Balance: 100, AccID: 22},
			wantErr: false,
		},
		{
			name: "Test for invalid account",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().GetAccount(22).Return(nil, errors.New("Invalid account")).Times(1)
				rd := RequestBalance{AccID: 22}
				return args{r: repositoryMock, request: rd}
			},
			want:    ResponseBalance{Error: errors.New("Invalid account")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.makeArgs(mockCtrl)

			got, err := BalanceUsecase(args.request, args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("BalanceUsecase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BalanceUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUsecase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type args struct {
		request RequestCreate
		r       RepositoryInterface
	}
	tests := []struct {
		name     string
		args     args
		makeArgs func(mockCtrl *gomock.Controller) args
		want     ResponseBalance
		wantErr  bool
	}{
		{
			name: "Test for new account",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().NewAccount(float32(120)).Return(&entity.Account{ID: 22, Balance: 120}, nil).Times(1)
				rd := RequestCreate{Value: 120}
				return args{r: repositoryMock, request: rd}
			},
			want:    ResponseBalance{Balance: 120, AccID: 22},
			wantErr: false,
		},
		{
			name: "Test for fail on create account",
			makeArgs: func(mockCtrl *gomock.Controller) args {
				repositoryMock := mocks.NewMockRepositoryInterface(mockCtrl)
				repositoryMock.EXPECT().NewAccount(float32(120)).Return(nil, errors.New("Failed on create account")).Times(1)
				rd := RequestCreate{Value: 120}
				return args{r: repositoryMock, request: rd}
			},
			want:    ResponseBalance{Error: errors.New("Failed on create account")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.makeArgs(mockCtrl)

			got, err := CreateUsecase(args.request, args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUsecase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
