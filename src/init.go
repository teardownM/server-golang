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

	if err := initializer.RegisterRpc("go_echo_sample", rpc.Echo); err != nil {
		return err
	}
	if err := initializer.RegisterRpc("rpc_create_match", rpc.CreateMatch); err != nil {
		return err
	}

	if err := initializer.RegisterMatch("match", func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
		return &match.Match{}, nil
	}); err != nil {
		return err
	}

	if err := initializer.RegisterEvent(func(ctx context.Context, logger runtime.Logger, evt *api.Event) {
		logger.Info("Received event: %+v", evt)
	}); err != nil {
		return err
	}

	return nil
}
