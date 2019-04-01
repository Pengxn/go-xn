package util

import (
	"database/sql"

	// _ "github.com/go-sql-driver/mysql"
	"go-xn/src/config"
)

// InitDB initialize database connection
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBUrl())

	if err != nil {
		return nil, err
	}

	return db, err
}
