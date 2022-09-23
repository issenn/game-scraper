package zap

import (
	"go.uber.org/zap"
)


// Logger is a Logger adapter for zap.
type Logger struct {
	logger *zap.Logger
}


func NewAdapter(l *zap.Logger) *Logger {
	return &Logger{
		logger: l,
	}
}

func New() *Logger {
	// Todo
	return nil
}

func NewNop() *Logger {
	// Todo
	return nil
}

func NewProduction() *Logger {
	// Todo
	return nil
}

func NewDevelopment() *Logger {
	// Todo
	return nil
}
