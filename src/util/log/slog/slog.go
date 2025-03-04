package slogger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func init() {
	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: "2006/01/02 - 15:04:05",
	}))

	slog.SetDefault(logger)
}
