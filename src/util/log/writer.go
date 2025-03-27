package log

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Pengxn/go-xn/src/util/file"
)

// writerLog writes log to the specified writer buffer.
// Example: os.Stderr, a file opened in write mode, a socket...
func writerLog() io.Writer {
	logFile, err := LogFilePath("fyj.log")
	if err != nil {
		log.Printf("get log file Path %s error: %v", logFile, err)
		return os.Stderr
	}

	// Logging to a file, append logging if the file already exists.
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("open log file %s error: %v", logFile, err)
		return os.Stderr
	}

	return f
}

// LogFilePath returns the log file path.
//
// Deprecated: use fixed log file name instead, with `./logs` directory.
// This function is deprecated and will be removed in the future.
func LogFilePath(fileName string) (string, error) {
	if file.IsExist(fileName) && file.IsFile(fileName) {
		return fileName, nil // => ./<logFile>
	}

	logDir := "logs"
	if !file.IsExist(logDir) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return "", err // => ./<logFile>
		}
	}

	return filepath.Join(logDir, fileName), nil // => ./logs/<logFile>
}
