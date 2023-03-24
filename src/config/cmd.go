package config

import (
	"github.com/urfave/cli/v2"

	"github.com/Pengxn/go-xn/src/util/file"
	"github.com/Pengxn/go-xn/src/util/log"
)

func OverrideConfigByFlag(c *cli.Context) {
	if c.IsSet("config") { // specified config file
		f := c.Path("config")
		if !file.IsExist(f) || !file.IsFile(f) {
			log.Error("Specified config file not found: " + f)
		}
		// TODO: load specified config file
	}
	// Server config
	if c.IsSet("port") {
		Config.Server.Port = c.String("port")
	}
	if c.IsSet("debug") {
		Config.Server.Debug = c.Bool("debug")
	}
}
