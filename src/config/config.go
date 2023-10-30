package config

import (
	"path/filepath"

	"github.com/go-ini/ini"
	"github.com/google/uuid"

	"github.com/Pengxn/go-xn/src/util/file"
	"github.com/Pengxn/go-xn/src/util/home"
	"github.com/Pengxn/go-xn/src/util/log"
)

var Config appConfig // Global config object

func init() {
	Config = appConfig{
		Server: ServerConfig{Port: "7991", JwtToken: uuid.New().String()},
		DB:     DBConfig{Type: "SQLite3", Name: "fyj.db", SSLMode: "disable"},
	}

	configPath := getConfigPathByFlag()
	if configPath == "" {
		configPath = defaultConfigPath()
	}

	if err := loadConfig(configPath); err != nil {
		log.Errorf("Load config file failed, %+v", err)
	}
}

func loadConfig(file string) error {
	configFile, err := ini.LooseLoad(file)
	if err != nil {
		return err
	}

	configFile.BlockMode = false // Only read the config file, not write.
	if err := configFile.MapTo(&Config); err != nil {
		return err
	}
	return nil
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

// appConfig is configuration for global object.
type appConfig struct {
	Server   ServerConfig   `ini:"server"`
	DB       DBConfig       `ini:"database"`
	Redis    RedisConfig    `ini:"redis"`
	DNS      DNSConfig      `ini:"dns"`
	Sentry   SentryConfig   `ini:"sentry"`
	Logger   LoggerConfig   `ini:"log"`
	WebAuthn WebAuthnConfig `ini:"webauthn"`
	SMTP     SMTPConfig     `ini:"smtp"`
}

// ServerConfig is configuration for server.
type ServerConfig struct {
	Debug    bool   `ini:"debug"`
	Port     string `ini:"port"`
	TLS      bool   `ini:"tls"`
	CertFile string `ini:"certFile"`
	KeyFile  string `ini:"keyFile"`
	JwtToken string `ini:"jwtToken"`
	JwtExp   int    `ini:"jwtExp"`
}

// DBConfig is custom configuration for DB.
type DBConfig struct {
	Type     string `ini:"type"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Port     string `ini:"port"`
	Name     string `ini:"name"`
	Url      string `ini:"url"`
	SSLMode  string `ini:"sslMode"`
}

// RedisConfig is custom configuration for Redis.
type RedisConfig struct {
	URL      string `ini:"url"`
	Password string `ini:"password"`
	DB       int    `ini:"db"`
}

// DNSConfig is DNS configuration for Tencent Cloud.
type DNSConfig struct {
	SecretID  string `ini:"secretID"`
	SecretKey string `ini:"secretKey"`
}

// SentryConfig is configuration for Sentry.
type SentryConfig struct {
	DSN   string `ini:"DSN"`
	Debug bool   `ini:"debug"`
}

// LoggerConfig is configuration for logger.
type LoggerConfig struct {
	Level string `ini:"level"`
	APP   string `ini:"app"`
	Route string `ini:"route"`
}

// WebAuthnConfig is the WebAuthn configuration.
type WebAuthnConfig struct {
	RPID          string   `ini:"rpID"`
	RPDisplayName string   `ini:"rpDisplayName"`
	RPOrigins     []string `ini:"rpOrigins"`
}

// SMTPConfig is the mail SMTP configuration.
type SMTPConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	SkipTLS  bool   `ini:"skipTLS"`
}
