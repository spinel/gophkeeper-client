package main

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spinel/gophkeeper-client/cmd"
	"github.com/spinel/gophkeeper-client/pkg/config"
	"github.com/spinel/gophkeeper-client/pkg/db"
	"github.com/spinel/gophkeeper-client/pkg/logger"
	"github.com/spinel/gophkeeper-client/pkg/services"
)

func main() {
	l := logger.Get()
	if err := run(l); err != nil {
		l.Fatal().Msgf("init error, %s", err.Error())
	}

	cmd.Execute()
}

func run(l *logger.Logger) error {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		return errors.Wrap(err, "config load error")
	}

	db, err := db.New(&cfg)
	if err != nil {
		return errors.Wrap(err, "store")
	}

	err = services.NewManager(ctx, db, &cfg)
	if err != nil {
		return errors.Wrap(err, "manager")
	}

	return nil
}
