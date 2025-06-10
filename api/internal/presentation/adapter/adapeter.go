package adapter

import (
	"github.com/tamaco489/grpc_go_sample/api/internal/presentation/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	healthcheck_pb "github.com/tamaco489/grpc_go_sample/api/internal/gen/proto/healthcheck"
	product_pb "github.com/tamaco489/grpc_go_sample/api/internal/gen/proto/product"

	healthcheck_usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/healthcheck"
	product_usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/product"
)

// SetupControllers resolves dependencies from UseCase creation to Controller creation and returns them
func SetupControllers() *controller.Controllers {
	hcUsecase := healthcheck_usecase.NewUseCase()
	productUsecase := product_usecase.NewUseCase()
	return controller.NewControllers(
		hcUsecase,
		productUsecase,
	)
}

// RegisterGRPCServices registers each service to the grpcServer
func RegisterGRPCServices(grpcServer *grpc.Server, controllers *controller.Controllers) {
	// HealthCheck service
	healthcheck_pb.RegisterHealthCheckServiceServer(grpcServer, controllers.HealthCheck)
	product_pb.RegisterProductServiceServer(grpcServer, controllers.Product)

	// NOTE: Register other services here
	// other.RegisterGRPCServer(grpcServer, controllers.Other)
}

// NewGRPCServer creates a gRPC server with services registered
func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	controllers := SetupControllers()
	RegisterGRPCServices(s, controllers)
	reflection.Register(s)
	return s
}
