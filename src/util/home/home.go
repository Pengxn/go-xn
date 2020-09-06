package home

import (
	"errors"
	"os"
	"runtime"
)

// HomeDir detects and returns the home directory for the executing user.
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
		// Unix-like system, so just assume Unix.
		// It's mainly linux.
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

// ConfigDir retrieves app's config directory for specified OS.
// darwin:  ~/Library/Preferences/app
// Windows: ~\AppData\Local\app\Config
// Unix:    ~/.config/app
func ConfigDir(app string) string {
	switch runtime.GOOS {
	case "drawin":
		return HomeDir() + "/Library/Preferences/" + app
	case "windows":
		return HomeDir() + `\AppData\Local\` + app + `\Config`
	default:
		return HomeDir() + "/.config/" + app
	}
}

// DataDir retrieves app's data directory for specified OS.
// darwin:  ~/Library/Application Support/app
// Windows: ~\AppData\Local\app
// Unix:    ~/.local/share/app
func DataDir(app string) string {
	switch runtime.GOOS {
	case "drawin":
		return HomeDir() + "/Library/Application Support/" + app
	case "windows":
		return HomeDir() + `\AppData\Local\` + app
	default:
		return HomeDir() + "/.local/share/" + app
	}
}

// CacheDir retrieves app's cache directory for specified OS.
// darwin:  ~/Library/Caches/app
// Windows: ~\AppData\Local\app\Cache
// Unix:    ~/.cache/app
func CacheDir(app string) string {
	switch runtime.GOOS {
	case "drawin":
		return HomeDir() + "/Library/Caches/" + app
	case "windows":
		return HomeDir() + `\AppData\Local\` + app + `Cache`
	default:
		return HomeDir() + "/.cache/" + app
	}
}

// LogDir retrieves app's log directory for specified OS.
// darwin:  ~/Library/Logs/app
// Windows: ~\AppData\Local\app\Logs
// Unix:    ~/.local/share/app
func LogDir(app string) string {
	switch runtime.GOOS {
	case "drawin":
		return HomeDir() + "/Library/Logs/" + app
	case "windows":
		return HomeDir() + `\AppData\Local\` + app + `Log`
	default:
		return HomeDir() + "/.local/share/" + app
	}
}
