package main

import (
	"context"
	"log"
	"math"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt function invoked with: %v\n", in)

	number := in.Number
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Received a negative number: %d", number)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
