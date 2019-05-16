package config

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-xn/src/config"
)

func TestHomeDir(t *testing.T) {
	assert := assert.New(t)

	home, _ := user.Current()
	dir, _ := config.HomeDir()

	assert.Equal(home.HomeDir, dir)
}
