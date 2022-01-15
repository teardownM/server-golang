package match

import (
	"context"
	"database/sql"
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

/**
* Reads the config file in ./modules/config.yml
**/
func readYMLConfig() structs.ServerConfig {
	content, fileErr := ioutil.ReadFile("./data/gamemodes/config.yml")
	if fileErr != nil {
		log.Fatal("Could not read config.yml")
		panic(fileErr)
	}

	serverConfig := structs.ServerConfig{}
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
**/
func registerServer(serverConfig structs.ServerConfig) {
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

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	serverConfig := readYMLConfig()
	registerServer(serverConfig)

	L.SetContext(ctx)
	L.PreloadModule("player", player.ModuleLoader)

	spawnPoints := make([]vector3.Vector3, 3)
	spawnPoints[0] = *vector3.New(135, 9, -72)
	spawnPoints[1] = *vector3.New(135, 8, -66)
	spawnPoints[2] = *vector3.New(125, 8, -66)
	structs.MState.SpawnPoints = spawnPoints

	LuaGamemodeInit(L, serverConfig)

	if structs.MState.Debug {
		logger.Info("match init, starting with debug: %v", structs.MState.Debug)
	}

	tickRate := 28
	label := "dev"

	return structs.MState, tickRate, label
}
