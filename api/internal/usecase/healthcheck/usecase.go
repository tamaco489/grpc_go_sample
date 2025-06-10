package healthcheck

import (
	"context"

	"github.com/tamaco489/grpc_go_sample/api/internal/domain/entity"
)

// UseCase defines the use case for health check
type UseCase interface {
	GetHeartbeat(ctx context.Context) (*entity.HealthCheck, error)
}

// useCase provides the implementation of UseCase
type useCase struct{}

// NewUseCase creates a new UseCase instance
func NewUseCase() UseCase {
	return &useCase{}
}

// GetHeartbeat returns the current time of the server
func (u *useCase) GetHeartbeat(ctx context.Context) (*entity.HealthCheck, error) {
	return entity.NewHealthCheck(), nil
}
