package main

import (
	"context"
	"log"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--createBlog was invoked--")
	blog := &pb.Blog{
		AuthorId: "Mike",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Undexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}
