// Package logger represents a generic logging interface

package logger


import (
	"github.com/sirupsen/logrus"

	logrusAdapter "github.com/issenn/game-scraper/pkg/logger/adapter/logrus"
)


var (
	_ Logger = (*logrus.Entry)(nil)
	_ Logger = (*logrus.Logger)(nil)

	_ Logger = (*logrusAdapter.Entry)(nil)
	_ Logger = (*logrusAdapter.Logger)(nil)
)

// Log is a package level variable, every program should access logging function through "Log"
var Log Logger

// SetLogger is the setter for log variable, it should be the only way to assign value to log
func SetLogger(newLogger Logger) {
	Log = newLogger
}

func init() {
	// Todo
}
