package database

import (
	"database/sql"
	"errors"
	"fmt"
	"microtwo/database/drivers"
	"os"
)

var ErrNoDatabaseTypeSpecified = errors.New("no database type specified")

func GetDbSqlConnection() (*sql.DB, error) {
	var db *sql.DB

	switch os.Getenv("DB_TYPE") {
	case "mysql":
		db = drivers.NewMysqlConnection()

	case "postgres":
		db = drivers.NewPostgresConnection()
	default:
		fmt.Println("No database type specified")
		return nil, ErrNoDatabaseTypeSpecified
	}

	fmt.Println("Connected to database type: " + os.Getenv("DB_TYPE"))
	return db, nil
}
