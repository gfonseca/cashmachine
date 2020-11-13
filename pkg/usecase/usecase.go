package usecase

import (
	"cashmachine/pkg/entity"
	"errors"
)

// RepositoryInterface represents an repository abstraction
type RepositoryInterface interface {
	NewAccount(float32) (*entity.Account, error)
	GetAccount(int) (*entity.Account, error)
	UpdateAccount(entity.Account) error
}

// RequestCreate Represents a user request for Create new account
type RequestCreate struct {
	Value float32 `json:"value,omitempty"`
}

// RequestBalance Represents a user request for balance
type RequestBalance struct {
	AccID int
}

// ResponseBalance is the output to the check balance operation
type ResponseBalance struct {
	Error   error   `json:"error,omitempty"`
	AccID   int     `json:"account_id,omitempty"`
	Balance float32 `json:"balance,omitempty"`
}

// RequestWithdraw Represents a user request for withdraw
type RequestWithdraw struct {
	Value int `json:"value,omitempty"`
	AccID int `json:"id,omitempty"`
}

// ResponseWithdraw is the output to the withdraw operation
type ResponseWithdraw struct {
	Error error         `json:"error,omitempty"`
	Bills []entity.Bill `json:"bills,omitempty"`
}

// RequestDeposit Represents a user request for deposit
type RequestDeposit struct {
	Value float32 `json:"value,omitempty"`
	AccID int     `json:"id,omitempty"`
}

// ResponseGeneral is the output to general operation
type ResponseGeneral struct {
	Error error  `json:"error,omitempty"`
	Msg   string `json:"msg,omitempty"`
}

// WithdrawUsecase withdraw some value from account
func WithdrawUsecase(request RequestWithdraw, r RepositoryInterface) (ResponseWithdraw, error) {
	Acc, err := r.GetAccount(request.AccID)
	if err != nil {
		return ResponseWithdraw{Error: err}, err
	}

	if !Acc.VerifyBalance(float32(request.Value)) {
		errMsg := errors.New("Insufficient funds")
		return ResponseWithdraw{Error: errMsg}, errMsg
	}

	bills, err := Acc.Withdraw(request.Value)
	err = r.UpdateAccount(*Acc)
	if err != nil {
		return ResponseWithdraw{Error: err}, err
	}

	return ResponseWithdraw{Bills: bills}, nil
}

// DepositUsecase depoisit some value in the user account
func DepositUsecase(request RequestDeposit, r RepositoryInterface) (ResponseGeneral, error) {
	Acc, err := r.GetAccount(request.AccID)
	if err != nil {
		return ResponseGeneral{Error: err}, err
	}
	Acc.Balance += float32(request.Value)

	err = r.UpdateAccount(*Acc)
	if err != nil {
		return ResponseGeneral{Error: err}, err
	}

	return ResponseGeneral{Msg: "ok"}, nil
}

// BalanceUsecase check account balance
func BalanceUsecase(request RequestBalance, r RepositoryInterface) (ResponseBalance, error) {
	Acc, err := r.GetAccount(request.AccID)
	if err != nil {
		return ResponseBalance{Error: err}, err
	}

	return ResponseBalance{Balance: Acc.Balance, AccID: Acc.ID}, nil
}

// CreateUsecase Create a new account
func CreateUsecase(request RequestCreate, r RepositoryInterface) (ResponseBalance, error) {
	Acc, err := r.NewAccount(float32(request.Value))
	if err != nil {
		return ResponseBalance{Error: err}, err
	}

	return ResponseBalance{AccID: Acc.ID, Balance: Acc.Balance}, nil
}
