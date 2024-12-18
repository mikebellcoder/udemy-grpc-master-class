package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/blog/proto"
)

var svrAddr = "localhost:5051"

func main() {
	conn, err := grpc.NewClient(svrAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connec: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	// // readBlog(c, "non-Id") // invalid id
	updateBlog(c, id)
	listBlogs(c)
	deleteBlog(c, id)
}
