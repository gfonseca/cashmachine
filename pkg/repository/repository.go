package repository

import (
	"cashmachine/pkg/entity"
	"fmt"
)

// DBInterface represents a connection with some dabase system
type DBInterface interface {
	Get(int) (float32, error)
	Update(int, float32) error
	Create(float32) (int, error)
}

// Repository is an data access interface
type Repository struct {
	dBDriver DBInterface
}

// NewRepository build and configure Repository instance
func NewRepository(dbdriver DBInterface) *Repository {
	r := &Repository{dBDriver: dbdriver}
	return r
}

//NewAccount create an database register to store account balance
func (r Repository) NewAccount(value float32) (acc *entity.Account, err error) {
	id, err := r.dBDriver.Create(value)

	if err != nil {
		err = fmt.Errorf("Failed to create account in database; %s", err)
		return
	}
	acc = &entity.Account{ID: id, Balance: value}
	return
}

//UpdateAccount persist Account state in database
func (r Repository) UpdateAccount(acc entity.Account) error {
	err := r.dBDriver.Update(acc.ID, acc.Balance)

	if err != nil {
		return fmt.Errorf("Failed to create account in database; %s", err)
	}
	return nil
}

//GetAccount get an Account from database
func (r Repository) GetAccount(id int) (*entity.Account, error) {
	balance, err := r.dBDriver.Get(id)
	if err != nil {
		return nil, err
	}

	return &entity.Account{ID: id, Balance: balance}, nil
}
