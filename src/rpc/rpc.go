package rpc

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func GetMatches(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	matches, err := nk.MatchList(ctx, 10, true, "dev", nil, nil, "")
	if err != nil {
		logger.WithField("err", err).Error("Match list error.")
		return "", err
	} else {
		for _, match := range matches {
			logger.Info("Found match with id: %s", match.GetMatchId())
		}

		jsonMatches, _ := json.Marshal(matches)
		return string(jsonMatches), nil
	}
}
