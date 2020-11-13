package utils

import (
	"cashmachine/pkg/driver"
	"cashmachine/pkg/repository"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// BuildRepository create and configure a repository instance
func BuildRepository() (*repository.Repository, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	database := os.Getenv("DB_DATABASE")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.WithField("error", err).Fatal("Failed to parse DB_PORT field from config")
		return nil, err
	}

	dbDriver, err := driver.NewPGDriver(host, port, user, pass, database)
	if err != nil {
		log.WithField("error", err).Fatal("Fail in database connection")
		return nil, err
	}

	return repository.NewRepository(dbDriver), nil
}
