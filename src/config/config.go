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

// DBConfig is custom configuration for DB.
type DBConfig struct {
	Type     string `ini:"type"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Port     string `ini:"port"`
	Name     string `ini:"name"`
	Url      string `ini:"url"`
}

// GetDBConfig returns database configuration.
func GetDBConfig() *DBConfig {
	database := new(DBConfig)
	if err := config.Section("database").MapTo(database); err != nil {
		log.Warnln("Fail to parse database configuration.", err)
	}

	return database
}

// DNSConfig is DNS configuration for Tencent Cloud.
type DNSConfig struct {
	SecretID  string `ini:"secretID"`
	SecretKey string `ini:"secretKey"`
}

// GetDNSConfig returns DNS configuration.
func GetDNSConfig() *DNSConfig {
	dns := new(DNSConfig)
	if err := config.Section("dns").MapTo(dns); err != nil {
		log.Warnln("Fail to parse DNS configuration.", err)
	}

	return dns
}

// SentryConfig is configuration for Sentry.
type SentryConfig struct {
	DSN   string `ini:"DSN"`
	Debug bool   `ini:"debug"`
}

// GetSentryConfig returns snetry configuration.
func GetSentryConfig() *SentryConfig {
	sentry := new(SentryConfig)
	if err := config.Section("sentry").MapTo(sentry); err != nil {
		log.Warnln("Fail to parse sentry configuration.", err)
	}

	return sentry
}

// LoggerConfig is configuration for logger.
type LoggerConfig struct {
	Route string `ini:"route"`
	APP   string `ini:"app"`
}

// GetLoggerConfig returns logger configuration.
func GetLoggerConfig() *LoggerConfig {
	logger := new(LoggerConfig)
	if err := config.Section("log").MapTo(logger); err != nil {
		log.Warnln("Fail to parse logger configuration.", err)
	}

	return logger
}
