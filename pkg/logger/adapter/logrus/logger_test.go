package logrus_test


import (
	"os"
	// "bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"

	. "github.com/issenn/game-scraper/pkg/logger/adapter/logrus"
)

func TestLogrusAdapter_New(t *testing.T) {
	// out := &bytes.Buffer{}
	l := logrus.New()
	l.SetReportCaller(true)
	// l.SetOutput(out)

	logger := NewLogrusAdapt(l)
	fmt.Println(reflect.ValueOf(logger).Type())
	logger.Trace("This is log message.")
	logger.Debug("This is log message.")
	logger.Info("This is log message.")
	logger.Print("This is log message.")
	logger.Warn("This is log message.")
	logger.Warning("This is log message.")
	logger.Error("This is log message.")
	logger.Fatal("This is log message.")
	logger.Panic("This is log message.")
}

func TestLogrusAdapter_NewLogger(t *testing.T) {
	var log = logrus.New()

	// out := &bytes.Buffer{}
	// log.SetOutput(out)
	log.SetOutput(os.Stdout)

	// log.SetFormatter(&logrus.JSONFormatter{})
	log.SetFormatter(&logrus.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
	})

	// log.SetLevel(logrus.WarnLevel)
	log.SetLevel(logrus.DebugLevel)

	// logrus.SetReportCaller(true)

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	// log.WithFields(logrus.Fields{
	// 	"omg":    true,
	// 	"number": 100,
	// }).Fatal("The ice breaks!")

	contextLogger := log.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other": "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
