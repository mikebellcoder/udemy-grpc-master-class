package main

import (
	"log"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials())) // nb: using NewClient instead of Dial
	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	readBlog(c, "aNonExistingID")
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}
