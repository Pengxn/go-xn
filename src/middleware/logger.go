package middleware

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/config"
)

// LoggerToFile is a custom logger middleware, it writes logs to os.Stdout and a file.
// And uses a custom log format function.
func LoggerToFile() gin.HandlerFunc {
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
	if config.Config.Logger.Route == "" {
		return os.Stdout
	}

	logFile := filepath.Join("logs", "route.log") // -> ./logs/route.log

	// Logging to a file, append logging if the file already exists.
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		slog.Error("open log file error", slog.Any("error", err), slog.String("logFile", logFile))
		return os.Stdout
	}

	return f
}
