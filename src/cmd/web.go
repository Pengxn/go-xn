package cmd

import (
	"github.com/urfave/cli/v2"

	app "github.com/Pengxn/go-xn/src"
	"github.com/Pengxn/go-xn/src/config"
)

var (
	// Web is "web" subcommand. It's used to run web server
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
		Aliases: []string{"r"},
		Usage:   "Web root path that used by server",
		Value:   config.HomeDir(),
	}
)

// runWeb run web server
func runWeb(c *cli.Context) error {
	app.Run()

	return nil
}
