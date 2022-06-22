package match

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

/**
* Reads the config file in ./modules/config.yml
**/
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

func WatchFiles() {
	gamemode := serverConfig.Gamemode
	path := "./data/gamemodes/" + gamemode + "/"
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
		for file, lastModified := range filemap {
			filepath := path + file
			fileInfo, err := os.Stat(filepath)
			if err != nil {
				log.Fatal(err)
			}
			
			if fileInfo.ModTime().Unix() > lastModified {
				changed = true
				filemap[file] = fileInfo.ModTime().Unix()
			}
		
			// Check for new or deleted files
			if !fileInfo.IsDir() {
				if _, ok := filemap[file]; !ok {
					changed = true
					filemap[file] = fileInfo.ModTime().Unix()
				}
				
				if _, ok := filemap[file]; ok {
					if fileInfo.ModTime().Unix() > filemap[file] {
						changed = true
						filemap[file] = fileInfo.ModTime().Unix()
					}
				}
				
				if _, ok := filemap[file]; ok {
					if fileInfo.ModTime().Unix() < filemap[file] {
						changed = true
						delete(filemap, file)
					}
				}
			}
		}
		
		if changed {
			if err := L.DoFile("./data/gamemodes/" + serverConfig.Gamemode + "/main.lua"); err != nil {
				fmt.Printf("Could not reload main.lua for gamemode %s", serverConfig.Gamemode)
				panic(err)
			}
			
			LuaGamemodeInit(L, serverConfig)
			
			changed = false
		}
	}
}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	serverConfig = ReadYMLConfig()
	if !serverConfig.Debug {
		registerServer()
	}
	
	go WatchFiles()

	L.SetContext(ctx)
	L.PreloadModule("player", player.ModuleLoader)

	structs.MState.Map = serverConfig.Name
	structs.MState.Debug = serverConfig.Debug

	spawnPoints := make([]vector3.Vector3, 3)
	for index, spawnPoint := range serverConfig.SpawnPoints {
		spawnPoints[index] = *vector3.New(spawnPoint[0], spawnPoint[1], spawnPoint[2])
	}

	structs.MState.SpawnPoints = spawnPoints

	if err := L.DoFile("./data/gamemodes/" + serverConfig.Gamemode + "/main.lua"); err != nil {
		log.Fatalf("Could not find main.lua for gamemode " + serverConfig.Gamemode + ". Make sure the folder name matches exactly the gamemode name. (no spaces)")
		panic(err)
	}
	
	L.SetGlobal("LogGeneral", L.NewFunction(LuaLogGeneral))
	
	LuaGamemodeInit(L, serverConfig)

	if structs.MState.Debug {
		logger.Info("match init, starting with debug: %v", structs.MState.Debug)
	}

	tickRate := 28
	label := serverConfig.Gamemode

	return structs.MState, tickRate, label
}
