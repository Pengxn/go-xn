package config

import (
	"path/filepath"

	"github.com/go-ini/ini"

	"github.com/Pengxn/go-xn/src/util/file"
	"github.com/Pengxn/go-xn/src/util/home"
	"github.com/Pengxn/go-xn/src/util/log"
)

var configINI *ini.File // Global setting object

func init() {
	configPath := defaultConfigPath()
	configFile, err := ini.Load(configPath)
	if err != nil {
		log.Fatalf("Fail to read config file %s, %+v", configPath, err)
	}
	configINI = configFile
}

func defaultConfigPath() string {
	files := []string{
		"fyj.ini",
		filepath.Join(home.ConfigDir("fyj"), "fyj.ini"), // ~/.config/fyj/fyj.ini
		filepath.Join(home.HomeDir(), "fyj.ini"),        // ~/fyj.ini
	}
	for _, f := range files {
		if file.IsExist(f) && file.IsFile(f) {
			return f
		}
	}

	return files[0] // default is fyj.ini
}

// ServerConfig is configuration for server.
type ServerConfig struct {
	Debug    bool   `ini:"debug"`
	Port     string `ini:"port"`
	TLS      bool   `ini:"tls"`
	CertFile string `ini:"certFile"`
	KeyFile  string `ini:"keyFile"`
}

// GetServerConfig returns server configuration.
func GetServerConfig() ServerConfig {
	var server ServerConfig
	if err := configINI.Section("server").MapTo(&server); err != nil {
		log.Warnln("Fail to parse server configuration.", err)
	}

	return server
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
func GetDBConfig() DBConfig {
	var database DBConfig
	if err := configINI.Section("database").MapTo(&database); err != nil {
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
func GetDNSConfig() DNSConfig {
	var dns DNSConfig
	if err := configINI.Section("dns").MapTo(&dns); err != nil {
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
func GetSentryConfig() SentryConfig {
	var sentry SentryConfig
	if err := configINI.Section("sentry").MapTo(&sentry); err != nil {
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
func GetLoggerConfig() LoggerConfig {
	var logger LoggerConfig
	if err := configINI.Section("log").MapTo(&logger); err != nil {
		log.Warnln("Fail to parse logger configuration.", err)
	}

	return logger
}

// WebAuthnConfig is the WebAuthn configuration.
type WebAuthnConfig struct {
	RPID          string   `ini:"rpID"`
	RPDisplayName string   `ini:"rpDisplayName"`
	RPOrigins     []string `ini:"rpOrigins"`
}

// GetWebAuthnConfig returns WebAuthn configuration.
func GetWebAuthnConfig() WebAuthnConfig {
	var webAuthn WebAuthnConfig
	if err := configINI.Section("webauthn").MapTo(&webAuthn); err != nil {
		log.Warnln("Fail to parse WebAuthn configuration.", err)
	}

	return webAuthn
}

// SMTPConfig is the mail SMTP configuration.
type SMTPConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	SkipTLS  bool   `ini:"skipTLS"`
}

// GetSMTPConfig returns SMTP configuration.
func GetSMTPConfig() SMTPConfig {
	var smtp SMTPConfig
	if err := configINI.Section("smtp").MapTo(&smtp); err != nil {
		log.Warnln("Fail to parse SMTP configuration.", err)
	}

	return smtp
}
