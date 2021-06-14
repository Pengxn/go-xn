package home

import (
	"path/filepath"
)

// HomeDir detects and returns the home directory for the executing user.
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func HomeDir() string {
	return home()
}

// ConfigDir retrieves app's config directory for specified OS.
func ConfigDir(app string) string {
	return filepath.Join(HomeDir(), ".config", app)
}

// CacheDir retrieves app's cache directory for specified OS.
func CacheDir(app string) string {
	return filepath.Join(HomeDir(), ".cache", app)
}

// DataDir retrieves app's data directory for specified OS.
func DataDir(app string) string {
	return filepath.Join(HomeDir(), ".local", "share", app)
}

// LogDir retrieves app's log directory for specified OS.
func LogDir(app string) string {
	return filepath.Join(HomeDir(), ".local", "share", app)
}
