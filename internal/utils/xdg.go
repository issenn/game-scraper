package utils


import (
	"os"
	"fmt"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)


var (
	HOME            string
	XDG_CONFIG_HOME string
	XDG_CACHE_HOME  string
	XDG_DATA_HOME   string
)

type _XDGBaseDirectory struct {
	Home          string
	XDGConfigHome string
	XDGCacheHome  string
	XDGDataHome   string
}

var XDGBaseDirectory _XDGBaseDirectory

func GetXDG() _XDGBaseDirectory {
	return _XDGBaseDirectory{
		Home: GetHome(),
		XDGConfigHome: GetXDGConfigHome(),
		XDGCacheHome: "",
		XDGDataHome: "",
	}
}

func GetHome() string {
	// Find home directory.
	HOME, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return HOME
}

func GetXDGConfigHome() string {
	// Find XDG Base Directory.
	// fmt.Println("XDG_CONFIG_HOME:", viper.GetString("XDG_CONFIG_HOME"))
	if viper.GetString("XDG_CONFIG_HOME") != "" {
		XDG_CONFIG_HOME = viper.GetString("XDG_CONFIG_HOME")
	} else if xdg.ConfigHome != "" {
		XDG_CONFIG_HOME = xdg.ConfigHome
	} else {
		// Search config in home directory with name ".config" (without extension).
		home := GetHome()
		XDG_CONFIG_HOME = filepath.Join(home, ".config")
	}
	return XDG_CONFIG_HOME
}

func init() {
	XDGBaseDirectory = GetXDG()
}
