package rpc

import (
	"context"

	pb "github.com/Pengxn/go-xn/src/rpc/proto"
)

type Server struct {
	pb.HeathCheckServer
}

func (s *Server) HeartBeat(_ context.Context, _ *pb.Empty) (*pb.Pong, error) {
	return &pb.Pong{Message: "ok!"}, nil
}
