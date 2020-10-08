package main

import (
	"fmt"
	"log"
	"net"

	"github.com/meghashyamc/letschat-server/chat"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":"+chat.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	fmt.Println("server running on port", chat.Port)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
