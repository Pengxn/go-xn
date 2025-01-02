package rpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Pengxn/go-xn/src/rpc/proto"
)

func Client(address string) (*pb.Pong, error) {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c := pb.NewHeathCheckClient(conn)
	r, err := c.Ping(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	return r, nil
}
