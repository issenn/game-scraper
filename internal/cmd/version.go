package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/issenn/game-scraper/internal/meta"
)


var (
	VERSION    string
	GO_VERSION string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version number.",
	Long: `
Show the necro version number, the go version, the build target
OS and architecture, the runtime OS and kernel version and bitness,
build tags and the type of executable (static or dynamic).

For example:
    $ necro version
    necro v1.55.0
    - os/version: ubuntu 18.04 (64 bit)
    - os/kernel: 4.15.0-136-generic (x86_64)
    - os/type: linux
    - os/arch: amd64
    - go/version: go1.16
    - go/linking: static
    - go/tags: none

Note: ...

If you supply the --check flag, then it will do an online check to
compare your version with the latest release and the latest beta.
  $ necro version --check
  yours:  1.42.0.6
  latest: 1.42          (released 2022-07-16)
  beta:   1.42.0.5      (released 2022-07-17)
Or
  $ necro version --check
  yours:  1.41
  latest: 1.42          (released 2022-07-16)
    upgrade: https://downloads.issenn.ml/necro/v1.42
  beta:   1.42.0.5      (released 2022-07-17)
    upgrade: https://beta.issenn.ml/necro/v1.42-005-g56e1e820

`,
	Run: runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {
	checkFlag, _ := cmd.Flags().GetBool("check")
	if !checkFlag {
		if VERSION == "" {
			fmt.Println("could not determine build information")
		} else {
			// fmt.Printf("version called: Necro Static Site Generator v%v -- HEAD\n", VERSION)
			ShowVersion()
		}
	}
}

// strip a leading v off the string
func stripV(s string) string {
	if len(s) > 0 && s[0] == 'v' {
		return s[1:]
	}
	return s
}

// ShowVersion prints the version to stdout
func ShowVersion() {
	osVersion, osKernel := "", ""
	if osVersion == "" {
		osVersion = "unknown"
	}
	if osKernel == "" {
		osKernel = "unknown"
	}

	linking, tagString := "", ""

	fmt.Printf("necro %s\n", VERSION)
	fmt.Printf("- os/version: %s\n", osVersion)
	fmt.Printf("- os/kernel: %s\n", osKernel)
	fmt.Printf("- os/type: %s\n", meta.VersionMeta.GOOS)
	fmt.Printf("- os/arch: %s\n", meta.VersionMeta.GOARCH)
	fmt.Printf("- go/version: %s\n", meta.VersionMeta.GoVersion)
	fmt.Printf("- go/linking: %s\n", linking)
	fmt.Printf("- go/tags: %s\n", tagString)
}

// GetVersion gets the version available for download
func GetVersion() {
	// Todo
}

// CheckVersion checks the installed version against available downloads
func CheckVersion() {
	// Todo
}

func init() {
	initByLDFlags()

	rootCmd.AddCommand(versionCmd)

	var (
		shortened = false
		output    = "string"
		check     = false
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	versionCmd.Flags().BoolVarP(&shortened, "short", "s", false, "Use shortened output for version information.")
	versionCmd.Flags().StringVarP(&output, "output", "o", "json", "Output format. One of 'string', 'yaml' or 'json'.")
	versionCmd.Flags().BoolVarP(&check, "check", "", false, "Check for new version")
}

func initByLDFlags() {
	VERSION    = meta.VersionMeta.Version
	GO_VERSION = meta.VersionMeta.GoVersion
}
