package utils


import (
	"os"
	// "io/fs"
	"errors"
	"path/filepath"
)


func PathExists(path string) (bool, error) {
	if path == "" {
		return false, nil
	}
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return err == nil, err
}

func LookupConfigPath(name string) []string {
	var paths []string

	paths = append(paths, ".")
	paths = updateConfigPath(paths, "config")
	paths = updateConfigPath(paths, filepath.Join(XDGBaseDirectory.XDGConfigHome, name))
	paths = updateConfigPath(paths, filepath.Join(XDGBaseDirectory.Home, "." + name))
	paths = updateConfigPath(paths, filepath.Join("/etc", name))

	return paths
}

func updateConfigPath(paths []string, path string) []string {
	if fs, _ := PathExists(path); fs {
		return append(paths, path)
	}
	return paths
}
