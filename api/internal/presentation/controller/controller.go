package controller

import usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/healthcheck"

type Controllers struct {
	HealthCheck *HealthCheckController
}

func NewControllers(
	healthCheckUseCase usecase.UseCase,
) *Controllers {
	return &Controllers{
		HealthCheck: NewHealthCheckController(healthCheckUseCase),
	}
}
