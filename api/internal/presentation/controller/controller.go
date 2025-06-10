package controller

import usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/healthcheck"

// 複数コントローラをまとめる構造体
type Controllers struct {
	HealthCheck *HealthCheckController
	// 他コントローラもあればここに追加
}

// NewControllers はユースケースを受け取って各コントローラを初期化して返す
func NewControllers(healthCheckUseCase usecase.UseCase) *Controllers {
	return &Controllers{
		HealthCheck: NewHealthCheckController(healthCheckUseCase),
		// 他コントローラ初期化もここに追記
	}
}
