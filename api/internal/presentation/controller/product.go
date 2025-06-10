package controller

import (
	"context"

	product "github.com/tamaco489/grpc_go_sample/api/internal/gen/proto/product"
	usecase "github.com/tamaco489/grpc_go_sample/api/internal/usecase/product"
)

// ProductController is a struct that implements the gRPC Product service
type ProductController struct {
	usecase usecase.UseCase

	product.UnimplementedProductServiceServer
}

// NewProductController creates a new ProductController instance
func NewProductController(u usecase.UseCase) *ProductController {
	return &ProductController{usecase: u}
}

// GetProductList implements the gRPC Product service method
func (c *ProductController) GetProductList(ctx context.Context, req *product.GetProductListRequest) (*product.GetProductListResponse, error) {
	// Delegate the processing to the usecase layer
	products, err := c.usecase.GetProductList(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the entity to the gRPC response
	productsPb := make([]*product.Product, len(products))
	for i, p := range products {
		productsPb[i] = &product.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Categories:  p.Categories,
		}
	}
	return &product.GetProductListResponse{
		Products: productsPb,
	}, nil
}
