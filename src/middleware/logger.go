package middleware

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/util/file"
	"github.com/Pengxn/go-xn/src/util/home"
	"github.com/Pengxn/go-xn/src/util/log"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: customLog,
		Output:    writerLog(),
	})
}

// customLog is the custom log format function Logger middleware uses.
func customLog(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("[FYJ] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// writerLog writes log to the specified writer buffer.
// Example: os.Stdout, a file opened in write mode, a socket...
func writerLog() io.Writer {
	logFile := "route.log"
	if !file.IsExist(logFile) {
		logDir := filepath.Join(home.LogDir("fyj"), "logs")
		if file.IsExist(logDir) {
			logFile = filepath.Join(logDir, logFile)
		} else {
			if err := os.MkdirAll(logDir, 0755); err != nil {
				log.Errorf("Mkdir folder %s error: %+v", logDir, err)
			} else {
				logFile = filepath.Join(logDir, logFile)
			}
		}
	}

	// Logging to a file, append logging if the file already exists.
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Errorf("Open log file %s error: %+v", logFile, err)
		return os.Stdout
	}

	return io.MultiWriter(f, os.Stdout)
}
