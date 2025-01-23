package graph

import (
	"context"

	"github.com/levelstudio/payroll-4ta-crud/pkg/models"
)

func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	return r.ProductsRepo.GetProducts()
}

func (r *queryResolver) Product(ctx context.Context, id string) (*models.Product, error) {
	return r.ProductsRepo.GetProduct(id)
}
