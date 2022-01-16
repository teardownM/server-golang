package match

import (
	"context"
	"database/sql"

	"github.com/alexandargyurov/teardownM/match/structs"

	"github.com/heroiclabs/nakama-common/runtime"
)

func (m *Match) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	if structs.MState.Debug {
		logger.Info("match signal match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick)
		logger.Info("match signal match_id %v data %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), data)
	}
	return state, "signal received: " + data
}
