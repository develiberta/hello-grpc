package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")

	// TCP connection listen port open
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := chat.Server{}

	// connection port with gRPC Server
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
