// +build !windows

package home

import (
	"os"
	"os/user"
)

// key returns the env var name for the user's home dir based on
// the platform being run on
func key() string {
	return "HOME"
}

// home returns the home directory of the current user with the help of
// environment variables depending on the target operating system.
// Returned path should be used with "path/filepath" to form new paths.
//
// If linking statically with cgo enabled against glibc, ensure the
// osusergo build tag is used.
//
// If needing to do nss lookups, do not disable cgo or set osusergo.
func home() string {
	home := os.Getenv(key())
	if home == "" {
		if u, err := user.Current(); err == nil {
			return u.HomeDir
		}
	}
	return home
}
