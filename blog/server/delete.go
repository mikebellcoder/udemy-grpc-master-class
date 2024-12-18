package main

import (
	"context"
	"log"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot parse ID")
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot delete object in MongoDB: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Blog was not found")
	}

	log.Println("Blog deleted successfully!")

	return &emptypb.Empty{}, nil
}
