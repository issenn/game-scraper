package meta

var (
	// Version is the compile-time set version
	Version = "v0.0.0-unknown"

	Commit = ""
	Revision = Commit

	Date = ""
	buildDate = Date

	BuiltBy = ""

	// UserAgent is the http client identifier derived from Version
	// UserAgent string
)

// func init() {
// 	UserAgent = "Watchtower/" + Version
// }
