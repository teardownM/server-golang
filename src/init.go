package main

import (
	"context"
	"database/sql"

	"github.com/alexandargyurov/teardownM/match"
	"github.com/alexandargyurov/teardownM/rpc"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	if err := initializer.RegisterRpc("rpc_get_matches", rpc.GetMatches); err != nil {
		return err
	}

	// Register as match handler, this call should be in InitModule.
	if err := initializer.RegisterMatch("sandbox", func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
		return &match.Match{}, nil
	}); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}

	if err := initializer.RegisterEvent(func(ctx context.Context, logger runtime.Logger, evt *api.Event) {
		logger.Info("Received event: %+v", evt)
	}); err != nil {
		return err
	}

	nk.MatchCreate(ctx, "sandbox", nil)

	return nil
}
