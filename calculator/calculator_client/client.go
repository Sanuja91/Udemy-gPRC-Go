package main

import (
	"context"
	"fmt"
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
	c := calculatorpb.NewSumServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)

}

func doUnary(c calculatorpb.SumServiceClient){
	fmt.Println("Starting to do a Unary RPC ...")
	req := &calculatorpb.SumRequest{
		Sum:&calculatorpb.Sum{
			A:5,
			B:4,
		},
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil{
		log.Fatalf("Error while callign Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)

}
