package main

import (
	"context"

	pb "github.com/rakshitgondwal/grpc-go/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResp, error) {
	return &pb.HelloResp{
		Message: "Hello",
	}, nil
}
