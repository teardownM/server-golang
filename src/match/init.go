package match

import (
	"context"
	"database/sql"

	"github.com/deeean/go-vector/vector3"
	"github.com/heroiclabs/nakama-common/runtime"
)

type TeardownPlayer struct {
	position vector3.Vector3
	health   float32
}

type UserId string

type MatchState struct {
	debug     bool
	presences map[UserId]TeardownPlayer
}

type Match struct{}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	state := &MatchState{
		debug: true, // hardcode debug for now
	}

	if state.debug {
		logger.Info("match init, starting with debug: %v", state.debug)
	}

	tickRate := 28
	label := "sandbox"

	return state, tickRate, label
}
