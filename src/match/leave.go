package match

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	mState, _ := state.(*MatchState)

	for _, presence := range presences {
		logger.Info("match leave username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId())
		delete(mState.presences, UserId(presence.GetUserId()))
	}

	return mState
}
