package cmd

import (
	"fmt"
	"os"
	// "path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// "github.com/issenn/game-scraper/internal/meta"
	// "github.com/issenn/game-scraper/internal/util"
	"github.com/issenn/game-scraper/internal/config"
	"github.com/issenn/game-scraper/internal/logger"
)


const (
	// appName is an identifier-like name used anywhere this app needs to be identified.
	//
	// It identifies the application itself, the actual instance needs to be identified via environment
	// and other details.
	appName = "necro"

	// friendlyAppName is the visible name of the application.
	friendlyAppName = "Necro Go Application"
)


func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   appName,
		Short: "Show help for necro commands and flags.",
		Long: `
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
		Run: runRoot,
	}

	// meta.SetDefaultAppName(appName)

	// Here you will define your flags and configuration settings.

	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	cmd.PersistentFlags().String("config-dir", "config/", "Set config dir (default is /path/to/dir)")
	cmd.PersistentFlags().StringP("config-file", "c", "config/default.yaml", "Set configuration file (default is config/default.yaml)")
	cmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug output")
	cmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	cmd.PersistentFlags().BoolP("quiet", "q", false, "Do not output to stdout")
	// cmd.PersistentFlags().String("logfile", "", "Set logfile")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	cmd.Flags().BoolP("version", "V", false, "Show version information")
	// cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// viper.SetDefault("configFile", filepath.Join("config", "default.yaml"))

	_ = viper.BindPFlag("config-dir", cmd.PersistentFlags().Lookup("config-dir"))
	_ = viper.BindPFlag("config-file", cmd.PersistentFlags().Lookup("config-file"))
	_ = viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose"))
	_ = viper.BindPFlag("quiet", cmd.PersistentFlags().Lookup("quiet"))
	// _ = viper.BindPFlag("logfile", cmd.PersistentFlags().Lookup("logfile"))

	// viper.AutomaticEnv() // read in environment variables that match
	_ = viper.BindEnv("XDG_CONFIG_HOME")
	_ = viper.BindEnv("XDG_CACHE_HOME")
	_ = viper.BindEnv("XDG_DATA_HOME")
	return cmd
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = NewRootCmd()

// runRoot implements the main rclone command with no subcommands
func runRoot(cmd *cobra.Command, args []string) {
	versionFlag, _ := cmd.Flags().GetBool("version")
	if versionFlag {
		ShowVersion()
	} else {
		// _ = cmd.Usage()
		// if len(args) > 0 {
		// 	_, _ = fmt.Fprintf(os.Stderr, "Command not found.\n")
		// }

		// configFile, _ := cmd.Flags().GetString("config-file")
		// fmt.Println(configFile)

		// config.InitializeConfig(configFile)
		// conf := config.GetGlobalConfig()
		// conf.LoadConfigFile("config/default.yaml")

		fmt.Println("\n---runRoot---")
		fmt.Printf("\nGetGlobalConfig: %+v\n\n", config.GetGlobalConfig())
		// fmt.Printf("Log: %+v\n", config.GetGlobalConfig().Log)
		fmt.Println("LogEnabled:", config.LogEnabled())

		// fmt.Println(viper.GetStringMapString("log"))
		// fmt.Println(viper.GetString("log.dir"))
		// fmt.Println(viper.GetString("log.level"))

		// if len(args) < 1 {
		// 	cmd.Help()
		// 	return
		// }
	}
}

// rootCmd.CompletionOptions.DisableDefaultCmd = true

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if strings.Contains(err.Error(), "unknown flag") ||
		    strings.Contains(err.Error(), "unknown shorthand flag") {
			// exit code 126: Command invoked cannot execute
			os.Exit(126)
		}
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configFile := viper.GetString("config-file")

	config.InitializeConfig(configFile)

	logger.InitLogger()
}
