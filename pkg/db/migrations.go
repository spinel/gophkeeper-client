package db

import (
	"fmt"

	"github.com/spinel/gophkeeper-client/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// runMigrations runs migrations.
func runMigrations() error {

		

	cfg, _ := config.LoadConfig()
	if cfg.MigrationsPath == "" {
		return nil
	}
	if cfg.DBUrl == "" {
		return errors.New("no cfg.DBURL provided")
	}
	m, err := migrate.New(
		cfg.MigrationsPath,
		fmt.Sprintf("sqlite3://%s", cfg.DBUrl),
	)
	if err != nil {

		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
