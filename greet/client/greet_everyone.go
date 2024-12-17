package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone function was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("failed to create client stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{{FirstName: "Mike"}, {FirstName: "Clement"}, {FirstName: "Maria"}, {FirstName: "Test"}}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending: %v\n", req)
			err = stream.Send(req)
			if err != nil {
				log.Fatalf("failed to send from client stream: %v\n", err)
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Printf("error receiving on client stream: %v\n", err)
				break
			}

			log.Printf("Received: %s\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
