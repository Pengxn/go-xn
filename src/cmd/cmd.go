package cmd

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/Pengxn/go-xn/src/util/home"
)

// Execute to run cmd
func Execute() error {
	app := &cli.App{
		Name:    "go-xn",
		Usage:   "The platform for publishing and running your blog",
		Version: Version,
		Commands: []*cli.Command{
			Web,
			VersionCmd,
		},
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Custom configuration `file` path",
				Value:   home.HomeDir() + string(os.PathSeparator) + "fyj.ini",
			},
		},
	}

	return app.Run(os.Args)
}
