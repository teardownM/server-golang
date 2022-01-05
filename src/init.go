package main

import (
	"context"
	"database/sql"

	"teardownNakamaServer/match"
	"teardownNakamaServer/rpc"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {

	// TODO: Remove me once we build server browser
	if err := initializer.RegisterRpc("rpc_get_match_id", rpc.CreateMatch); err != nil {
		return err
	}

	// if err := initializer.RegisterMatch("match", func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	// 	return &match.Match{}, nil
	// }); err != nil {
	// 	return err
	// }

	// Register as match handler, this call should be in InitModule.
	if err := initializer.RegisterMatch("pingpong", func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
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

	return nil
}
