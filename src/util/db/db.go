package db

import (
	"strings"

	_ "github.com/go-sql-driver/mysql" // MySQL/MariaDB driver
	_ "github.com/lib/pq"              // PostgreSQL driver
	_ "github.com/mattn/go-sqlite3"    // SQLite driver
	"xorm.io/xorm"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/util/log"
)

// getDBUrl returns database type and its DSN.
func getDBUrl() (dbType, dsn string) {
	db := config.Config.DB
	dbType = strings.ToLower(db.Type)

	switch dbType {
	case "mysql":
		dsn = db.User + ":" + db.Password + "@tcp(" + db.Url +
			":" + db.Port + ")/" + db.Name + "?charset=utf8"
	case "postgresql":
		dbType = "postgres"
		dsn = strings.Join([]string{
			"dbname=" + db.Name,
			"user=" + db.User,
			"password=" + db.Password,
			"host=" + db.Url,
			"port=" + db.Port,
			"sslmode=" + db.SSLMode,
		}, " ")
	case "sqlite3":
		dsn = "file:" + db.Name + "?cache=shared&mode=rwc"
	default:
		dsn = ""
	}

	return
}

// DBEngine initializes database connection engine
// and returns *xorm.Engine.
func DBEngine() *xorm.Engine {
	orm, err := xorm.NewEngine(getDBUrl())
	if err != nil {
		log.Fatalln("Can't connect your DB.", err)
	}

	return orm
}
