package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Handler struct {
	Client *sql.DB
}

const connStr string = "postgresql://%s:%s@%s/%s?sslmode=disable"

var (
	dbUser     string
	dbPassword string
	dbIp       string
	dbName     string
)

func CreateDbHandler() *Handler {
	err := loadEnvs()
	if err != nil {
		panic(err.Error())
	}

	s := fmt.Sprintf(connStr, dbUser, dbPassword, dbIp, dbName)

	db, err := sql.Open("postgres", s)

	if err != nil {
		panic(err.Error())
	}

	return &Handler{Client: db}
}

func loadEnvs() error {
	dbIp = os.Getenv("DB_IP")
	if dbIp == "" {
		return errors.New("env DB_IP not found")
	}

	dbName = os.Getenv("DB_NAME")
	if dbName == "" {
		return errors.New("env DB_NAME not found")
	}

	dbPassword = os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return errors.New("env DB_PASSWORD not found")
	}

	dbUser = os.Getenv("DB_USER")
	if dbUser == "" {
		return errors.New("env DB_USER not found")
	}

	return nil
}
