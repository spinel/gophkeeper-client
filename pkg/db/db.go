package db

import (
	"context"
	"fmt"
	"time"

	"github.com/spinel/gophkeeper-client/pkg/config"
	"github.com/spinel/gophkeeper-client/pkg/db/bun"
	"github.com/spinel/gophkeeper-client/pkg/logger"
)

// Store main struct
type Store struct {
	Bun *bun.DB

	User   UserRepo
	Entity EntityRepo
}

// New - create store
func New(cfg *config.Config) (*Store, error) {
	bunDB, err := bun.Dial(cfg)
	if err != nil {
		return nil, fmt.Errorf("pgdb.Dial failed: %w", err)
	}

	var store Store

	// Init Postgres repositories
	if bunDB != nil {
		store.Bun = bunDB

		go store.KeepAlivePg(cfg)

		store.User = bun.NewUserRepo(bunDB)
		store.Entity = bun.NewEntityRepo(bunDB)

		// Run migrations
		store.User.Init(context.Background())
		store.Entity.Init(context.Background())
	}

	return &store, nil
}

// KeepAlivePollPeriod is a Pg/MySQL keepalive check time period
const KeepAlivePollPeriod = 3

// KeepAlivePg makes sure PostgreSQL is alive and reconnects if needed
func (store *Store) KeepAlivePg(cfg *config.Config) {
	logger := logger.Get()
	var err error
	for {
		// Check if PostgreSQL is alive every 3 seconds
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false
		if store.Bun == nil {
			lostConnect = true
		} else if _, err = store.Bun.Exec("SELECT 1"); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		logger.Debug().Msg("[store.KeepAlivePg] Lost PostgreSQL connection. Restoring...")
		store.Bun, err = bun.Dial(cfg)
		if err != nil {
			logger.Err(err)
			continue
		}
		logger.Debug().Msg("[store.KeepAlivePg] PostgreSQL reconnected")
	}
}
