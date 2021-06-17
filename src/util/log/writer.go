package log

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Pengxn/go-xn/src/util/file"
	"github.com/Pengxn/go-xn/src/util/home"
)

// writerLog writes log to the specified writer buffer.
// Example: os.Stderr, a file opened in write mode, a socket...
func writerLog() io.Writer {
	logFile, err := LogFilePath("fyj.log")
	if err != nil {
		log.Printf("Get log file Path %s error: %+v", logFile, err)
	}

	// Logging to a file, append logging if the file already exists.
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Open log file %s error: %+v", logFile, err)
		return os.Stderr
	}

	return io.MultiWriter(f, os.Stderr)
}

func LogFilePath(logFile string) (string, error) {
	if file.IsExist(logFile) && file.IsFile(logFile) {
		return logFile, nil // => ./<logFile>
	}

	logDir := "logs"
	if file.IsExist(logDir) && file.IsDir(logDir) {
		return filepath.Join(logDir, logFile), nil // => ./logs/<logFile>
	}

	logDir = filepath.Join(home.LogDir("fyj"), "logs")
	if !file.IsExist(logDir) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return logFile, err // => ./<logFile>
		}
	}
	return filepath.Join(logDir, logFile), nil // => ~/.local/share/fyj/logs/<logFile>
}
