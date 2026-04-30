package slogger

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// NewFilesWriter creates a new rolling files writer.
// It rotates the log file when it reaches a certain size, age...
func NewFilesWriter(filename string) io.Writer {
	finfo, err := os.Stat(filename)
	if err != nil && !os.IsNotExist(err) {
		slog.Error("failed to stat log file", slog.String("filename", filename), slog.Any("error", err))
		return io.Discard // or use `os.Stdout` instead
	}

	// TODO: check the file mode and permissions
	if finfo.IsDir() {
		slog.Error("log file is a directory", slog.String("filename", filename))
		return io.Discard
	}

	// TODO: customize lumberjack settings
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // megabytes
		MaxAge:     30, // days
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   true,
	}
}
