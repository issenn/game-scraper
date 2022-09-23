package logger


// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// StdLogger is what your logger-enabled library should take, that way
// it'll accept a stdlib logger and a logrus logger. There's no standard
// interface, this is the closest we get, unfortunately.
type StdLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

// Logger represent common interface for logging function
type Logger interface {
	// WithField(key string, value interface{}) Logger
	// WithFields(fields Fields) Logger

	// Trace(args ...interface{})
	// Tracef(format string, args ...interface{})

	Debug(args ...interface{})
	// Debugf(format string, args ...interface{})

	Info(args ...interface{})
	// Infof(format string, args ...interface{})

	// Printf(format string, args ...interface{})

	Warn(args ...interface{})
	// Warnf(format string, args ...interface{})

	// Warningf(format string, args ...interface{})

	Error(args ...interface{})
	// Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	// Fatalf(format string, args ...interface{})

	Panic(args ...interface{})
	// Panicf(format string, args ...interface{})
}
