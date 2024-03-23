package main

import (
	"log"
	pb "github.com/fazlul/grpc-practise/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":5000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect :%v",err)
	}
	defer conn.Close()
	client :=pb.NewGreetServicesClient(conn)
    names := &pb.NameList{
		Names: []string{"bodana", "lungi", "duti"},
	}

	callSayHello(client) //unart api

	callSayHelloServerStream(client, names) //serverstreaming api

	callSayHelloClientStream(client, names) //clientStreaming api

	callSayHelloBidirectionalStream(client, names)

}