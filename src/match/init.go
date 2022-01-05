package match

import (
	"context"
	"database/sql"

	"github.com/deeean/go-vector/vector3"
	"github.com/heroiclabs/nakama-common/runtime"
)

type UserId string

type TeardownPlayer struct {
	Position vector3.Vector3 `json:"position"`
	Health   float32         `json:"health"`
}

type Presences map[UserId]*TeardownPlayer

type MatchState struct {
	debug     bool
	presences Presences
}

type Match struct{}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	state := &MatchState{
		debug:     true, // hardcode debug for now
		presences: make(Presences),
	}

	if state.debug {
		logger.Info("match init, starting with debug: %v", state.debug)
	}

	tickRate := 28
	label := "dev"

	return state, tickRate, label
}
