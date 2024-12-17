package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax function was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v\n", err)
	}

	// (1,5,3,6,2,20)
	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	waitc := make(chan struct{})

	// send
	go func() {
		for _, req := range reqs {
			log.Printf("Sending: %v\n", req)
			err = stream.Send(req)
			if err != nil {
				log.Printf("Failed to send from client stream: %v\n", err)
				break
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// receive
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Printf("error receiving on client stream: %v\n", err)
				break
			}

			log.Printf("Received: %d\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
