package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetEveryone(stream grpc.BidiStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Println("GreetEveryone was invoked")

	for {
		// read from client
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatalf("error reading from client stream: %v\n", err)
		}
		log.Printf("Received: %v\n", req)

		// write to client
		res := fmt.Sprintf("Hello %s!", req.FirstName)
		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalf("error writing back to client stream: %v\n", err)
		}
		log.Printf("Sending: %s\n", res)
	}
}
