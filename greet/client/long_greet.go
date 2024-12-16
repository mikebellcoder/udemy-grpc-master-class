package main

import (
	"context"
	"log"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Mike"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	for _, r := range reqs {
		log.Printf("Sending req: %v", r)
		err := stream.Send(r)
		if err != nil {
			log.Fatalf("Error while streaming from LongGreet client: %v\n", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error closing stream: %v\n", err)
	}

	log.Printf("Received %s\n", res.Result)
}
