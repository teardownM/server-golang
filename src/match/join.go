package match

import (
	"context"
	"database/sql"

	"github.com/deeean/go-vector/vector3"
	"github.com/heroiclabs/nakama-common/runtime"
)

func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	logger.Info("match join attempt username %v user_id %v session_id %v node %v with metadata %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId(), metadata)

	return state, true, ""
}

func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	mState, _ := state.(*MatchState)

	for _, presence := range presences {
		logger.Info("match join username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId())

		// This line means the player will always spawn at 0, 0, 0
		// In the future, add support for players previous location (database required) or map spawn points
		mState.presences[UserId(presence.GetUserId())] = &TeardownPlayer{*vector3.New(0, 0, 0), 100}
		data := presence.GetUserId() + ",0," + "0," + "0"

		dispatcher.BroadcastMessage(PLAYER_JOINS, []byte(data), nil, nil, true)
	}

	return mState
}
