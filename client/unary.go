package main

import (
	"context"
	"log"
	"time"

	pb "github.com/fazlul/grpc-practise/proto"
)

func callSayHello(client pb.GreetServicesClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	//printing the response
	log.Printf("%s", res.Message)
}