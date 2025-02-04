package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/lib/webauthn"
	"github.com/Pengxn/go-xn/src/model"
	"github.com/Pengxn/go-xn/src/route"
	"github.com/Pengxn/go-xn/src/util/log"
)

var (
	// Web is "web" subcommand. It's used to run web server.
	webCmd = &cli.Command{
		Name:  "web",
		Usage: "Start web server interface for blog",
		Description: `Run a performant web server which serves the site for blog.
If '--port' flag is not used, it will use port 7991 by default.`,
		Action: runWeb,
		Flags: []cli.Flag{
			configFile,
			port,
			debug,
		},
	}

	configFile = &cli.PathFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Usage:       "Custom configuration file path",
		Value:       "fyj.ini",
		DefaultText: "fyj.ini",
	}
	port = &cli.IntFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Usage:   "Temporary port number to prevent conflict",
		Value:   7991,
	}
	debug = &cli.BoolFlag{
		Name:               "debug",
		Aliases:            []string{"d"},
		Usage:              "Enable debug mode, print extra debugging information",
		DisableDefaultText: true,
	}
)

// runWeb is the entry point to the web server.
// Parses the arguments, routes and others.
func runWeb(c *cli.Context) error {
	// Override config by cli flag
	config.OverrideConfigByFlag(c)
	model.InitTables()
	webauthn.InitWebAuthn()

	err := route.InitRoutes()
	if err != nil {
		log.Fatalln("fail to start web server:", err)
	}

	return err
}
