package main

import (
	"fmt"
	"log"
	"net"

	"../greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello World")

	// Port Binding
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Bind the port to the gPRC server
	s := grpc.NewServer()

	// Add the services here
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
