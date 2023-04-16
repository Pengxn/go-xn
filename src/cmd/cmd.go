package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

// Execute to run cmd
func Execute() error {
	app := &cli.App{
		Name:    "go-xn",
		Usage:   "The platform for publishing and running your blog",
		Version: Version,
		Commands: []*cli.Command{
			Web,
			Update,
			VersionCmd,
		},
	}

	return app.Run(os.Args)
}
