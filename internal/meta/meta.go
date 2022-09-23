package meta

import (
	"runtime"
	"runtime/debug"
)


// Version information.
var (
	AppName   string

	// Version is the compile-time set version
	Version   string

	Revision  = "unknown"
	Commit    = Revision

	Date      = "unknown"
	BuiltDate = Date

	BuiltBy   = "unknown"

	PipeID    = "unknown"

	GoVersion string

	// UserAgent is the http client identifier derived from Version
	// UserAgent string
)

type meta struct {
	AppName   string
}

//Details contains the state of the executable at build time
type versionMeta struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuiltDate string `json:"buildDate"`
	BuiltBy   string `json:"builtBy"`
	PipeID    string `json:"pipeID"`
	GoVersion string `json:"goVersion"`
	GOOS      string
	GOARCH    string
}

var Meta meta
var VersionMeta versionMeta

func getMeta() meta {
	return meta{
		AppName:   getAppName(),
	}
}

func SetDefaultAppName(name string) {
	if AppName == "" || AppName == "unknown" {
		AppName = name
		Meta.AppName = name
	}
}

func SetAppName(name string) {
	AppName = name
	Meta.AppName = name
}

func getAppName() string {
	if AppName == "" {
		AppName = "unknown"
	}
	return AppName
}

// Get returns a prepopulated Details struct
func getVersionMeta() versionMeta {
	return versionMeta{
		Version:   getVersion(),
		Commit:    Commit,
		BuiltDate: BuiltDate,
		BuiltBy:   BuiltBy,
		PipeID:    PipeID,
		GoVersion: getGoVersion(),
		GOOS:      runtime.GOOS,
		GOARCH:    runtime.GOARCH,
	}
}

func getVersion() string {
	if Version == "" {
		i, ok := debug.ReadBuildInfo()
		if !ok {
			Version = "0.0.0-unknown"
		}
		Version = "0.0.0-" + i.Main.Version
	}
	return Version
}

func getGoVersion() string {
	GoVersion = runtime.Version()
	if GoVersion == "" {
		i, ok := debug.ReadBuildInfo()
		if !ok {
			GoVersion = "unknown"
		}
		GoVersion = i.GoVersion
	}
	return GoVersion
}

func init() {
	Meta = getMeta()
	VersionMeta = getVersionMeta()
	// UserAgent = "Watchtower/" + Version
}
