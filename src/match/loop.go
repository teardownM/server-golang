package match

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	PLAYER_POS    int64 = 1
	PLAYER_SPAWN  int64 = 2
	PLAYER_SHOOTS int64 = 3
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
		case PLAYER_POS:
			m_clientPresenceUserId := UserId(message.GetUserId())

			if _, ok := mState.presences[m_clientPresenceUserId]; ok {
				s := message.GetData()
				data := IncomingData{}
				json.Unmarshal([]byte(s), &data)

				mState.presences[m_clientPresenceUserId].Position.Set(data.CurrentX, data.CurrentY, data.CurrentZ)
			}

			data, _ := json.Marshal(&mState.presences)

			// Sending nil for presenses means will send it to all players connected to the match
			dispatcher.BroadcastMessage(PLAYER_POS, data, nil, nil, true)
		case PLAYER_SPAWN:
			fmt.Println("Linux.")
		case PLAYER_SHOOTS:
			fmt.Println("Linux.")
		default:
			fmt.Printf("Invalid OP Code!")
		}
	}

	return mState
}
