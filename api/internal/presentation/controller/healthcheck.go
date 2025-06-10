package controller

import (
	"context"

	healthcheck "github.com/tamaco489/grpc_go_sample/api/internal/gen/proto/healthcheck"
	usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/healthcheck"
)

// HealthCheckController is a struct that implements the gRPC HealthCheck service
type HealthCheckController struct {
	usecase usecase.UseCase

	healthcheck.UnimplementedHealthCheckServiceServer
}

// NewHealthCheckController creates a new HealthCheckController instance
func NewHealthCheckController(u usecase.UseCase) *HealthCheckController {
	return &HealthCheckController{usecase: u}
}

// GetHeartbeat implements the gRPC HealthCheck service method
func (c *HealthCheckController) GetHeartbeat(ctx context.Context, req *healthcheck.GetHeartbeatRequest) (*healthcheck.GetHeartbeatResponse, error) {
	// Delegate the processing to the usecase layer
	health, err := c.usecase.GetHeartbeat(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the entity to the gRPC response
	return &healthcheck.GetHeartbeatResponse{
		ServerTime: health.ServerTime.Unix(), // Unix seconds are expected
	}, nil
}
