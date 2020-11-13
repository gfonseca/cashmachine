package driver

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	pgxp "github.com/jackc/pgx/v4/pgxpool"
)

//PGDriver Represents a connection with postgres-sql instance
type PGDriver struct {
	con *pgxp.Pool
}

// ErrInvalidAccount Invalid account id
var ErrInvalidAccount = errors.New("Invalid account id")

var dbCon *PGDriver

func formatURL(host string, port int, user, pass, db string) string {
	return fmt.Sprintf("postgresql://%s:%d/%s?user=%s&password=%s", host, port, db, user, pass)
}

// NewPGDriver build and configure a new Postgres connection
func NewPGDriver(host string, port int, user, pass, db string) (*PGDriver, error) {
	if dbCon != nil {
		return dbCon, nil
	}

	url := formatURL(host, port, user, pass, db)

	conn, err := pgxp.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %s", err)
	}
	dbCon = &PGDriver{con: conn}
	return dbCon, nil
}

//Get Return the balance of a given account
func (pgd *PGDriver) Get(id int) (balance float32, err error) {
	q := "SELECT balance FROM accounts WHERE id=$1"
	err = pgd.con.QueryRow(context.Background(), q, id).Scan(&balance)
	if err == pgx.ErrNoRows {
		err = ErrInvalidAccount
	}
	return
}

//Update the balance of a giving account
func (pgd *PGDriver) Update(id int, value float32) (err error) {
	q := "UPDATE accounts SET balance=$1 WHERE id=$2"
	res, err := pgd.con.Exec(context.Background(), q, value, id)

	if cnt := res.RowsAffected(); cnt == 0 {
		err = ErrInvalidAccount
	}

	return
}

//Create a new account and return it's ID
func (pgd *PGDriver) Create(value float32) (id int, err error) {
	q := "INSERT INTO accounts (balance) VALUES ($1) RETURNING(id)"
	row := pgd.con.QueryRow(context.Background(), q, value)
	err = row.Scan(&id)
	return
}
