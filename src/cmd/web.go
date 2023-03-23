package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/Pengxn/go-xn/src/route"
	"github.com/Pengxn/go-xn/src/util/home"
	"github.com/Pengxn/go-xn/src/util/log"
)

var (
	// Web is "web" subcommand. It's used to run web server.
	Web = &cli.Command{
		Name:  "web",
		Usage: "Start web server interface for blog",
		Description: `Run a performant web server which serves the site
for blog. If '--port' flag is not used, it will use
port 8080 by default.`,
		Action: runWeb,
		Flags: []cli.Flag{
			port,
			webroot,
		},
	}

	port = &cli.IntFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Usage:   "Temporary port number to prevent conflict",
		Value:   3000,
	}
	webroot = &cli.PathFlag{
		Name:    "webroot",
		Aliases: []string{"w"},
		Usage:   "Web root path that used by server",
		Value:   home.HomeDir(),
	}
)

// runWeb is the entry point to the web server.
// Parses the arguments, routes and others.
func runWeb(c *cli.Context) error {
	err := route.InitRoutes()
	if err != nil {
		log.Fatalln("Fail to Start Web Server.", err)
	}

	return err
}
