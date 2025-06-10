package adapter

import (
	"github.com/tamaco489/grpc_go_sample/api/internal/presentation/controller"
	"google.golang.org/grpc"

	healthcheck_pb "github.com/tamaco489/grpc_go_sample/api/internal/gen/proto/healthcheck"
	healthcheck_usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/healthcheck"
)

// SetupControllers は UseCase生成～Controller生成までの依存関係を解決して返す
func SetupControllers() *controller.Controllers {
	hcUsecase := healthcheck_usecase.NewUseCase()
	return controller.NewControllers(hcUsecase)
}

// RegisterGRPCServices grpcServerへ各サービスを登録する
func RegisterGRPCServices(grpcServer *grpc.Server, controllers *controller.Controllers) {
	// HealthCheckサービス
	healthcheck_pb.RegisterHealthCheckServiceServer(grpcServer, controllers.HealthCheck)

	// NOTE: Register other services here
	// other.RegisterGRPCServer(grpcServer, controllers.Other)
}

// NewGRPCServer は gRPC サーバーを作成し、サービス登録済みの状態で返します。
func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	controllers := SetupControllers()
	RegisterGRPCServices(s, controllers)
	return s
}
