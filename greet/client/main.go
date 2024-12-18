package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/mikebellcoder/udemy-grpc-master-class/greet/proto"
)

var svrAddr string = "localhost:5051"

func main() {
	opts := []grpc.DialOption{}
	tls := true
	if tls {
		creds, err := credentials.NewClientTLSFromFile("ssl/ca.crt", "")
		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.NewClient(svrAddr, opts...) // nb: using NewClient instead of Dial
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 3*time.Second)
}
