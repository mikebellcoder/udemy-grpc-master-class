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
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Printf("Error happend while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
