package adapter

import (
	"github.com/tamaco489/grpc_go_sample/api/internal/presentation/controller"
	"google.golang.org/grpc"

	healthcheck "github.com/tamaco489/grpc_go_sample/api/internal/gen/proto/healthcheck"
)

// RegisterGRPCServices grpcServerへ各サービスを登録する
func RegisterGRPCServices(grpcServer *grpc.Server, controllers *controller.Controllers) {
	// HealthCheckサービス
	healthcheck.RegisterHealthCheckServiceServer(grpcServer, controllers.HealthCheck)

	// 他サービスの登録もここでまとめる
	// other.RegisterGRPCServer(grpcServer)
}
