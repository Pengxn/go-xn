package slogger

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

// NewFilesWriter creates a new rolling files writer.
// It rotates the log file when it reaches a certain size, age...
func NewFilesWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // megabytes
		MaxAge:     30, // days
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   true,
	}
}
