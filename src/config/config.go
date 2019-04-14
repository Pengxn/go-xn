package config

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

// DBConfiguration is your custom configuration for this app
type DBConfiguration struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Port     string `ini:"port"`
	DBName   string `ini:"name"`
	DBUrl    string `ini:"url"`
}

// GetDBConfiguration will return database configuration
func GetDBConfiguration() *DBConfiguration {
	config, err := ini.Load("fyj.ini")

	if err != nil {
		fmt.Printf("Fail to read fyj.ini file: %v", err)

		os.Exit(1)
	}

	database := new(DBConfiguration)

	err = config.Section("database").MapTo(database)

	return database
}

// DBUrl will return database url
func DBUrl() string {
	db := GetDBConfiguration()

	return db.User + ":" + db.Password + "@tcp(" + db.DBUrl + ":" + db.Port + ")/" + db.DBName + "?charset=utf8&parseTime=true"
}
