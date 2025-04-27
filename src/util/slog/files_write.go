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
	// TODO: check the file mode and permissions
	_, err := os.Stat(filename)
	if err != nil || os.IsNotExist(err) {
		slog.Error("file not exist", slog.String("filename", filename), slog.Any("error", err))
		return io.Discard // or instead of it use os.Stdout
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
