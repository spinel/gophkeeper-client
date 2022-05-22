package cli

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spinel/gophkeeper-client/pkg/db"
	"github.com/spinel/gophkeeper-client/pkg/models"
)

// UserWebService ...
type UserWebService struct {
	ctx   context.Context
	store *db.Store
}

// NewUserWebService is a user service
func NewUserWebService(ctx context.Context, store *db.Store) *UserWebService {
	return &UserWebService{
		ctx:   ctx,
		store: store,
	}
}

// Create user service
func (svc UserWebService) Create(ctx context.Context, userRegisterForm models.UserForm) (*models.User, error) {
	user := &models.User{
		Email:    userRegisterForm.Email,
		Password: userRegisterForm.Password,
	}

	user, err := svc.store.User.Create(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Create error")
	}

	return user, nil
}

// SetToken user service
func (svc UserWebService) SetToken(ctx context.Context, userRegisterForm models.UserForm) (*models.User, error) {
	user := &models.User{
		Email:    userRegisterForm.Email,
		Password: userRegisterForm.Password,
		Token:    userRegisterForm.Token,
	}

	user, err := svc.store.User.Update(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "svc.User.Create error")
	}

	return user, nil
}

// GetAuthorised user service
func (svc UserWebService) GetAuthorised(ctx context.Context) (*models.User, error) {
	user, err := svc.store.User.GetAuthorised(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "svc.User.GetAuthorised error")
	}

	return user, nil
}
