//go:build windows

package home

import (
	"os"
)

// key returns the env var name for the user's home dir based on
// the platform being run on
func key() string {
	return "USERPROFILE"
}

// home returns the home directory of the current user with the help of
// environment variables depending on the target operating system.
// Returned path should be used with "path/filepath" to form new paths.
func home() string {
	return os.Getenv(key())
}
