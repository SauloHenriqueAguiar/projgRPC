package main

import (
	"fmt"
	"log"
	"net"

	"projgRPC/proto"

	"projgRPC/server/database"
	"projgRPC/server/handlers"

	"google.golang.org/grpc"
)

func main() {
	database.Connect()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterProductServiceServer(s, &handlers.ProductHandler{})

	fmt.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
