package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

// Writes a fucntion onto a pointer
func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Calculator function was invoked with %v\n", req)
	a := req.GetSum().GetA()
	b := req.GetSum().GetB()
	result := a + b
	res := &calculatorpb.SumResponse{
		Result: result,
	}
	return res, nil
}

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
	calculatorpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
