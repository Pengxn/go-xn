package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/Pengxn/go-xn/src/config"
)

// Execute to run cmd
func Execute() {
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
				Value:   config.HomeDir() + string(os.PathSeparator) + "fyj.ini",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln("Fail to Start Command app.", err.Error())
	}
}
