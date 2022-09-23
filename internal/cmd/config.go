package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


func NewConfigCmd(root *cobra.Command) *cobra.Command {
	c := &cobra.Command{
		Use:   "config",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("config called")

			configFile, _ := cmd.Flags().GetString("config-file")
			fmt.Println(configFile)
		},
	}
	root.AddCommand(c)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return c
}

// register the subcommand into rootCmd
// configCmd represents the config command
var configCmd = NewConfigCmd(rootCmd)

func configure(v *viper.Viper, f *cobra.Command) {
	// Todo
}
