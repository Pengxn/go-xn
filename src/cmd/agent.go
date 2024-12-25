package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	agent = &cli.Command{
		Name:  "agent",
		Usage: "Run as an agent",
		Description: `Run as an agent, which can be used to manage the server.
It run as a gRPC server, it only support gRPC protocol.`,
		Flags:  []cli.Flag{agentPort},
		Action: runGRPCServer,
	}

	agentPort = &cli.IntFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Usage:   "Port to listen",
		Value:   7992,
	}
)

func runGRPCServer(c *cli.Context) error {
	// TODO: run gRPC server
	fmt.Println(c.Int("port"))
	return nil
}
