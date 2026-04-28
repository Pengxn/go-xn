package cmd

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

// Execute to run cmd
func Execute(ctx context.Context) error {
	app := &cli.Command{
		Name:    "go-xn",
		Usage:   "The platform for publishing and running your blog",
		Version: version,
		Commands: []*cli.Command{
			webCmd,
			updateCmd,
			versionCmd,
			agentCmd,
		},
	}

	return app.Run(ctx, os.Args)
}
