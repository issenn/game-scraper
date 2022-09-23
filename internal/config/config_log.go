package config


import (
	// "fmt"
	"reflect"

	// "github.com/mitchellh/mapstructure"
	// "github.com/spf13/viper"

	"github.com/issenn/game-scraper/pkg/logger"
)


type baseLogConfig struct {
}

type LogConfig struct {
	Logging          string          `mapstructure:"logging" json:"logging" yaml:"logging" toml:"logging"`
	Development      bool            `mapstructure:"development" json:"development" yaml:"development" toml:"development"`
	Dir              string          `mapstructure:"dir" json:"dir" yaml:"dir" toml:"dir"`

	// logger.BaseLoggingConfig

	logger.Config    `mapstructure:",squash"`
}

var defaultLogConf = LogConfig{
	Logging: "zap",
	// baseLogConfig: baseLogConfig{
	// 	Logging: "zap",
	// },
	// BaseLoggingConfig: logger.BaseLoggingConfig{
	// 	Level: logger.NewAtomicLevel(),
	// },
}

// NewConfig creates a new config instance with default value.
func NewLogConfig() *LogConfig {
	logConf := defaultLogConf
	return &logConf
}

// LogEnabled uses to check whether enabled the table lock feature.
func LogEnabled() bool {
	return !GetGlobalConfig().Log.IsEmpty()
}

func (l LogConfig) IsEmpty() bool {
	return reflect.ValueOf(l).IsZero()
}
