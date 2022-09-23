package logrus

import (
	"github.com/sirupsen/logrus"
)


type Entry struct {
	entry *logrus.Entry
}


func NewEntryAdapt(e *logrus.Entry) *Entry {
	return &Entry{
		entry: e,
	}
}

func (entry *Entry) WithError(err error) *Entry {
	return NewEntryAdapt(entry.entry.WithError(err))
}

func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return NewEntryAdapt(entry.entry.WithField(key, value))
}

func (entry *Entry) WithFields(fields Fields) *Entry {
	return NewEntryAdapt(entry.entry.WithFields(logrus.Fields(fields)))
}

func (entry *Entry) Trace(args ...interface{}) {
	entry.entry.Trace(args...)
}

func (entry *Entry) Debug(args ...interface{}) {
	entry.entry.Debug(args...)
}

func (entry *Entry) Info(args ...interface{}) {
	entry.entry.Info(args...)
}

func (entry *Entry) Print(args ...interface{}) {
	entry.entry.Print(args...)
}

func (entry *Entry) Warn(args ...interface{}) {
	entry.entry.Warn(args...)
}

func (entry *Entry) Warning(args ...interface{}) {
	entry.entry.Warning(args...)
}

func (entry *Entry) Error(args ...interface{}) {
	entry.entry.Error(args...)
}

func (entry *Entry) Fatal(args ...interface{}) {
	entry.entry.Fatal(args...)
}

func (entry *Entry) Panic(args ...interface{}) {
	entry.entry.Panic(args...)
}

func (entry *Entry) Tracef(format string, args ...interface{}) {
	entry.entry.Tracef(format, args...)
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	entry.entry.Debugf(format, args...)
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	entry.entry.Infof(format, args...)
}

func (entry *Entry) Printf(format string, args ...interface{}) {
	entry.entry.Printf(format, args...)
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	entry.entry.Warnf(format, args...)
}

func (entry *Entry) Warningf(format string, args ...interface{}) {
	entry.entry.Warningf(format, args...)
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	entry.entry.Errorf(format, args...)
}

func (entry *Entry) Fatalf(format string, args ...interface{}) {
	entry.entry.Fatalf(format, args...)
}

func (entry *Entry) Panicf(format string, args ...interface{}) {
	entry.entry.Panicf(format, args...)
}

func (entry *Entry) Traceln(args ...interface{}) {
	entry.entry.Traceln(args...)
}

func (entry *Entry) Debugln(args ...interface{}) {
	entry.entry.Debugln(args...)
}

func (entry *Entry) Infoln(args ...interface{}) {
	entry.entry.Infoln(args...)
}

func (entry *Entry) Println(args ...interface{}) {
	entry.entry.Println(args...)
}

func (entry *Entry) Warnln(args ...interface{}) {
	entry.entry.Warnln(args...)
}

func (entry *Entry) Warningln(args ...interface{}) {
	entry.entry.Warningln(args...)
}

func (entry *Entry) Errorln(args ...interface{}) {
	entry.entry.Errorln(args...)
}

func (entry *Entry) Fatalln(args ...interface{}) {
	entry.entry.Fatalln(args...)
}

func (entry *Entry) Panicln(args ...interface{}) {
	entry.entry.Panicln(args...)
}
