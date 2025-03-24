package slogger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/lmittmann/tint"

	"github.com/Pengxn/go-xn/src/config"
)

// override default logger with `tint` logger.
func init() {
	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: "2006/01/02 - 15:04:05",
	}))

	// Set `tint` logger with colorized output as default logger.
	slog.SetDefault(logger)
}

// SetLogger sets the logger with the given config settings.
func SetLogger(c config.LoggerConfig) {
	if c.APP == "" {
		return
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: mapToLevel(c.Level),
	}))

	slog.SetDefault(logger)
}

// mapToLevel maps string level to slog.Level.
// If the level is not valid, default to 'INFO'.
func mapToLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
