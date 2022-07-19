package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go-necro/internal/config"
)


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "necro",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Printf("necro: %s", VERSION)
			return
		}

		// configFile, _ := cmd.Flags().GetString("config-file")
		// fmt.Println(configFile)

		// config.InitializeConfig(configFile)
		// conf := config.GetGlobalConfig()
		// conf.LoadConfigFile("config/default.yaml")

		// fmt.Println("GetGlobalConfig:", config.GetGlobalConfig())

		// fmt.Println(viper.GetStringMapString("log"))
		// fmt.Println(viper.GetString("log.level"))

		if len(args) < 1 {
			cmd.Help()
			return
		}
	},
}

// rootCmd.CompletionOptions.DisableDefaultCmd = true

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func init() {
	var Verbose bool
	var Version bool
	var ConfigFile string

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&ConfigFile, "config-file", "c", "config/default.yaml", "config file")

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolVarP(&Version, "version", "V", false, "Version information")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.SetDefault("configFile", filepath.Join("config", "default.yaml"))

	// viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// viper.AutomaticEnv() // read in environment variables that match
	viper.BindEnv("XDG_CONFIG_HOME")
	viper.BindEnv("XDG_CACHE_HOME")
	viper.BindEnv("XDG_DATA_HOME")
}

func initConfig() {
	configFile, _ := rootCmd.Flags().GetString("config-file")

	// Don't forget to read config either from configFile or from default directory!
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find work directory.
		viper.AddConfigPath("config")

		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Find XDG Base Directory.
		// fmt.Println("XDG_CONFIG_HOME:", viper.GetString("XDG_CONFIG_HOME"))
		if viper.GetString("XDG_CONFIG_HOME") != "" {
			viper.AddConfigPath(filepath.Join(viper.GetString("XDG_CONFIG_HOME"), "necro"))
		} else {
			// Search config in home directory with name ".config" (without extension).
			viper.AddConfigPath(filepath.Join(home, ".config", "necro"))
			viper.AddConfigPath(filepath.Join(xdg.ConfigHome, "necro"))
		}

		viper.AddConfigPath(filepath.Join(home, ".necro"))

		viper.AddConfigPath("/etc/necro")

		viper.SetConfigName("default")
		// viper.SetConfigType("yaml")  // REQUIRED if the config file does not have the extension in the name
	}

	// viper.AutomaticEnv()  // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		conf := config.Config{}
		if err := viper.Unmarshal(&conf); err != nil {
			fmt.Println("Unable to decode into struct:", err)
			os.Exit(1)
		}
		config.StoreGlobalConfig(&conf)
	}
}
