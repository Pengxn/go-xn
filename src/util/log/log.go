// log package is a wrapper around the [logrus] logger.
// It provides a simple interface for logging messages at different levels.
// It also provides functions for logging messages with formatting options.
//
// Deprecated: this package is deprecated and will be removed in the next release.
// Use the [log] and [log/slog] package instead for logging.
//
// [logrus]: https://github.com/sirupsen/logrus
package log

import (
	"github.com/sirupsen/logrus"

	"github.com/Pengxn/go-xn/src/util/log/hook"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

	logger.SetLevel(logrus.InfoLevel)
	logger.AddHook(hook.NewWriterHook(writerLog()))
}

// Error logs a message at level Error.
//
// Deprecated: it's recommended to use the [log/slog.Error] or [log/slog.ErrorContext] package instead.
func Error(args ...any) {
	logger.Error(args...)
}

// Errorf logs a message at level Error.
//
// Deprecated: it's recommended to use the [log/slog.Error] or [log/slog.ErrorContext] instead.
func Errorf(format string, args ...any) {
	logger.Errorf(format, args...)
}

// Errorln logs a message at level Error.
//
// Deprecated: it's recommended to use the [log/slog.Error] or [log/slog.ErrorContext] instead.
func Errorln(format string, args ...any) {
	logger.Errorln(args...)
}

// Info logs a message at level Info.
//
// Deprecated: it's recommended to use the [log/slog.Info] or [log/slog.InfoContext] instead.
func Info(args ...any) {
	logger.Info(args...)
}

// Infof logs a message at level Info.
//
// Deprecated: it's recommended to use the [log/slog.Info] or [log/slog.InfoContext] instead.
func Infof(format string, args ...any) {
	logger.Infof(format, args...)
}

// Infoln logs a message at level Info.
//
// Deprecated: it's recommended to use the [log/slog.Info] or [log/slog.InfoContext] instead.
func Infoln(args ...any) {
	logger.Infoln(args...)
}

// Debug logs a message at level Debug.
//
// Deprecated: it's recommended to use the [log/slog.Debug] or [log/slog.DebugContext] instead.
func Debug(args ...any) {
	logger.Debug(args...)
}

// Debugf logs a message at level Debug.
//
// Deprecated: it's recommended to use the [log/slog.Debug] or [log/slog.DebugContext] instead.
func Debugf(format string, args ...any) {
	logger.Debugf(format, args...)
}

// Debugln logs a message at level Debug.
//
// Deprecated: it's recommended to use the [log/slog.Debug] or [log/slog.DebugContext] instead.
func Debugln(args ...any) {
	logger.Debugln(args...)
}

// Fatal logs a message at level Fatal and then
// the process will exit with status set to 1.
//
// Deprecated: it's recommended to use the [log.Fatal] instead.
func Fatal(args ...any) {
	logger.Fatal(args...)
}

// Fatalf logs a message at level Fatal and then
// the process will exit with status set to 1.
//
// Deprecated: it's recommended to use the [log.Fatalf] instead.
func Fatalf(format string, args ...any) {
	logger.Fatalf(format, args...)
}

// Fatalln logs a message at level Fatal and then
// the process will exit with status set to 1.
//
// Deprecated: it's recommended to use the [log.Fatalln] instead.
func Fatalln(args ...any) {
	logger.Fatalln(args...)
}

// Warn logs a message at level Warn.
//
// Deprecated: it's recommended to use the [log/slog.Warn] or [log/slog.WarnContext] instead.
func Warn(args ...any) {
	logger.Warn(args...)
}

// Warnf logs a message at level Warn.
//
// Deprecated: it's recommended to use the [log/slog.Warn] or [log/slog.WarnContext] instead.
func Warnf(format string, args ...any) {
	logger.Warnf(format, args...)
}

// Warnln logs a message at level Warn.
//
// Deprecated: it's recommended to use the [log/slog.Warn] or [log/slog.WarnContext] instead.
func Warnln(args ...any) {
	logger.Warnln(args...)
}

// Panic logs a message at level Panic.
//
// Deprecated: it's recommended to use the [log.Panic] instead.
func Panic(args ...any) {
	logger.Panic(args...)
}

// Panicf logs a message at level Panic.
//
// Deprecated: it's recommended to use the [log.Panicf] instead.
func Panicf(format string, args ...any) {
	logger.Panicf(format, args...)
}

// Panicln logs a message at level Panic.
//
// Deprecated: it's recommended to use the [log.Panicln] instead.
func Panicln(format string, args ...any) {
	logger.Panicln(args...)
}
