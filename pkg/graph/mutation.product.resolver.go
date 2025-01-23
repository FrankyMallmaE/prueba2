package graph

import (
	"context"

	"github.com/levelstudio/payroll-4ta-crud/pkg/graph/model"
	"github.com/levelstudio/payroll-4ta-crud/pkg/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*models.Product, error) {
	product := models.Product{
		Name:  input.Name,
		Price: input.Price,
	}

	return r.ProductsRepo.CreateProduct(&product)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input model.UpdateProduct) (*models.Product, error) {
	product, err := r.ProductsRepo.GetProduct(id)
	if err != nil {
		return nil, err
	}

	if input.Price != nil {
		product.Price = *input.Price
	}

	if input.Name != nil {
		product.Name = *input.Name
	}

	return r.ProductsRepo.UpdateProduct(product)
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (*models.Product, error) {
	product, err := r.ProductsRepo.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return r.ProductsRepo.DeleteProduct(product)

}
