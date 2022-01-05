package match

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	PLAYER_POS    int64 = 1
	PLAYER_SPAWN  int64 = 2
	PLAYER_SHOOTS int64 = 3
)

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	for _, message := range messages {
		switch message.GetOpCode() {
		case PLAYER_POS:
			// A player has moved
			fmt.Println("Player is moving")

			fmt.Println(string(message.GetData())) // Look at string builder for performance?

			// local client_data = nk.json_decode(message.data)
			// local client_presence = state.presences[client_data.user_id] -- find the user in presences

			// if client_presence then
			//     client_presence.x = client_data.currentX
			//     client_presence.z = client_data.currentZ
			//     client_presence.y = client_data.currentY
			// end

			// Sending nil for presenses means will send it to all players connected to the match
			dispatcher.BroadcastMessage(PLAYER_POS, message.GetData(), nil, nil, true)
		case PLAYER_SPAWN:
			fmt.Println("Linux.")
		case PLAYER_SHOOTS:
			fmt.Println("Linux.")
		default:
			fmt.Printf("Invalid OP Code!")
		}
	}

	return state
}
