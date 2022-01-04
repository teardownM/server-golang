package rpc

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Echo(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Info("RUNNING IN GO")
	return payload, nil
}

func CreateMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	matchID, err := nk.MatchCreate(ctx, "match", map[string]interface{}{})

	if err != nil {
		return "", err
	}

	return matchID, nil
}
