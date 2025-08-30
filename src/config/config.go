package config

import (
	"log/slog"

	"github.com/go-ini/ini"
	"github.com/google/uuid"
)

var Config appConfig // Global config object

func init() {
	Config = appConfig{
		Server: ServerConfig{Port: "7991", JwtToken: uuid.New().String()},
		DB:     DBConfig{Type: "SQLite3", Name: "fyj.db", SSLMode: "disable"},
	}

	configPath := getConfigPathByFlag()
	if configPath == "" {
		configPath = "fyj.ini"
	}

	if err := loadConfig(configPath); err != nil {
		slog.Error("load config file failed", slog.Any("error", err))
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

// appConfig is configuration for global object.
type appConfig struct {
	Server   ServerConfig   `ini:"server"`
	DB       DBConfig       `ini:"database"`
	Redis    RedisConfig    `ini:"redis"`
	Sentry   SentryConfig   `ini:"sentry"`
	Logger   LoggerConfig   `ini:"log"`
	Otel     OtelConfig     `ini:"otel"`
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
	Username string `ini:"username"`
	Password string `ini:"password"`
	DB       int    `ini:"db"`
}

// SentryConfig is configuration for Sentry.
type SentryConfig struct {
	DSN   string `ini:"DSN"`
	Debug bool   `ini:"debug"`
}

// LoggerConfig is configuration for logger.
type LoggerConfig struct {
	Level    string `ini:"level"`
	Route    string `ini:"route"`
	APP      string `ini:"app"`      // log writer: bark, telegram, newrelic, default: os.Stdout
	Bark     string `ini:"bark"`     // bark token
	Telegram string `ini:"telegram"` // telegram bot token
	Newrelic string `ini:"newrelic"` // newrelic api key
}

// OtelConfig is the OpenTelemetry configuration.
type OtelConfig struct {
	EnableTrace      bool   `ini:"trace"`      // enable OpenTelemetry tracing
	EnableMetric     bool   `ini:"metric"`     // enable OpenTelemetry metrics
	EnableLog        bool   `ini:"log"`        // enable OpenTelemetry logging
	ClientType       string `ini:"clientType"` // client type, used to determine the OpenTelemetry client
	Endpoint         string `ini:"endpoint"`   // endpoint for OpenTelemetry, default: localhost:4317
	Header           string `ini:"header"`     // header name for authentication
	Token            string `ini:"token"`      // token or key for authentication
	EnablePrometheus bool   `ini:"prometheus"` // enable Prometheus metrics
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
