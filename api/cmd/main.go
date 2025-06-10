package main

import (
	"log"
	"net"

	"github.com/tamaco489/grpc_go_sample/api/internal/presentation/adapter"
	"github.com/tamaco489/grpc_go_sample/api/internal/presentation/controller"
	"github.com/tamaco489/grpc_go_sample/api/internal/usecase/healthcheck"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	// TCPリスナーを作成
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPCサーバーを作成
	s := grpc.NewServer()

	controllers := controller.NewControllers(healthcheck.NewUseCase())
	adapter.RegisterGRPCServices(s, controllers)

	// サーバーを起動
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
