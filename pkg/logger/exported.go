package logger


import (
	"log"
	"github.com/sirupsen/logrus"

	logrusAdapter "github.com/issenn/game-scraper/pkg/logger/adapter/logrus"
)


var (
	// _ StdLogger = &log.Logger{}
	_ StdLogger = (*log.Logger)(nil)

	_ StdLogger = (*logrus.Entry)(nil)
	_ StdLogger = (*logrus.Logger)(nil)

	_ StdLogger = (*logrusAdapter.Entry)(nil)
	_ StdLogger = (*logrusAdapter.Logger)(nil)
)

var (
	// std is the name of the standard logger in stdlib `log`
	std StdLogger
)

func StandardLogger() StdLogger {
	return std
}

// func WithField(key string, value interface{}) Logger {
// 	return std.WithField(key, value)
// }

// func WithFields(fields Fields) Logger {
// 	return std.WithFields(fields)
// }

// Trace logs a message at level Trace on the standard logger.
// func Trace(args ...interface{}) {
// 	std.Trace(args...)
// }

// Debug logs a message at level Debug on the standard logger.
// func Debug(args ...interface{}) {
// 	std.Debug(args...)
// }

// Info logs a message at level Info on the standard logger.
// func Info(args ...interface{}) {
// 	std.Info(args...)
// }

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	std.Print(args...)
}

// Warn logs a message at level Warn on the standard logger.
// func Warn(args ...interface{}) {
// 	std.Warn(args...)
// }

// Error logs a message at level Error on the standard logger.
// func Error(args ...interface{}) {
// 	std.Error(args...)
// }

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Tracef logs a message at level Trace on the standard logger.
// func Tracef(format string, args ...interface{}) {
// 	std.Tracef(format, args...)
// }

// Debugf logs a message at level Debug on the standard logger.
// func Debugf(format string, args ...interface{}) {
// 	std.Debugf(format, args...)
// }

// Infof logs a message at level Info on the standard logger.
// func Infof(format string, args ...interface{}) {
// 	std.Infof(format, args...)
// }

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
// func Warnf(format string, args ...interface{}) {
// 	std.Warnf(format, args...)
// }

// Errorf logs a message at level Error on the standard logger.
// func Errorf(format string, args ...interface{}) {
// 	std.Errorf(format, args...)
// }

// Panicf logs a message at level Panic on the standard logger.
// func Panicf(format string, args ...interface{}) {
// 	std.Panicf(format, args...)
// }

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
// func Fatalf(format string, args ...interface{}) {
// 	std.Fatalf(format, args...)
// }

// initialize the global standard logger.
func InitializeStdLogger(s StdLogger)  {
	std = s
}

func init() {
	InitializeStdLogger(logrus.StandardLogger())
	// InitializeStdLogger(logrusAdapter.StandardLogger())
	// InitializeStdLogger(logrusAdapter.NewLogrusAdapt(logrus.StandardLogger()))
}
