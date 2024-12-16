package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average from client: %v\n", err)
	}

	for n := int64(1); n < 5; n++ {
		req := &pb.AverageRequest{Number: n}
		log.Printf("Sending req: %v \n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error closing stream: %v\n", err)
	}

	log.Printf("Average: %.2f\n", res.Result)
}
