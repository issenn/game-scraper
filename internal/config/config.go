package config


import (
	"fmt"
	"io/ioutil"
	"sync/atomic"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var err error

// Config contains configuration options.
type Config struct {
	Log	LogConfig	`mapstructure:"log" yaml:"log"`
}

type LogConfig struct {
	Dir	string		`mapstructure:"dir" yaml:"dir"`
	Level2	string		`mapstructure:"level" yaml:"level"`
}

var defaultConf = Config{}

var (
	globalConf atomic.Value
)

// NewConfig creates a new config instance with default value.
func NewConfig() *Config {
	conf := defaultConf
	return &conf
}

func GetGlobalConfig() *Config {
	return globalConf.Load().(*Config)
}

// StoreGlobalConfig stores a new config to the globalConf. It mostly uses in the test to avoid some data races.
func StoreGlobalConfig(config *Config) {
	globalConf.Store(config)
}

// InitializeConfig initialize the global config handler.

func InitializeConfig(configFile string) {
	cfg := GetGlobalConfig()
	fmt.Println("init config", cfg)
	// cfg.LoadConfigFile(configFile)
}

func init() {
	conf := defaultConf
	StoreGlobalConfig(&conf)
}

func (c *Config) LoadConfigFile(configFile string) error {
	switch filepath.Ext(configFile) {
	case ".yaml", ".yml":
		conf := Config{}
		if yamlFile, err := ioutil.ReadFile(configFile); err != nil {
			panic("Read conf error: " + err.Error())
		} else if err = yaml.Unmarshal(yamlFile, &conf); err != nil {
			panic("Conf file unmarshal error: " + err.Error())
		}
		StoreGlobalConfig(&conf)
	case ".toml":
		fmt.Println("toml")
	default:
		fmt.Println("default")
	}
	return err
}
