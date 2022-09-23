package v1


import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"

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

// func (l *LogConfig) UnmarshalCustom(data interface{}) (bool, error) {
// 	dataMap, ok := data.(map[string]interface{})
// 	if !ok {
// 		return false, fmt.Errorf("can't unmarshal %+v to an LogConfig type", data)
// 	}
// 	logging := dataMap["logging"]
// 	loggerslice, ok := dataMap["logger"].([]interface{})
// 	for _, logger := range loggerslice {
// 		loggerMap, ok := logger.(map[string]interface{})
// 		if !ok {
// 			continue
// 		}
// 		fmt.Println(loggerMap["name"], loggerMap["logging"])
// 	}
// 	fmt.Printf("================\n\n")
// 	fmt.Println(logging)
// 	return false, nil
// }

// StringToTimeHookFunc returns a DecodeHookFunc that converts
// strings to time.Time.
func ZeroValueToDefaultStringHookFunc(m map[string]interface{}) mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
			fmt.Printf("%+v", m)

		return "", nil
	}
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
