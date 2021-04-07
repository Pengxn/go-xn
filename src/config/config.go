package config

import (
	"os"

	"github.com/go-ini/ini"

	"github.com/Pengxn/go-xn/src/util/home"
	"github.com/Pengxn/go-xn/src/util/log"
)

var config *ini.File // Global setting object

func init() {
	configFile, err := ini.Load(home.HomeDir() + string(os.PathSeparator) + "fyj.ini")
	if err != nil {
		log.Fatalln("Fail to read fyj.ini file.", err)
	}
	config = configFile
}

// DBConfiguration is custom configuration for DB.
type DBConfiguration struct {
	Type     string `ini:"type"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Port     string `ini:"port"`
	Name     string `ini:"name"`
	Url      string `ini:"url"`
}

// getDBConfiguration returns database configuration.
func GetDBConfig() *DBConfiguration {
	database := new(DBConfiguration)
	if err := config.Section("database").MapTo(database); err != nil {
		log.Warnln("Fail to parse database configuration.", err)
	}

	return database
}

// DNSConfiguration is DNS configuration for Tencent Cloud.
type DNSConfiguration struct {
	SecretID  string `ini:"secretID"`
	SecretKey string `ini:"secretKey"`
}

func DNSConfig() *DNSConfiguration {
	dns := new(DNSConfiguration)
	if err := config.Section("dns").MapTo(dns); err != nil {
		log.Warnln("Fail to parse DNS configuration.", err)
	}

	return dns
}
