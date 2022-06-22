package match

import (
	"fmt"
	"log"
	"time"

	"github.com/alexandargyurov/teardownM/match/structs"

	lua "github.com/yuin/gopher-lua"
)

func LuaLogGeneral(L *lua.LState) int {
	format := L.ToString(1)
	nargs := L.GetTop()
	args := make([]interface{}, nargs-1)
	for i := 2; i <= nargs; i++ {
		args[i-2] = L.ToString(i)
	}
	
	now := time.Now()
	
	fmt.Println("[" + now.Format("2006-01-02 15:04:05") + "] [Lua] [General] " + fmt.Sprintf(format, args...))
	
	return 0
}

func LuaGamemodeInit(L *lua.LState, serverConfig structs.ServerConfig) {	
	initFunction := L.GetGlobal("OnInitialize")
	
	if initFunction != lua.LNil {
		err := L.CallByParam(lua.P{
			Fn:      initFunction,
			NRet:    0,
			Protect: true,
		}); if err != nil {
			log.Fatalf("Error calling OnInitialize: %v", err)
			panic(err)
		}
	} else {
		fmt.Println("[Lua] [General] OnInitialize function not found")
	}
}

func LuaGamemodeOnJoin(L *lua.LState, userId string) {
	connectFunc := L.GetGlobal("OnConnected")
	if connectFunc != lua.LNil {
		err := L.CallByParam(lua.P{
			Fn:      connectFunc,
			NRet:    0,
			Protect: true,
		}, lua.LString(userId)); if err != nil {
			log.Fatalf("Error calling OnConnected: %v", err)
			panic(err)
		}
	}
}

func LuaGamemodeOnLeave(L *lua.LState, userId string) {
	disconnectFunc := L.GetGlobal("OnDisconnected")
	if disconnectFunc != lua.LNil {
		err := L.CallByParam(lua.P{
			Fn:      disconnectFunc,
			NRet:    0,
			Protect: true,
		}, lua.LString(userId)); if err != nil {
			log.Fatalf("Error calling OnDisconnected: %v", err)
			panic(err)
		}
	}
}