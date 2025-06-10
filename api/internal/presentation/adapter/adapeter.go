package adapter

import (
	"github.com/tamaco489/grpc_go_sample/api/internal/presentation/controller"
	"google.golang.org/grpc"

	healthcheck_pb "github.com/tamaco489/grpc_go_sample/api/internal/gen/proto/healthcheck"
	healthcheck_usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/healthcheck"
)

// SetupControllers resolves dependencies from UseCase creation to Controller creation and returns them
func SetupControllers() *controller.Controllers {
	hcUsecase := healthcheck_usecase.NewUseCase()
	return controller.NewControllers(hcUsecase)
}

// RegisterGRPCServices registers each service to the grpcServer
func RegisterGRPCServices(grpcServer *grpc.Server, controllers *controller.Controllers) {
	// HealthCheck service
	healthcheck_pb.RegisterHealthCheckServiceServer(grpcServer, controllers.HealthCheck)

	// NOTE: Register other services here
	// other.RegisterGRPCServer(grpcServer, controllers.Other)
}

// NewGRPCServer creates a gRPC server with services registered
func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	controllers := SetupControllers()
	RegisterGRPCServices(s, controllers)
	return s
}
