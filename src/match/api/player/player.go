package player

import (
	"github.com/alexandargyurov/teardownM/match/structs"

	lua "github.com/yuin/gopher-lua"
)

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
}

// Given a user_id, returns the health of that player
func GetHealth(L *lua.LState) int {
	userID := L.ToString(1)
	teardownPlayer := structs.MState.Presences[structs.UserID(userID)]

	if teardownPlayer != nil {
		L.Push(lua.LNumber(teardownPlayer.Health))
	} else {
		L.Push(lua.LString("Error: No player found with id of " + userID))
	}

	return 1
}
