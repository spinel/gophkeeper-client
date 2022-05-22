package db

import (
	"context"

	"github.com/spinel/gophkeeper-client/pkg/models"
)

// UserRepo is a store for Users.
type UserRepo interface {
	Init(ctx context.Context) error
	Create(ctx context.Context, reqUser *models.User) (*models.User, error)
	Update(ctx context.Context, reqUser *models.User) (*models.User, error)
	GetAuthorised(ctx context.Context) (*models.User, error)
}

// EntityRepo is a store for entities.
type EntityRepo interface {
	Init(ctx context.Context) error
}
