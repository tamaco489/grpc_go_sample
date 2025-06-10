package healthcheck

import (
	"context"

	"github.com/tamaco489/grpc_go_sample/api/internal/domain/entity"
)

// UseCase はヘルスチェックのユースケースを定義します
type UseCase interface {
	GetHeartbeat(ctx context.Context) (*entity.HealthCheck, error)
}

// useCase はUseCaseの実装を提供します
type useCase struct{}

// NewUseCase は新しいUseCaseのインスタンスを作成します
func NewUseCase() UseCase {
	return &useCase{}
}

// GetHeartbeat はサーバーの現在時刻を返します
func (u *useCase) GetHeartbeat(ctx context.Context) (*entity.HealthCheck, error) {
	return entity.NewHealthCheck(), nil
}
