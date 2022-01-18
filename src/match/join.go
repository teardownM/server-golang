package match

import (
	"context"
	"database/sql"
	"math/rand"

	"github.com/deeean/go-vector/vector3"
	"github.com/heroiclabs/nakama-common/runtime"

	"github.com/alexandargyurov/teardownM/match/structs"
)

func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	logger.Info("match join attempt username %v user_id %v session_id %v node %v with metadata %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId(), metadata)

	return state, true, ""
}

func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {

	for _, presence := range presences {
		userId := structs.UserID(presence.GetUserId())

		logger.Info("match join username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId())

		quaternion := structs.Quaternion{X: 0, Y: 0, Z: 0, W: 0}

		playerSpawnPoint := structs.MState.SpawnPoints[rand.Intn(len(structs.MState.SpawnPoints))]

		structs.MState.Presences[userId] = &structs.TeardownPlayer{Position: *vector3.New(playerSpawnPoint.X, playerSpawnPoint.Y, playerSpawnPoint.Z), Rotation: quaternion, Health: 100}

		//dataToSend := fmt.Sprintf("%f", playerSpawnPoint.X) + "," + fmt.Sprintf("%f", playerSpawnPoint.Y) + "," + fmt.Sprintf("%f", playerSpawnPoint.Z)
		clientRecipients := make([]runtime.Presence, 1)
		clientRecipients[0] = presence

		//dispatcher.BroadcastMessage(PLAYER_SPAWN, []byte(dataToSend), clientRecipients, nil, true)

		LuaGamemodeOnJoin(L, userId)
	}

	return structs.MState
}
