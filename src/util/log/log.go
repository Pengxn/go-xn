package log

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/Pengxn/go-xn/src/util/log/hook"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

	logger.SetLevel(logrus.DebugLevel)
	logger.AddHook(hook.NewWriterHook(writerLog()))
}

// level returns specified log level, default is Info level.
func level(level string) logrus.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "FATAL":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}

// Error logs a message at level Error.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf logs a message at level Error.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Errorln logs a message at level Error.
func Errorln(format string, args ...interface{}) {
	logger.Errorln(args...)
}

// Info logs a message at level Info.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof logs a message at level Info.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Infoln logs a message at level Info.
func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

// Debug logs a message at level Debug.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf logs a message at level Debug.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Debugln logs a message at level Debug.
func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

// Fatal logs a message at level Fatal and then
// the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf logs a message at level Fatal and then
// the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Fatalln logs a message at level Fatal and then
// the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

// Warn logs a message at level Warn.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf logs a message at level Warn.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Warnln logs a message at level Warn.
func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

// Panic logs a message at level Panic.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf logs a message at level Panic.
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Panicln logs a message at level Panic.
func Panicln(format string, args ...interface{}) {
	logger.Panicln(args...)
}
