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
	logFile := "fyj.log"
	if !file.IsExist(logFile) {
		logDir := filepath.Join(home.LogDir("fyj"), "logs")
		if file.IsExist(logDir) {
			logFile = filepath.Join(logDir, logFile)
		} else {
			if err := os.MkdirAll(logDir, 0755); err != nil {
				log.Printf("Mkdir folder %s error: %+v", logDir, err)
			} else {
				logFile = filepath.Join(logDir, logFile)
			}
		}
	}

	// Logging to a file, append logging if the file already exists.
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Open log file %s error: %+v", logFile, err)
		return os.Stderr
	}

	return io.MultiWriter(f, os.Stderr)
}
