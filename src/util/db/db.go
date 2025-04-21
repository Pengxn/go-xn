package db

import (
	"fmt"
	"log"
	"log/slog"
	"strings"

	_ "github.com/go-sql-driver/mysql" // MySQL/MariaDB driver
	_ "github.com/lib/pq"              // PostgreSQL driver
	_ "github.com/mattn/go-sqlite3"    // SQLite driver
	"xorm.io/xorm"

	"github.com/Pengxn/go-xn/src/config"
)

// getDBUrl returns database type and its DSN.
func getDBUrl() (dbType, dsn string) {
	db := config.Config.DB
	dbType = strings.ToLower(db.Type)

	switch dbType {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			db.User, db.Password, db.Url, db.Port, db.Name)
	case "postgresql":
		dbType = "postgres"
		dsn = fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
			db.Name, db.User, db.Password, db.Url, db.Port, db.SSLMode)
	case "sqlite3":
		dsn = "file:" + db.Name + "?cache=shared&mode=rwc"
	default:
		slog.Warn("unknown database type")
		// TODO: set default database settings
		dsn = ""
	}

	return
}

// DBEngine initializes database connection engine
// and returns *xorm.Engine.
func DBEngine() *xorm.Engine {
	orm, err := xorm.NewEngine(getDBUrl())
	if err != nil {
		log.Fatalln("can't connect your DB", err)
	}

	return orm
}
