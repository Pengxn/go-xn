package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-xn/src/config"
)

func TestDBUrl(t *testing.T) {
	assert := assert.New(t)

	result := "root:password@tcp(127.0.0.1:3306)/fyj?charset=utf8"
	url := config.DBUrl()

	assert.Equal(url, result)
}
