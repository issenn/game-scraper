// Package logger represents a generic logging interface

package logger

// Log is a package level variable, every program should access logging function through "Log"
var Log Logger

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// Logger represent common interface for logging function
type Logger interface {
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
}

// SetLogger is the setter for log variable, it should be the only way to assign value to log
func SetLogger(newLogger Logger) {
	Log = newLogger
}
