package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"../calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// Defer is executed once everything else is done
	defer cc.Close()

	// Calculator Service Client
	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	// doUnary(c)
	// doServerStreaming(c)

	doClientStreaming(c)
}

// Streams data from the client
func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a ComputeAverage Client Streaming RPC...")
	stream,err :=c.ComputeAverage(context.Background())

	if err != nil{
		log.Fatalf("Error while opening stream: %v", err)
	}

	numbers:= []int32{3,5,9,54,23}

	for _, number := range numbers{
		fmt.Printf("Sending number: %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
		Number:number,
		})
	}

	res,err := stream.CloseAndRecv()
	
	if err != nil{
		log.Fatalf("Error while receiving response: %v", err)
	}

	fmt.Printf("The Average is: %v", res.GetAverage())
}
// Streams data from the Server
func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		PrimeNumberDecomposition: &calculatorpb.PrimeNumberDecomposition{
			A: 500,
		}}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDecomposition RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		// If the error is because the file has reached the End of the File
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream:%v", err)
		}

		log.Printf("Response from PrimeNumberDecomposition: %v", msg.GetResult())

	}
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC ...")
	req := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Sum{
			A: 5,
			B: 4,
		},
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while callign Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)

}
