package meta

import (
	"runtime"
	"runtime/debug"
)


var (
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

var VersionMeta versionMeta

// Get returns a prepopulated Details struct
func Get() versionMeta {
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
	VersionMeta = Get()
	// UserAgent = "Watchtower/" + Version
}
