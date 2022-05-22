package services

import (
	"context"
	"fmt"

	"github.com/spinel/gophkeeper-client/pkg/config"
	"github.com/spinel/gophkeeper-client/pkg/db"
	"github.com/spinel/gophkeeper-client/pkg/services/cli"
)

var Mgr *Manager

// Manager is just a collection of all services we have in the project
type Manager struct {
	User   UserService
	Config *config.Config
}

// NewManager creates new service manager
func NewManager(ctx context.Context, store *db.Store, cfg *config.Config) error {
	if store == nil {
		return fmt.Errorf("no store provided")
	}

	Mgr = &Manager{
		Config: cfg,
		User:   cli.NewUserWebService(ctx, store),
	}

	return nil
}
