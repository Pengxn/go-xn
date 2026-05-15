package rpc

import (
	"context"

	pb "github.com/Pengxn/go-xn/rpc/proto"
)

type Server struct {
	pb.HealthCheckServer
}

func (s *Server) Ping(_ context.Context, _ *pb.Empty) (*pb.Pong, error) {
	return &pb.Pong{Message: "ok!"}, nil
}
