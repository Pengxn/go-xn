package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"

	"github.com/Pengxn/go-xn/src/rpc"
	"github.com/Pengxn/go-xn/src/rpc/proto"
)

var (
	// agentCmd is "agent" subcommand.
	// It's used to run a gRPC server as an agent side.
	agentCmd = &cli.Command{
		Name:        "agent",
		Usage:       "Run as an agent",
		Description: `Run a gRPC serveras an agent side, it only support gRPC protocol.`,
		Flags:       []cli.Flag{agentPort},
		Action:      runGRPCServer,
		Subcommands: []*cli.Command{agentPingCmd},
	}
	agentPingCmd = &cli.Command{
		Name:        "ping",
		Usage:       "Ping the agent",
		Description: `Ping the agent, which can be used to check the agent status.`,
		Flags:       []cli.Flag{pingUrl},
		Action:      pingAgent,
	}

	pingUrl = &cli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Usage:   "URL to ping",
		Value:   "127.0.0.1:7992",
	}

	agentPort = &cli.IntFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Usage:   "Port to listen",
		Value:   7992,
	}
)

func runGRPCServer(c *cli.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Int("port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gs := grpc.NewServer()
	defer gs.Stop()
	proto.RegisterHeathCheckServer(gs, &rpc.Server{})

	log.Printf("server listening at %v", lis.Addr())
	err = gs.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func pingAgent(c *cli.Context) error {
	pong, err := rpc.Client(c.String("url"))
	if err != nil {
		log.Fatalln("failed to ping agent:", err)
	}

	fmt.Println(pong.Message)

	return nil
}
