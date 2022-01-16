package player

import (
	"fmt"
	"strconv"

	"github.com/alexandargyurov/teardownM/match/structs"

	lua "github.com/yuin/gopher-lua"
)

// Checks if object exists
func exists(L *lua.LState, teardownPlayer *structs.TeardownPlayer, userID string) bool {
	if teardownPlayer != nil {
		return true
	} else {
		fmt.Println("Error: No player found with id of " + userID)
		L.Push(lua.LBool(false))
		return false
	}
}

func ModuleLoader(L *lua.LState) int {
	// register functions to the table
	module := L.SetFuncs(L.NewTable(), exports)

	// register other stuff
	//L.SetField(mod, "name", lua.LString("value"))

	L.Push(module)
	return 1
}

var exports = map[string]lua.LGFunction{
	"GetHealth": GetHealth,
	"SetHealth": SetHealth,
	"GetPos":    GetPos,
	"SetPos":    SetPos,
}

// Given a user_id (UserID), return the health of that player
func GetHealth(L *lua.LState) int {
	userID := L.ToString(1)
	teardownPlayer := structs.MState.Presences[structs.UserID(userID)]

	if exists(L, teardownPlayer, userID) {
		L.Push(lua.LNumber(teardownPlayer.Health))
	}

	return 1
}

// Given a user_id (UserID) and health (number), set the health of that player
func SetHealth(L *lua.LState) int {
	userID := L.ToString(1)
	newHealth := L.ToInt(2)

	teardownPlayer := structs.MState.Presences[structs.UserID(userID)]

	if exists(L, teardownPlayer, userID) {
		teardownPlayer.Health = float32(newHealth)
		L.Push(lua.LBool(true))
	}

	return 1
}

// Given a user_id (UserID), return the position table of that player
func GetPos(L *lua.LState) int {
	userID := L.ToString(1)
	teardownPlayer := structs.MState.Presences[structs.UserID(userID)]

	if exists(L, teardownPlayer, userID) {
		playerPosTable := &lua.LTable{}
		playerPosTable.RawSetString("X", lua.LNumber(teardownPlayer.Position.X))
		playerPosTable.RawSetString("Y", lua.LNumber(teardownPlayer.Position.Y))
		playerPosTable.RawSetString("Z", lua.LNumber(teardownPlayer.Position.Z))

		L.Push(playerPosTable)
	}

	return 1
}

// Given a user_id (UserID), X, Y, Z (number), set the position for that player
func SetPos(L *lua.LState) int {
	userID := L.ToString(1)
	newX := L.ToString(2)
	newY := L.ToString(3)
	newZ := L.ToString(4)

	teardownPlayer := structs.MState.Presences[structs.UserID(userID)]

	if exists(L, teardownPlayer, userID) {
		x, _ := strconv.ParseFloat(newX, 64)
		y, _ := strconv.ParseFloat(newY, 64)
		z, _ := strconv.ParseFloat(newZ, 64)

		teardownPlayer.Position.X = x
		teardownPlayer.Position.X = y
		teardownPlayer.Position.X = z

		L.Push(lua.LBool(true))
	}

	return 1
}
