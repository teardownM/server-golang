package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexandargyurov/teardownM/match"
	"github.com/alexandargyurov/teardownM/rpc"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"

	supabase "github.com/nedpals/supabase-go"
	lua "github.com/yuin/gopher-lua"
)

type ServerConfig struct {
	Title    string `json:"name"`
	Gamemode string `json:"gamemode"`
	Version  string `json:"version"`
	Map      struct {
		Name        string  `yaml:"name"`
		SpawnPoints [][]int `yaml:"spawn_points"`
	}
}

func readYMLConfig() ServerConfig {
	content, fileErr := ioutil.ReadFile("./data/modules/config.yml")
	if fileErr != nil {
		log.Fatal("Could not read config.yml")
		panic(fileErr)
	}

	serverConfig := ServerConfig{}
	err := yaml.Unmarshal(content, &serverConfig)
	if err != nil {
		log.Fatalf("Unable to read server config: %v", err)
		panic(err)
	}

	return serverConfig
}

/**
* Registers server to global list of Teardown servers.
*
* NOTE: This is SUPER dangerous and requires getting teardownM
* developers to create API keys for them to register their server
* to the global list.
 */
func registerServer(serverConfig ServerConfig) {
	err := godotenv.Load("./data/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabaseClient := supabase.CreateClient(supabaseUrl, supabaseKey, true)

	var results map[string]interface{}
	supabaseClient.DB.From("servers").Insert(serverConfig).Execute(results)
}

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	serverConfig := readYMLConfig()
	registerServer(serverConfig)

	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("./data/modules/" + serverConfig.Gamemode + "/init.lua"); err != nil {
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

	if err := initializer.RegisterRpc("rpc_get_matches", rpc.GetMatches); err != nil {
		return err
	}

	// Register as match handler, this call should be in InitModule.
	if err := initializer.RegisterMatch("sandbox", func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
		return &match.Match{}, nil
	}); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}

	if err := initializer.RegisterEvent(func(ctx context.Context, logger runtime.Logger, evt *api.Event) {
		logger.Info("Received event: %+v", evt)
	}); err != nil {
		return err
	}

	nk.MatchCreate(ctx, "sandbox", nil)

	return nil
}
