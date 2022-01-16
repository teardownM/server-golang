package match

import (
	"fmt"
	"log"

	"github.com/alexandargyurov/teardownM/match/structs"

	lua "github.com/yuin/gopher-lua"
)

func LuaGamemodeInit(L *lua.LState, serverConfig structs.ServerConfig) {
	if err := L.DoFile("./data/gamemodes/" + serverConfig.Gamemode + "/init.lua"); err != nil {
		log.Fatalf("Could not find init.lua for gamemode " + serverConfig.Gamemode + ". Make sure the folder name matches exactly the gamemode name. (no spaces)")
		panic(err)
	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("Init"),
		NRet:    1,
		Protect: true,
	}, lua.LString(serverConfig.Gamemode)); err != nil {
		panic(err)
	}

	// Get the returned value from the stack and cast it to a lua.LString
	if str, ok := L.Get(-1).(lua.LString); ok {
		fmt.Println(str)
	}

	// Pop the returned value from the stack
	L.Pop(1)
}

func LuaGamemodeOnJoin(L *lua.LState, userId structs.UserID) {
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("OnJoin"),
		NRet:    0,
		Protect: true,
	}, lua.LString(userId)); err != nil {
		panic(err)
	}
}

func LuaGamemodeOnLeave(L *lua.LState, userId structs.UserID) {
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("OnLeave"),
		NRet:    0,
		Protect: true,
	}, lua.LString(userId)); err != nil {
		panic(err)
	}
}
