package config

import (
	"os"
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-xn/src/config"
)

func patchEnv(key, value string) func() {
	bck := os.Getenv(key)
	deferFunc := func() {
		os.Setenv(key, bck)
	}

	if value != "" {
		os.Setenv(key, value)
	} else {
		os.Unsetenv(key)
	}

	return deferFunc
}

func TestHomeDir(t *testing.T) {
	assert := assert.New(t)

	home, _ := user.Current()
	dir, _ := config.HomeDir()

	assert.Equal(home.HomeDir, dir)
}
