package v1


import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync/atomic"
	// "reflect"

	"gopkg.in/yaml.v3"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	"github.com/issenn/game-scraper/internal/meta"
	"github.com/issenn/game-scraper/internal/utils"
	viperOption "github.com/issenn/game-scraper/pkg/viper"
	mapstructureHooks "github.com/issenn/game-scraper/pkg/mapstructure"
)


// Config holds any kind of configuration options that comes from the outside world and
// is necessary for running the application.
type Config struct {
	// Log configuration
	Log     LogConfig       `mapstructure:"log" json:"log" yaml:"log" toml:"log"`
}

var defaultConf = Config{
	Log: defaultLogConf,
}

var (
	globalConf atomic.Value
)

// NewConfig creates a new config instance with default value.
func NewConfig() *Config {
	conf := defaultConf
	return &conf
}

// GetGlobalConfig returns the global configuration for this server.
// It should store configuration from command line and configuration file.
// Other parts of the system can read the global configuration use this function.
func GetGlobalConfig() *Config {
	return globalConf.Load().(*Config)
}

// StoreGlobalConfig stores a new config to the globalConf. It mostly uses in the test to avoid some data races.
func StoreGlobalConfig(config *Config) {
	globalConf.Store(config)
}

// removedConfig contains items that are no longer supported.
// they might still be in the config struct to support import,
// but are not actively used.
var removedConfig = map[string]struct{}{
	"example.key":          {},
}

// isAllRemovedConfigItems returns true if all the items that couldn't validate
// belong to the list of removedConfig items.
func isAllRemovedConfigItems(items []string) bool {
	for _, item := range items {
		if _, ok := removedConfig[item]; !ok {
			return false
		}
	}
	return true
}

// InitializeConfig initialize the global config handler.
func InitializeConfig(configFile string) {
	cfg := GetGlobalConfig()
	// fmt.Printf("Initialize config: %#v\n\n", cfg)
	cfg.Load(configFile)
	fmt.Println("[InitializeConfig] Using Logging Framework:", cfg.Log.Logging)
	// fmt.Println("[InitializeConfig] Using Log Configuration:", cfg.Log)
	// fmt.Print("[InitializeConfig] Using Log Configuration:")
	// utils.PrettyPrint(cfg.Log)
	StoreGlobalConfig(cfg)
}

// RemovedVariableCheck checks if the config file contains any items
// which have been removed. These will not take effect any more.
func (c *Config) RemovedVariableCheck(configFile string) error {
	// Todo
	return nil
}

// UpdateGlobal updates the global config, and provide a restore function that can be used to restore to the original.
func UpdateGlobal(f func(conf *Config)) {
	g := GetGlobalConfig()
	newConf := *g
	f(&newConf)
	StoreGlobalConfig(&newConf)
}

// RestoreFunc gets a function that restore the config to the current value.
func RestoreFunc() (restore func()) {
	g := GetGlobalConfig()
	return func() {
		StoreGlobalConfig(g)
	}
}

func init() {
	conf := defaultConf
	fmt.Printf("[init] default config: %#v\n\n", conf)
	StoreGlobalConfig(&conf)
}

// Load loads config options from a config file.
func (c *Config) Load(configFile string) (err error) {
	// c.LoadConfigFile(configFile)
	if ok, _ := utils.PathExists(configFile); ok {
		c.LoadViper(configFile, "", "", "")
	} else {
		c.LoadViper(configFile, meta.Meta.AppName, "log", "")
	}
	return err
}

func (c *Config) LoadViper(configFile string, configDir string,
			   configName string, configType string) (err error) {
	// Don't forget to read config either from configFile or from default directory!
	if ok, _ := utils.PathExists(configFile); ok {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath("$NECRO_CONFIG_DIR")
		paths := utils.LookupConfigPath(configDir)
		for _, path := range paths {
			viper.AddConfigPath(path)
		}

		viper.SetConfigName(configName)
		if configType != "" {
			// REQUIRED if the config file does not have the extension in the name
			viper.SetConfigType(configType)
		}
	}

	// viper.AutomaticEnv()  // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		_, configFileNotFound := err.(viper.ConfigFileNotFoundError)
		if !configFileNotFound {
			fmt.Println("Failed to read configuration:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("[Load config] Using config file:", viper.ConfigFileUsed())
		fmt.Println("[Load config] Using Logging Framework:", viper.GetString("log.Logging"))
		// fmt.Println(viper.GetStringMapString("log"))
		// v := viper.Get("log.logger")
		// fmt.Printf("  %T\n  %+v\n", v, v)
		// utils.PrettyPrint(viper.Get("log"))
		fmt.Println(viper.AllSettings())
		// fmt.Println(reflect.ValueOf(c).Elem().Type())
		// fmt.Println(reflect.Zero(reflect.ValueOf(c).Elem().Type()))
		// fmt.Printf("\n%+v\n\n", *c)

		var metadata mapstructure.Metadata
		var metadataOption = viperOption.Metadata(&metadata)
		// var metadataOption = viperOption.WithOption("Metadata", &metadata)

		var squashOption = viperOption.Squash(true)
		// var squashOption = viperOption.WithOption("Squash", true)

		var decodeHookOption = viper.DecodeHook(
			mapstructure.ComposeDecodeHookFunc(
				// mapstructureHooks.TestUnmarshalerHookFunc(),
				// ZeroValueToDefaultStringHookFunc(viper.AllSettings()),
				// mapstructureHooks.UnmarshalerHookFunc(),
				// mapstructureHooks.CustomUnmarshalerHookFunc(),
				// mapstructure.StringToTimeDurationHookFunc(),
				// mapstructure.StringToSliceHookFunc(","),
				mapstructureHooks.TextUnmarshalerHookFunc(),
			),
		)

		if err := viper.Unmarshal(c, metadataOption, squashOption, decodeHookOption); err != nil {
			fmt.Println("Failed to unmarshal configuration:", err)
			os.Exit(1)
		}
		// fmt.Printf("Metadata: %+v\n\n", metadata)
		// utils.PrettyPrint(*c)
		// fmt.Println(*c)
	}
	return err
}

func (c *Config) LoadConfigFile(configFile string) (err error) {
	switch filepath.Ext(configFile) {
	case ".yaml", ".yml":
		// conf := Config{}
		if yamlFile, err := ioutil.ReadFile(configFile); err != nil {
			panic("Read conf error: " + err.Error())
		} else if err = yaml.Unmarshal(yamlFile, &c); err != nil {
			panic("Conf file unmarshal error: " + err.Error())
		}
	case ".toml":
		fmt.Println("toml")
		// Todo
	default:
		fmt.Println("default")
		// Todo
	}
	return err
}

// Process post-processes configuration after loading it.
func (c *Config) Process() error {
	// Todo
	return nil
}

// Valid checks if this config is valid.
func (c *Config) Valid() error {
	// Todo
	return nil
}

// hideConfig is used to filter a single line of config for hiding.
var hideConfig = []string{
	"example.key",
}

// jsonifyPath converts the item to json path, so it can be extracted.
func jsonifyPath(str string) string {
	s := strings.Split(str, ".")
	return fmt.Sprintf("$.\"%s\"", strings.Join(s, "\".\""))
}

// GetJSONConfig returns the config as JSON with hidden items removed
// It replaces the earlier HideConfig() which used strings.Split() in
// an way that didn't work for similarly named items (like enable).
func GetJSONConfig() (string, error) {
	// Todo
	return "", nil
}

// ContainHiddenConfig checks whether it contains the configuration that needs to be hidden.
func ContainHiddenConfig(s string) bool {
	s = strings.ToLower(s)
	for _, hc := range hideConfig {
		if strings.Contains(s, hc) {
			return true
		}
	}
	for dc := range removedConfig {
		if strings.Contains(s, dc) {
			return true
		}
	}
	return false
}
