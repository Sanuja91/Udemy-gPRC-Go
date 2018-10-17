package main

import (
	"context"
	"fmt"
	"log"

	"../greetpb"
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

	// Greet Service Client
	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do a Unary RPC ...")
	req := &greetpb.GreetRequest{
		Greeting:&greetpb.Greeting{
			FirstName:"Sanuja",
			LastName:"Cooray",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil{
		log.Fatalf("Error while callign Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)

}
