package main

import (
	"context"
	"io"
	"log"

	pb "github.com/fazlul/grpc-practise/proto"
)

func callSayHelloServerStream(client pb.GreetServicesClient, names *pb.NameList) {
	log.Printf("Streaming from server started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}
		log.Println(message)
	}

	log.Printf("Streaming Finished")
}