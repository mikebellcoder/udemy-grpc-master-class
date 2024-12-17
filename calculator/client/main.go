package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/calculator/proto"
)

var svrAddr = "localhost:5051"

func main() {
	conn, err := grpc.NewClient(svrAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connec: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doSum(c)
	// doPrimes(c)
	// doAverage(c)
	// doMax(c)
	// doSqrt(c, 10)
	doSqrt(c, -2)
}
