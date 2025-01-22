package graph

import (
	"context"

	"github.com/levelstudio/payroll-4ta-crud/pkg/models"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.UsersRepo.GetUsers()
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UsersRepo.GetUser(id)
}
