package main

import (
	"log"
	"net"
	pb "github.com/fazlul/grpc-practise/proto"
	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

type helloServer struct {
	pb.GreetServicesServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServicesServer(grpcServer, &helloServer{})

	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
