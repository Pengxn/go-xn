package slogger

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/lmittmann/tint"

	"github.com/Pengxn/go-xn/src/config"
)

// override default logger with `tint` logger, default to [DEBUG] level.
func init() {
	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: "2006/01/02 - 15:04:05",
	}))

	// Set `tint` logger with colorized output as default logger.
	slog.SetDefault(logger)
}

// CtxKey is the key type for context.
type CtxKey int

const (
	CtxVersionKey CtxKey = iota // version key
)

// SetLogger sets the logger with the given config settings.
// The default logger is [tint] logger with colorized output.
//
// [tint]: https://github.com/lmittmann/tint
func SetLogger(ctx context.Context, c config.LoggerConfig) {
	if c.APP == "" {
		return
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: mapToLevel(c.Level),
		})).WithGroup("app").
		With("server", c.APP).
		With("os", runtime.GOOS).
		With("arch", runtime.GOARCH)

	// Set extra `version` fields for logger, if any.
	if version := ctx.Value(CtxVersionKey); version != nil {
		logger = logger.With("version", version)
	}

	slog.SetDefault(logger)
}

// mapToLevel maps string level to [log/slog.Level].
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
		slog.Warn("invalid log level, use default level: INFO")
		return slog.LevelInfo
	}
}
