package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

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

// Stream Server
func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition function was invoked with %v\n", req)
	N := int(req.GetPrimeNumberDecomposition().GetA())
	k := 2

	for N > 1 {
		if N%k == 0 { // if k evenly divides into N
			
			// strconv.Itoa converst an integer to a string
			result := "N = " + strconv.Itoa(N) + " Value = " + strconv.Itoa(k)
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				Result: result,
			}
	
			stream.Send(res)
			time.Sleep(1000 + time.Second)

			N = N / k // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}


	}
	return nil
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
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
