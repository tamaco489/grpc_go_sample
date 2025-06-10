package main

import (
	"log"
	"net"

	"github.com/tamaco489/grpc_go_sample/api/internal/presentation/adapter"
)

const port = ":50051"

func main() {
	// TCPリスナーを作成
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPCサーバーを起動
	s := adapter.NewGRPCServer()

	// サーバーを起動
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
