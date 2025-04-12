package config

import (
	"context"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/Pengxn/go-xn/src/util/file"
)

func OverrideConfigByFlag(ctx context.Context, c *cli.Command) {
	if c.IsSet("config") { // specified config file
		f := c.String("config")
		if !file.IsExist(f) || !file.IsFile(f) {
			slog.Error("specified config file not found", slog.String("filename", f))
		}
		if err := loadConfig(f); err != nil {
			slog.Error("load specified config file failed", slog.Any("error", err))
		}
	}
	// Server config
	if c.IsSet("port") {
		Config.Server.Port = c.String("port")
	}
	if c.IsSet("debug") {
		Config.Server.Debug = c.Bool("debug")
	}
}

func getConfigPathByFlag() string {
	if args := os.Args; len(args) > 0 {
		for k, v := range args {
			if v == "-c" || v == "--config" {
				// Load config file from command line
				if len(args) > k+1 {
					return args[k+1]
				}
			}
		}
	}

	return ""
}
