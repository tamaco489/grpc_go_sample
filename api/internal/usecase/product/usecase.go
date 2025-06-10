package product

import (
	"context"

	"github.com/tamaco489/grpc_go_sample/api/internal/domain/entity"
)

// UseCase defines the use case for product
type UseCase interface {
	GetProductList(ctx context.Context) ([]*entity.Product, error)
}

// useCase provides the implementation of UseCase
type useCase struct{}

// NewUseCase creates a new UseCase instance
func NewUseCase() UseCase {
	return &useCase{}
}

// GetProductList returns a list of products
func (u *useCase) GetProductList(ctx context.Context) ([]*entity.Product, error) {
	// NOTE: This is a mock implementation
	return []*entity.Product{
		{
			ID:          "1",
			Name:        "Product 1",
			Description: "Description 1",
			Price:       100,
			Categories:  []string{"Category 1"},
		},
		{
			ID:          "2",
			Name:        "Product 2",
			Description: "Description 2",
			Price:       200,
			Categories:  []string{"Category 2"},
		},
	}, nil
}
