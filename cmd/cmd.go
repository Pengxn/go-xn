package cmd

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

// Execute to run cmd
func Execute() error {
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

	return app.Run(context.Background(), os.Args)
}
