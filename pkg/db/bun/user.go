package bun

import (
	"context"

	"github.com/spinel/gophkeeper-client/pkg/models"
)

// UserPgRepo ...
type UserPgRepo struct {
	db *DB
}

// NewUserRepo ...
func NewUserRepo(db *DB) *UserPgRepo {
	return &UserPgRepo{db: db}
}

// Init creates User table
func (repo *UserPgRepo) Init(ctx context.Context) error {
	_, err := repo.db.NewCreateTable().Model((*models.User)(nil)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Create creates User
func (repo *UserPgRepo) Create(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := repo.db.NewInsert().
		Model(user).
		Returning("*").
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update User
func (repo *UserPgRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := repo.db.NewUpdate().
		Model(user).
		Column("token").
		Where("email=?", user.Email).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAuthorised User
func (repo *UserPgRepo) GetAuthorised(ctx context.Context) (*models.User, error) {
	var user models.User
	err := repo.db.NewSelect().
		Model(&user).
		Where("token != ''").
		Limit(1).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetByLogin retrieve user by login.
func (repo *UserPgRepo) GetByLogin(ctx context.Context, login string) (*models.User, error) {
	var user models.User
	err := repo.db.NewSelect().
		Model(&user).
		Where("login = ?", login).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
