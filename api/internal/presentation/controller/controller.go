package controller

import usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/healthcheck"

// Controllers is a collection of controllers
type Controllers struct {
	// HealthCheck is a controller for the health check service
	HealthCheck *HealthCheckController
}

// NewControllers creates a new Controllers instance
func NewControllers(
	healthCheckUseCase usecase.UseCase,
) *Controllers {
	return &Controllers{
		HealthCheck: NewHealthCheckController(healthCheckUseCase),
	}
}
