package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/greet/proto"
)

var svrAddr string = "localhost:5051"

func main() {
	conn, err := grpc.NewClient(svrAddr, grpc.WithTransportCredentials(insecure.NewCredentials())) // nb: using NewClient instead of Dial
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	// doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	doGreetWithDeadline(c, 3*time.Second)
}
