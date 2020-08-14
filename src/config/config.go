package config

import (
	"log"
	"os"

	"github.com/go-ini/ini"
)

// DBConfiguration is custom configuration for DB
type DBConfiguration struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Port     string `ini:"port"`
	DBName   string `ini:"name"`
	DBUrl    string `ini:"url"`
}

// getDBConfiguration will return database configuration
func getDBConfiguration() *DBConfiguration {
	config, err := ini.Load(HomeDir() + string(os.PathSeparator) + "fyj.ini")
	if err != nil {
		log.Fatalln("Fail to read fyj.ini file.", err.Error())
		os.Exit(1)
	}

	database := new(DBConfiguration)

	err = config.Section("database").MapTo(database)
	if err != nil {
		log.Fatalln("Fail to parse database configuration.", err.Error())
	}

	return database
}

// DBUrl will return database url
func DBUrl() string {
	db := getDBConfiguration()

	return db.User + ":" + db.Password + "@tcp(" + db.DBUrl + ":" + db.Port + ")/" + db.DBName + "?charset=utf8"
}
