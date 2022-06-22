package match

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/deeean/go-vector/vector3"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"

	"github.com/alexandargyurov/teardownM/match/api/player"
	"github.com/alexandargyurov/teardownM/match/structs"

	lua "github.com/yuin/gopher-lua"
	"gopkg.in/yaml.v2"
)

var serverConfig structs.ServerConfig
var gamemodeConfig structs.GamemodeConfig

func ReadYMLConfig() structs.ServerConfig {
	content, fileErr := ioutil.ReadFile("./data/config.yml")
	if fileErr != nil {
		log.Fatal("Could not read config.yml")
		panic(fileErr)
	}

	config := structs.ServerConfig{}
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Unable to read server config: %v", err)
		panic(err)
	}
	
	return config
}

/**
* Registers server to global list of Teardown servers.
*
* NOTE: This is SUPER dangerous and requires getting teardownM
* developers to create API keys for them to register their server
* to the global list.
**/
func registerServer() {
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

var L = lua.NewState()

type Match struct{}

func WatchFiles(ctx context.Context) {
	gamemode := serverConfig.Gamemode
	path := "./gamemodes/" + gamemode + "/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	
	filemap := make(map[string]int64)
	for _, file := range files {
		filemap[file.Name()] = file.ModTime().Unix()
	}
	
	changed := false
	
	// Watch all files in the gamemode folder and reload them when they change
	for {
		// read all files in the gamemode folder, recursively
		err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			} else if info.IsDir() {
				return nil
			} else if filepath.Ext(path) != ".lua" {
				return nil
			}
			
			// Check if the file has changed
			if filemap[info.Name()] != info.ModTime().Unix() {
				filemap[info.Name()] = info.ModTime().Unix()
				changed = true
			}
			
			// Check for new files
			if _, ok := filemap[info.Name()]; !ok {
				filemap[info.Name()] = info.ModTime().Unix()
				changed = true
			}
			
			return nil
		}); if err != nil {
			log.Fatal(err)
		}
		
		if changed {
			fmt.Printf("Detected change in %s, reloading...\n", gamemode)
			
			// Reset the Lua state
			// SetupLua(ctx)
			
			if err := L.DoFile("./gamemodes/" + serverConfig.Gamemode + "/main.lua"); err != nil {
				fmt.Printf("Could not reload main.lua, error: %s\n", err.Error())
			} else {
				LuaGamemodeInit(L, serverConfig)	
			}
			
			changed = false
		}
	}
}

func SetupLua(ctx context.Context) {
	if L != nil {
		L.Close()
	}
	
	L = lua.NewState()
	
	L.SetContext(ctx)
	L.PreloadModule("player", player.ModuleLoader)
	L.SetGlobal("LogGeneral", L.NewFunction(LuaLogGeneral))
}

func GetGamemodeConfig() structs.GamemodeConfig {
	path := "./gamemodes/" + serverConfig.Gamemode + "/config.yml"
	content, fileErr := ioutil.ReadFile(path)
	if fileErr != nil {
		log.Fatal("Could not read config.yml")
		panic(fileErr)
	}
	
	config := structs.GamemodeConfig{}
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Unable to read gamemode config: %v", err)
		panic(err)
	}
	
	return config
}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	serverConfig = ReadYMLConfig()
	if !serverConfig.Debug {
		registerServer()
	}
	
	gamemodeConfig = GetGamemodeConfig()
	
	go WatchFiles(ctx)

	structs.MState.Map = serverConfig.Name
	structs.MState.Debug = serverConfig.Debug

	spawnPoints := make([]vector3.Vector3, 0)
	
	spawnPoint := gamemodeConfig.SpawnPoints[serverConfig.Map]
	for _, point := range spawnPoint {
		spawnPoints = append(spawnPoints, vector3.Vector3{point[0], point[1], point[2]})
	}

	structs.MState.SpawnPoints = spawnPoints
	
	SetupLua(ctx)
	
	if err := L.DoFile("./gamemodes/" + serverConfig.Gamemode + "/main.lua"); err != nil {
		fmt.Errorf("Could not find main.lua for gamemode '%s'. Make sure the folder name matches exactly the gamemode name. (no spaces)", serverConfig.Gamemode)
		panic(err)
	}
	
	LuaGamemodeInit(L, serverConfig)

	tickrate := serverConfig.Tickrate
	label := fmt.Sprintf("%s-%s", serverConfig.Name, serverConfig.Gamemode)

	return structs.MState, tickrate, label
}