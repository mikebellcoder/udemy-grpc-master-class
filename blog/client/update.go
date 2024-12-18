package main

import (
	"context"
	"log"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Mike",
		Title:    "A new title",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Printf("Blog was updated: %v\n", newBlog)
}
