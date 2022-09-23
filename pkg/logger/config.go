package logger


import (
	// "fmt"
	// "reflect"

	// "github.com/mitchellh/mapstructure"

	// zapAdapter "github.com/issenn/game-scraper/pkg/logger/adapter/zap"
	// logrusAdapter "github.com/issenn/game-scraper/pkg/logger/adapter/logrus"
)

type BaseLoggingConfig struct {
	Level            AtomicLevel `mapstructure:"level" json:"level" yaml:"level" toml:"level"`
	LevelCmp         string          `mapstructure:"levelCmp" json:"levelCmp" yaml:"levelCmp" toml:"levelCmp"`

	// Log format, one of json or plain-text.
	Encoding         string          `mapstructure:"encoding" json:"encoding" yaml:"encoding" toml:"encoding"`
}

type baseLoggerConfig struct {
}

type baseCoreConfig struct {
}

// type LoggerConfigType interface {
// 	zapAdapter.LoggerConfig | logrusAdapter.LoggerConfig
// }

type Config struct {
	// BaseLoggingConfig

	Logger           []LoggerConfig  `mapstructure:"logger" json:"logger" yaml:"logger" toml:"logger"`
}

type LoggerConfig struct {
	Name             string          `mapstructure:"name" json:"name" yaml:"name" toml:"name"`
	Level            AtomicLevel     `mapstructure:"level" json:"level" yaml:"level" toml:"level"`
	Writer           []WriterConfig  `mapstructure:"writer" json:"writer" yaml:"writer" toml:"writer"`
}

type WriterConfig struct {
	Use           bool            `mapstructure:"use" json:"use" yaml:"use" toml:"use"`
}

// type LoggerConfig struct {
// 	Name             string          `mapstructure:"name" json:"name" yaml:"name" toml:"name"`

	// baseLoggerConfig

	// EncoderConfig    EncoderConfig   `mapstructure:"encoderConfig" json:"encoderConfig" yaml:"encoderConfig" toml:"encoderConfig"`
	// ErrorOutputPaths []string        `mapstructure:"errorOutputPaths" json:"errorOutputPaths" yaml:"errorOutputPaths"`
// }

// type CoreConfig struct {
// 	Level         AtomicLevel `mapstructure:"level" json:"level" yaml:"level" toml:"level"`
// 	LevelEnabled  string          `mapstructure:"levelEnabled" json:"levelEnabled" yaml:"levelEnabled" toml:"levelEnabled"`

// 	OutputPaths   []string        `mapstructure:"outputPaths" json:"outputPaths" yaml:"outputPaths" toml:"outputPaths"`

// 	Encoding      string          `mapstructure:"encoding" json:"encoding" yaml:"encoding" toml:"encoding"`
// 	EncoderConfig EncoderConfig   `mapstructure:"encoderConfig" json:"encoderConfig" yaml:"encoderConfig" toml:"encoderConfig"`
// }

// type EncoderConfig struct {
// 	MessageKey       string                  `mapstructure:"messageKey" json:"messageKey" yaml:"messageKey" toml:"messageKey"`
// 	LevelKey         string                  `mapstructure:"levelKey" json:"levelKey" yaml:"levelKey" toml:"levelKey"`
// 	TimeKey          string                  `mapstructure:"timeKey" json:"timeKey" yaml:"timeKey" toml:"timeKey"`

// 	EncodeLevel      zapcore.LevelEncoder    `mapstructure:"levelEncoder" json:"levelEncoder" yaml:"levelEncoder" toml:"levelEncoder"`
// 	EncodeTime       zapcore.TimeEncoder     `mapstructure:"timeEncoder" json:"timeEncoder" yaml:"timeEncoder" toml:"timeEncoder"`
// 	EncodeDuration   zapcore.DurationEncoder `mapstructure:"durationEncoder" json:"durationEncoder" yaml:"durationEncoder" toml:"durationEncoder"`
// 	EncodeCaller     zapcore.CallerEncoder   `mapstructure:"callerEncoder" json:"callerEncoder" yaml:"callerEncoder" toml:"callerEncoder"`

// 	EncodeName       zapcore.NameEncoder     `mapstructure:"nameEncoder" json:"nameEncoder" yaml:"nameEncoder" toml:"nameEncoder"`

// 	ConsoleSeparator string                  `mapstructure:"consoleSeparator" json:"consoleSeparator" yaml:"consoleSeparator" toml:"consoleSeparator"`
// }
