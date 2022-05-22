package services

import (
	"context"

	"github.com/spinel/gophkeeper-client/pkg/models"
)

type UserService interface {
	Create(ctx context.Context, userRegisterForm models.UserForm) (*models.User, error)
	SetToken(ctx context.Context, userRegisterForm models.UserForm) (*models.User, error)
	GetAuthorised(ctx context.Context) (*models.User, error)
}
