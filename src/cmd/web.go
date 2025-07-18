package cmd

import (
	"context"
	"log"

	"github.com/urfave/cli/v3"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/lib/webauthn"
	"github.com/Pengxn/go-xn/src/model"
	"github.com/Pengxn/go-xn/src/route"
	"github.com/Pengxn/go-xn/src/util/otel"
	slogger "github.com/Pengxn/go-xn/src/util/slog"
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

	configFile = &cli.StringFlag{
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
		Name:    "debug",
		Aliases: []string{"d"},
		Usage:   "Enable debug mode, print extra debugging information",
	}
)

// runWeb is the entry point to the web server.
// Parses the arguments, routes and others.
func runWeb(ctx context.Context, c *cli.Command) error {
	// Override config by cli flag
	config.OverrideConfigByFlag(ctx, c)

	// Initialize database and tables
	model.InitTables()

	// Initialize webauthn
	webauthn.InitWebAuthn(ctx, config.Config.WebAuthn)

	// Initialize the logger
	ctx = context.WithValue(ctx, slogger.CtxVersionKey, version)
	slogger.SetLogger(ctx, config.Config.Logger)

	// Initialize OpenTelemetry
	shutdown := otel.SetOtel(ctx, config.Config.Otel)
	defer shutdown(ctx)

	err := route.InitRoutes(ctx, config.Config.Server)
	if err != nil {
		log.Fatalln("fail to init routes", err)
	}

	return err
}
