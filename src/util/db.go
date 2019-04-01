package util

import (
	"database/sql"

	"go-xn/src/config"
	// _ "github.com/go-sql-driver/mysql"
)

// InitDB initialize database connection
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBUrl())

	if err != nil {
		return nil, err
	}

	return db, err
}
