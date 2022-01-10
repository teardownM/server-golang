package match

import (
	"context"
	"database/sql"

	"github.com/deeean/go-vector/vector3"
	"github.com/heroiclabs/nakama-common/runtime"
)

type UserId string

type Quaternion struct {
	X, Y, Z, W float64
}

type TeardownPlayer struct {
	Position vector3.Vector3
	Rotation Quaternion // Quaternion
	Health   float32
}

type Presences map[UserId]*TeardownPlayer

type MatchState struct {
	debug       bool
	presences   Presences
	_map        string
	spawnPoints []vector3.Vector3
}

type Match struct{}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	spawnPoints := make([]vector3.Vector3, 3)
	spawnPoints[0] = *vector3.New(135, 9, -72)
	spawnPoints[1] = *vector3.New(135, 8, -66)
	spawnPoints[2] = *vector3.New(125, 8, -66)

	state := &MatchState{
		debug:       true, // hardcode debug for now
		presences:   make(Presences),
		_map:        "villa_gordon",
		spawnPoints: spawnPoints,
	}

	if state.debug {
		logger.Info("match init, starting with debug: %v", state.debug)
	}

	tickRate := 28
	label := "dev"

	return state, tickRate, label
}
