package match

import (
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	//L.SetField(mod, "name", lua.LString("value"))

	// returns the module
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"GetHealth": GetHealth,
	"Test":      Test,
}

func Test(L *lua.LState) int {
	L.Push(lua.LString("test"))
	return 1
}

// User ID
func GetHealth(L *lua.LState) int {
	userId := L.ToString(1)
	teardownPlayer := mState.presences[UserId(userId)]

	if teardownPlayer != nil {
		L.Push(lua.LNumber(teardownPlayer.Health))
	} else {
		L.Push(lua.LString("Error: No player found with id of " + userId))
	}

	return 1
}
