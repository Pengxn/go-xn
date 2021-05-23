package middleware

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/util/log"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: CustomLog,
		Output:    WriterLog(),
	})
}

// WriterLog writes log to the specified writer buffer.
// Example: os.Stdout, a file opened in write mode, a socket...
func WriterLog() io.Writer {
	// Logging to a file.
	f, err := os.Create("fyj.log")
	if err != nil {
		log.Errorln("Create log file error: %v", err)
	}

	return io.MultiWriter(f, os.Stdout)
}

// CustomLog is the custom log format function Logger middleware uses.
func CustomLog(param gin.LogFormatterParams) string {
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
