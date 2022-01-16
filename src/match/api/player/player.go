package player

import (
	"fmt"

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
}

// Given a user_id (UserID), returns the health of that player
func GetHealth(L *lua.LState) int {
	userID := L.ToString(1)
	teardownPlayer := structs.MState.Presences[structs.UserID(userID)]

	if exists(L, teardownPlayer, userID) {
		L.Push(lua.LNumber(teardownPlayer.Health))
	}

	return 1
}

// Given a user_id (UserID) and health (int), set the health of that player
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
