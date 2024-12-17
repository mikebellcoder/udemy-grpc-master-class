package main

import (
	"io"
	"log"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Max(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {
	log.Println("Max has been invoked")

	maxNum := int64(0)

	for {
		// read from client
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatalf("error reading from client stream: %v\n", err)
		}
		log.Printf("Received: %v\n", req)

		// check if received new max number
		if req.Number > maxNum {
			maxNum = req.Number

			// write to client
			err = stream.Send(&pb.MaxResponse{Result: maxNum})
			if err != nil {
				log.Fatalf("error writing to client stream: %v\n", err)
			}
			log.Printf("Sending: %d\n", maxNum)
		}
	}
}
