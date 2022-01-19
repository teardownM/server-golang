package match

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexandargyurov/teardownM/match/structs"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	PLAYER_MOVE        int64 = 1
	PLAYER_SPAWN       int64 = 2
	PLAYER_SHOOTS      int64 = 3
	PLAYER_GRABS       int64 = 5
	PLAYER_TOOL_CHANGE int64 = 6
)

type IncomingData struct {
	UserId   string  `json:"user_id"`
	CurrentX float64 `json:"currentX"`
	CurrentY float64 `json:"currentY"`
	CurrentZ float64 `json:"currentZ"`
}

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	for _, message := range messages {
		switch message.GetOpCode() {
		case PLAYER_MOVE:
			m_clientPresenceUserId := structs.UserID(message.GetUserId())

			if _, ok := structs.MState.Presences[m_clientPresenceUserId]; ok {
				data := strings.Split(string(message.GetData()), ",")

				x, _ := strconv.ParseFloat(data[0], 64)
				y, _ := strconv.ParseFloat(data[1], 64)
				z, _ := strconv.ParseFloat(data[2], 64)
				rx, _ := strconv.ParseFloat(data[3], 64)
				ry, _ := strconv.ParseFloat(data[4], 64)
				rz, _ := strconv.ParseFloat(data[5], 64)
				rw, _ := strconv.ParseFloat(data[6], 64)
				
				tx, _ := strconv.ParseFloat(data[7], 64)
				ty, _ := strconv.ParseFloat(data[8], 64)
				tz, _ := strconv.ParseFloat(data[9], 64)
				tr_x, _ := strconv.ParseFloat(data[10], 64)
				tr_y, _ := strconv.ParseFloat(data[11], 64)
				tr_z, _ := strconv.ParseFloat(data[12], 64)
				tr_w, _ := strconv.ParseFloat(data[13], 64)

				structs.MState.Presences[m_clientPresenceUserId].Position.Set(x, y, z)
				structs.MState.Presences[m_clientPresenceUserId].ToolPosition.Set(tx, ty, tz)
				structs.MState.Presences[m_clientPresenceUserId].Rotation.X = rx
				structs.MState.Presences[m_clientPresenceUserId].Rotation.Y = ry
				structs.MState.Presences[m_clientPresenceUserId].Rotation.Z = rz
				structs.MState.Presences[m_clientPresenceUserId].Rotation.W = rw
				structs.MState.Presences[m_clientPresenceUserId].ToolRotation.X = tr_x
				structs.MState.Presences[m_clientPresenceUserId].ToolRotation.Y = tr_y
				structs.MState.Presences[m_clientPresenceUserId].ToolRotation.Z = tr_z
				structs.MState.Presences[m_clientPresenceUserId].ToolRotation.W = tr_w

				dataToSend := message.GetUserId() + "," + data[0] + "," + data[1] + "," + data[2] + "," + data[3] + "," + data[4] + "," + data[5] + "," + data[6] + "," + data[7] + "," + data[8] + "," + data[9] + "," + data[10] + "," + data[11] + "," + data[12] + "," + data[13]

				// Sending nil for presenses means will send it to all players connected to the match
				dispatcher.BroadcastMessage(PLAYER_MOVE, []byte(dataToSend), nil, nil, true)
			}

		case PLAYER_SPAWN:
			dispatcher.BroadcastMessage(PLAYER_SPAWN, []byte(message.GetUserId()), nil, nil, true)
		case PLAYER_SHOOTS:
			m_clientPresenceUserId := structs.UserID(message.GetUserId())

			if _, ok := structs.MState.Presences[m_clientPresenceUserId]; ok {
				tool := string(message.GetData())
				dataToSend := message.GetUserId() + "," + tool
				// This user has shot (M1) with this tool ^^
				dispatcher.BroadcastMessage(PLAYER_SHOOTS, []byte(dataToSend), nil, nil, true)
			}
		case PLAYER_TOOL_CHANGE:
			m_clientPresenceUserId := structs.UserID(message.GetUserId())

			if _, ok := structs.MState.Presences[m_clientPresenceUserId]; ok {
				tool := string(message.GetData())
				dataToSend := message.GetUserId() + "," + tool
				// This user has changed to this tool ^^
				dispatcher.BroadcastMessage(PLAYER_TOOL_CHANGE, []byte(dataToSend), nil, nil, true)
			}

		default:
			fmt.Printf("Invalid OP Code!")
		}
	}

	return structs.MState
}
