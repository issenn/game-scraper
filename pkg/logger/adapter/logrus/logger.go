package logrus

import (
	"github.com/sirupsen/logrus"
)


// Logger is a Logger adapter for Logrus.
type Logger struct {
	logger *logrus.Logger
}


func NewLogrusAdapt(l *logrus.Logger) *Logger {
	return &Logger{
		logger: l,
	}
}

func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return NewEntryAdapt(logger.logger.WithField(key, value))
}

func (logger *Logger) WithFields(fields Fields) *Entry {
	return NewEntryAdapt(logger.logger.WithFields(logrus.Fields(fields)))
}

func (logger *Logger) Trace(args ...interface{}) {
	logger.logger.Trace(args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.logger.Debug(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.logger.Info(args...)
}

func (logger *Logger) Print(args ...interface{}) {
	logger.logger.Print(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.logger.Warn(args...)
}

func (logger *Logger) Warning(args ...interface{}) {
	logger.logger.Warning(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.logger.Error(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.logger.Fatal(args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.logger.Panic(args...)
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	logger.logger.Tracef(format, args...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.logger.Debugf(format, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.logger.Infof(format, args...)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	logger.logger.Printf(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.logger.Warnf(format, args...)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	logger.logger.Warningf(format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.logger.Errorf(format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.logger.Fatalf(format, args...)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.logger.Panicf(format, args...)
}

func (logger *Logger) Traceln(args ...interface{}) {
	logger.logger.Traceln(args...)
}

func (logger *Logger) Debugln(args ...interface{}) {
	logger.logger.Debugln(args...)
}

func (logger *Logger) Infoln(args ...interface{}) {
	logger.logger.Infoln(args...)
}

func (logger *Logger) Println(args ...interface{}) {
	logger.logger.Println(args...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	logger.logger.Warnln(args...)
}

func (logger *Logger) Warningln(args ...interface{}) {
	logger.logger.Warningln(args...)
}

func (logger *Logger) Errorln(args ...interface{}) {
	logger.logger.Errorln(args...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	logger.logger.Fatalln(args...)
}

func (logger *Logger) Panicln(args ...interface{}) {
	logger.logger.Panicln(args...)
}
