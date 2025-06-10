package controller

import (
	healthcheck_usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/healthcheck"
	product_usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/product"
)

// Controllers is a collection of controllers
type Controllers struct {
	// HealthCheck is a controller for the health check service
	HealthCheck *HealthCheckController

	// Product is a controller for the product service
	Product *ProductController
}

// NewControllers creates a new Controllers instance
func NewControllers(
	healthCheckUseCase healthcheck_usecase.UseCase,
	productUseCase product_usecase.UseCase,
) *Controllers {
	return &Controllers{
		HealthCheck: NewHealthCheckController(healthCheckUseCase),
		Product:     NewProductController(productUseCase),
	}
}
