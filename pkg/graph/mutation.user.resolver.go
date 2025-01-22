package graph

import (
	"context"

	"github.com/levelstudio/payroll-4ta-crud/pkg/graph/model"
	"github.com/levelstudio/payroll-4ta-crud/pkg/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	user := models.User{
		Name: input.Name,
		Age:  input.Age,
	}

	return r.UsersRepo.CreateUser(&user)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUser) (*models.User, error) {
	user, err := r.UsersRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	if input.Age != nil {
		user.Age = *input.Age
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	return r.UsersRepo.UpdateUser(user)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	user, err := r.UsersRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return r.UsersRepo.DeleteUser(user)

}
