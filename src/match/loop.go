package match

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	PLAYER_MOVE   int64 = 1
	PLAYER_SPAWN  int64 = 2
	PLAYER_SHOOTS int64 = 3
	PLAYER_GRABS  int64 = 5
)

type IncomingData struct {
	UserId   string  `json:"user_id"`
	CurrentX float64 `json:"currentX"`
	CurrentY float64 `json:"currentY"`
	CurrentZ float64 `json:"currentZ"`
}

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	mState, _ := state.(*MatchState)

	for _, message := range messages {
		switch message.GetOpCode() {
		case PLAYER_MOVE:
			m_clientPresenceUserId := UserId(message.GetUserId())

			if _, ok := mState.presences[m_clientPresenceUserId]; ok {
				data := strings.Split(string(message.GetData()), ",")

				x, _ := strconv.ParseFloat(data[0], 64)
				y, _ := strconv.ParseFloat(data[1], 64)
				z, _ := strconv.ParseFloat(data[2], 64)
				rx, _ := strconv.ParseFloat(data[3], 64)
				ry, _ := strconv.ParseFloat(data[4], 64)
				rz, _ := strconv.ParseFloat(data[5], 64)
				rw, _ := strconv.ParseFloat(data[6], 64)

				mState.presences[m_clientPresenceUserId].Position.Set(x, y, z)
				mState.presences[m_clientPresenceUserId].Rotation.X = rx
				mState.presences[m_clientPresenceUserId].Rotation.Y = ry
				mState.presences[m_clientPresenceUserId].Rotation.Z = rz
				mState.presences[m_clientPresenceUserId].Rotation.W = rw

				dataToSend := message.GetUserId() + "," + data[0] + "," + data[1] + "," + data[2] + "," + data[3] + "," + data[4] + "," + data[5] + "," + data[6]

				// Sending nil for presenses means will send it to all players connected to the match
				dispatcher.BroadcastMessage(PLAYER_MOVE, []byte(dataToSend), nil, nil, true)
			}

		case PLAYER_SPAWN:
			dispatcher.BroadcastMessage(PLAYER_SPAWN, []byte(message.GetUserId()), nil, nil, true)
		case PLAYER_SHOOTS:
			m_clientPresenceUserId := UserId(message.GetUserId())

			if _, ok := mState.presences[m_clientPresenceUserId]; ok {
				tool := string(message.GetData())
				dataToSend := message.GetUserId() + "," + tool

				dispatcher.BroadcastMessage(PLAYER_SHOOTS, []byte(dataToSend), nil, nil, true)
			}
		default:
			fmt.Printf("Invalid OP Code!")
		}
	}

	return mState
}
