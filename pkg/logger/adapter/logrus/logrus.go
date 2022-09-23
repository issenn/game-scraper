/**
 * @Title Logrus Adapter
 * @Description
 * @Author issenn
 * @Date 2022-07-21
 **/

// Package logrus provides a Logger adapter for Logrus.

package logrus

import (
	"github.com/sirupsen/logrus"
)


// Fields type, used to pass to `WithFields`.
// type Fields map[string]interface{}
type Fields logrus.Fields

// The FieldLogger interface generalizes the Entry and Logger types
// type FieldLogger interface {
// 	WithField(key string, value interface{}) *Entry
// 	WithFields(fields Fields) *Entry
// 	WithError(err error) *Entry

// 	Debug(args ...interface{})
// 	Info(args ...interface{})
// 	Print(args ...interface{})
// 	Warn(args ...interface{})
// 	Warning(args ...interface{})
// 	Error(args ...interface{})
// 	Fatal(args ...interface{})
// 	Panic(args ...interface{})

// 	Debugf(format string, args ...interface{})
// 	Infof(format string, args ...interface{})
// 	Printf(format string, args ...interface{})
// 	Warnf(format string, args ...interface{})
// 	Warningf(format string, args ...interface{})
// 	Errorf(format string, args ...interface{})
// 	Fatalf(format string, args ...interface{})
// 	Panicf(format string, args ...interface{})

// 	Debugln(args ...interface{})
// 	Infoln(args ...interface{})
// 	Println(args ...interface{})
// 	Warnln(args ...interface{})
// 	Warningln(args ...interface{})
// 	Errorln(args ...interface{})
// 	Fatalln(args ...interface{})
// 	Panicln(args ...interface{})
// }
type FieldLogger logrus.FieldLogger
