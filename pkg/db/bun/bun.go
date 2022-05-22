package bun

import (
	"database/sql"

	"github.com/spinel/gophkeeper-client/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*bun.DB
}

const notDeleted = "deleted_at is null"

// Dial creates new database connection to postgres
func Dial(cfg *config.Config) (*DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, cfg.DBUrl)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
