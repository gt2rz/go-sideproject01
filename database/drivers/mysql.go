package drivers

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

func NewMysqlConnection() *sql.DB {
	// Capture connection properties
	cfg := mysql.Config{
		User:   os.Getenv("DB_USERNAME"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}

	// ParseTime=true is required to parse DATE and DATETIME values
	cfg.ParseTime = true

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	// Ping to check the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
