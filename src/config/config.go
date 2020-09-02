package config

import (
	"os"
	"strings"

	"github.com/go-ini/ini"

	"github.com/Pengxn/go-xn/src/util/home"
	"github.com/Pengxn/go-xn/src/util/log"
)

// DBConfiguration is custom configuration for DB
type DBConfiguration struct {
	Type     string `ini:"type"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Port     string `ini:"port"`
	Name     string `ini:"name"`
	Url      string `ini:"url"`
}

// getDBConfiguration returns database configuration.
func getDBConfiguration() *DBConfiguration {
	config, err := ini.Load(home.HomeDir() + string(os.PathSeparator) + "fyj.ini")
	if err != nil {
		log.Fatalln("Fail to read fyj.ini file.", err.Error())
	}

	database := new(DBConfiguration)

	err = config.Section("database").MapTo(database)
	if err != nil {
		log.Warnln("Fail to parse database configuration.", err.Error())
	}

	return database
}

// DBUrl returns database type and its DSN.
func DBUrl() (dbType, dsn string) {
	db := getDBConfiguration()
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
		}, " ")
	case "sqlite3":
		dsn = "file:" + db.Name + "?cache=shared&mode=rwc"
	default:
		dsn = ""
	}

	return
}
