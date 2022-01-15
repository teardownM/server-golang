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

	lua "github.com/yuin/gopher-lua"
	"gopkg.in/yaml.v2"
)

type UserId string

type Quaternion struct {
	X, Y, Z, W float64
}

type TeardownPlayer struct {
	Position vector3.Vector3
	Rotation Quaternion // Quaternion
	Health   float32
}

type Presences map[UserId]*TeardownPlayer

type MatchState struct {
	debug       bool
	presences   Presences
	_map        string
	spawnPoints []vector3.Vector3
	luaState    *lua.LState
	onJoin      func(*lua.LState)
}

type Match struct{}

type ServerConfig struct {
	Title    string `json:"name"`
	Gamemode string `json:"gamemode"`
	Version  string `json:"version"`
	Map      struct {
		Name        string  `yaml:"name"`
		SpawnPoints [][]int `yaml:"spawn_points"`
	}
}

/**
* Reads the config file in ./modules/config.yml
**/
func readYMLConfig() ServerConfig {
	content, fileErr := ioutil.ReadFile("./data/gamemodes/config.yml")
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
**/
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

var L = lua.NewState()

var mState = &MatchState{
	debug:       true, // hardcode debug for now
	presences:   make(Presences),
	_map:        "villa_gordon",
	spawnPoints: nil,
	luaState:    L,
}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	serverConfig := readYMLConfig()
	registerServer(serverConfig)

	L.SetContext(ctx)
	L.PreloadModule("player", Loader)

	spawnPoints := make([]vector3.Vector3, 3)
	spawnPoints[0] = *vector3.New(135, 9, -72)
	spawnPoints[1] = *vector3.New(135, 8, -66)
	spawnPoints[2] = *vector3.New(125, 8, -66)
	mState.spawnPoints = spawnPoints

	LuaGamemodeInit(L, serverConfig)

	if mState.debug {
		logger.Info("match init, starting with debug: %v", mState.debug)
	}

	tickRate := 28
	label := "dev"

	return mState, tickRate, label
}
