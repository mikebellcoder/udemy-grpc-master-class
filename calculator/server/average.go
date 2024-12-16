package main

import (
	"io"
	"log"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Average(stream grpc.ClientStreamingServer[pb.AverageRequest, pb.AverageResponse]) error {
	log.Println("Average function was invoked")

	count := float64(0)
	total := float64(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := total / count
			log.Printf("Count: %v Total: %v \n", count, total)
			log.Printf("Average: %.2f\n", avg)
			return stream.SendAndClose(&pb.AverageResponse{Result: avg})
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		log.Printf("Recieved: %v\n", req)
		count++
		total += float64(req.Number)
	}
}
