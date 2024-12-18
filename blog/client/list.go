package main

import (
	"context"
	"io"
	"log"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("---listbBlogs was invoked---")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling listBlogs: %v\n", err)
	}

	for {
		blog, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading stream during listBlogs: %v\n", err)
		}

		log.Printf("Received blog: %v\n", blog)
	}
}
