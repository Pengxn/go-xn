package db

import (
	_ "github.com/go-sql-driver/mysql" // MySQL/MariaDB driver
	_ "github.com/lib/pq"              // PostgreSQL driver
	_ "github.com/mattn/go-sqlite3"    // SQLite driver
	"xorm.io/xorm"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/util/log"
)

// DBEngine initializes database connection engine
// and returns *xorm.Engine.
func DBEngine() *xorm.Engine {
	orm, err := xorm.NewEngine(config.DBUrl())
	if err != nil {
		log.Fatalln("Can't connect your DB.", err)
	}

	return orm
}
