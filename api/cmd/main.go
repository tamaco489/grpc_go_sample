package main

import (
	"log"
	"net"

	"github.com/tamaco489/grpc_go_sample/api/internal/presentation/adapter"
)

const port = ":50051"

func main() {
	// Create a TCP listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the gRPC server
	s := adapter.NewGRPCServer()

	// Start the server
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
