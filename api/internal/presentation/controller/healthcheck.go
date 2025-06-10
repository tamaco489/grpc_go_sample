package controller

import (
	"context"

	healthcheck "github.com/tamaco489/grpc_go_sample/api/internal/gen/proto/healthcheck"
	usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/healthcheck"
)

// HealthCheckController はgRPCのHealthCheckサービスの実装を担う構造体
type HealthCheckController struct {
	usecase usecase.UseCase

	healthcheck.UnimplementedHealthCheckServiceServer
}

// NewHealthCheckController はコントローラのコンストラクタ
func NewHealthCheckController(u usecase.UseCase) *HealthCheckController {
	return &HealthCheckController{usecase: u}
}

// GetHeartbeat はgRPCのHealthCheckサービスのメソッドを実装
func (c *HealthCheckController) GetHeartbeat(ctx context.Context, req *healthcheck.GetHeartbeatRequest) (*healthcheck.GetHeartbeatResponse, error) {
	// usecase層に処理を委譲
	health, err := c.usecase.GetHeartbeat(ctx)
	if err != nil {
		return nil, err
	}

	// entityからgRPCのレスポンスへ変換
	return &healthcheck.GetHeartbeatResponse{
		ServerTime: health.ServerTime.Unix(), // Unix秒で渡す想定
	}, nil
}
