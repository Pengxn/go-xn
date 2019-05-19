package config

import (
	"errors"
	"os"
	"runtime"
)

// HomeDir detectes and returns the home directory for the executing user.
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func HomeDir() string {
	var (
		home string
		err  error
	)

	if runtime.GOOS == "windows" {
		home, err = dirWindows()
	} else {
		// Unix-like system, so just assume Unix
		home, err = dirUnix()
	}

	if err != nil {
		panic(err)
	}

	return home
}

func dirWindows() (string, error) {
	home := os.Getenv("USERPROFILE")

	if home == "" {
		return "", errors.New("Can't find 'USERPROFILE' environment variable")
	}

	return home, nil
}

func dirUnix() (string, error) {
	home := os.Getenv("HOME")

	if home == "" {
		return "", errors.New("Can't find 'HOME' environment variable")
	}

	return home, nil
}
