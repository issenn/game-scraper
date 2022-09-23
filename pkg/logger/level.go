package logger


import (
	"bytes"
	"errors"
	"fmt"
	"encoding"
	// "reflect"

	// "golang.org/x/exp/constraints"

	"go.uber.org/atomic"
	// "go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
	// "github.com/sirupsen/logrus"

	zapAdapter "github.com/issenn/game-scraper/pkg/logger/adapter/zap"
	// logrusAdapter "github.com/issenn/game-scraper/pkg/logger/adapter/logrus"
)


var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel

	InvalidLevel = _maxLevel + 1
)

func ParseLevel(text string) (Level, error) {
	var level Level
	err := level.UnmarshalText([]byte(text))
	return level, err
}

// String converts a Level to string.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case DPanicLevel:
		return "dpanic"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug", "DEBUG":
		*l = DebugLevel
	case "info", "INFO", "": // make the zero value useful
		*l = InfoLevel
	case "warn", "WARN":
		*l = WarnLevel
	case "error", "ERROR":
		*l = ErrorLevel
	case "dpanic", "DPANIC":
		*l = DPanicLevel
	case "panic", "PANIC":
		*l = PanicLevel
	case "fatal", "FATAL":
		*l = FatalLevel
	default:
		return false
	}
	return true
}

func (l Level) Enabled(lvl Level) bool {
	return lvl >= l
}


type LevelAdapter interface {
	zapAdapter.Level
}

type Leveler interface {
	fmt.Stringer
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

type AtomicLevel struct {
	l *atomic.Int32
}

// var (
// 	_ AdapterLevel = (*zapcore.Level)(nil)
// )

func NewAtomicLevel() AtomicLevel {
	return AtomicLevel{
		l: atomic.NewInt32(int32(InfoLevel)),
	}
}

func NewAtomicLevelAt(l Level) AtomicLevel {
	a := NewAtomicLevel()
	a.SetLevel(l)
	return a
}

func ParseAtomicLevel(text string) (AtomicLevel, error) {
	a := NewAtomicLevel()
	l, err := ParseLevel(text)
	if err != nil {
		return a, err
	}

	a.SetLevel(l)
	return a, nil
}

func (lvl AtomicLevel) Enabled(l Level) bool {
	return lvl.Level().Enabled(l)
}

// Level returns the minimum enabled log level.
func (lvl AtomicLevel) Level() Level {
	return Level(int8(lvl.l.Load()))
}

// SetLevel alters the logging level.
func (lvl AtomicLevel) SetLevel(l Level) {
	lvl.l.Store(int32(l))
}

// String returns the string representation of the underlying Level.
func (lvl AtomicLevel) String() string {
	return lvl.Level().String()
}

func (lvl *AtomicLevel) UnmarshalText(text []byte) error {
	if lvl.l == nil {
		lvl.l = &atomic.Int32{}
	}

	var l Level
	if err := l.UnmarshalText(text); err != nil {
		return err
	}

	lvl.SetLevel(l)
	return nil
}

func (lvl AtomicLevel) MarshalText() (text []byte, err error) {
	return lvl.Level().MarshalText()
}

type ComparableFunc func(x, y any) bool

// func (f ComparableFunc) Enabled(c any) bool { return f(c) }

func (f *ComparableFunc) UnmarshalText(text []byte) error {
	if !f.unmarshalText(text) && !f.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized comparison operators: %q", text)
	}
	return nil
}

func (f *ComparableFunc) unmarshalText(text []byte) bool {
	switch string(text) {
	case "eq", "==", "=":
		*f = func(x, y any) bool { return x == y }
	case "ne", "!=":
		*f = func(x, y any) bool { return x != y }
	// case "gt", ">":
	// 	*f = func(x, y any) bool { return x > y }
	// case "ge", ">=", "":  // make the zero value useful
	// 	*f = func(x, y any) bool { return x >= y }
	// case "lt", "<":
	// 	*f = func(x, y any) bool { return x < y }
	// case "le", "<=":
	// 	*f = func(x, y any) bool { return x <= y }
	default:
		return false
	}
	return true
}

// func Equal[T constraints.Ordered](x, y T) bool {
// 	return x == y
// }

// func NotEqual(x any, y any) bool {
// 	return x != y
// }

type levelEnabled struct {}

type Comparable[T any] interface {
	Equal(rhs T) bool
}

type Ordered[T any] interface {
	Comparable[T]
	Less(rhs T) bool
}

