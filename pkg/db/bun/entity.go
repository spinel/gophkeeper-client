package bun

import (
	"context"

	"github.com/spinel/gophkeeper-client/pkg/models"
)

// EntityRepo ...
type EntityRepo struct {
	db *DB
}

// NewEntityRepo ...
func NewEntityRepo(db *DB) *EntityRepo {
	return &EntityRepo{db: db}
}

// Init creates entity table
func (repo *EntityRepo) Init(ctx context.Context) error {
	_, err := repo.db.NewCreateTable().Model((*models.Entity)(nil)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
