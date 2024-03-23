package main

import (
	"io"
	"log"

	 pb "github.com/fazlul/grpc-practise/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetServices_SayHelloBidirectionalStreamingServer) error {
	for {

		//receiving the request
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)

		//make a response
		res := &pb.HelloResponse {
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}