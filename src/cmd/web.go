package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"

	"github.com/Pengxn/go-xn/src/route"
	"github.com/Pengxn/go-xn/src/util/home"
	"github.com/Pengxn/go-xn/src/util/log"
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
		Value:   home.HomeDir(),
	}
)

// runWeb is the entry point to the web server.
// Parses the arguments, routes and others.
func runWeb(c *cli.Context) error {
	gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	route.InitRoutes(g)

	if err := g.Run(":" + c.String("port")); err != nil {
		log.Fatalln("Fail to Start Web Server.", err.Error())
	}

	return nil
}
