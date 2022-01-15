package structs

import "github.com/deeean/go-vector/vector3"

type UserID string

type Quaternion struct {
	X, Y, Z, W float64
}

type TeardownPlayer struct {
	Position vector3.Vector3
	Rotation Quaternion // Quaternion
	Health   float32
}

type Presences map[UserID]*TeardownPlayer

type ServerConfig struct {
	Title    string `json:"name"`
	Gamemode string `json:"gamemode"`
	Version  string `json:"version"`
	Map      struct {
		Name        string  `yaml:"name"`
		SpawnPoints [][]int `yaml:"spawn_points"`
	}
}

type MatchState struct {
	Debug       bool
	Presences   Presences
	Map         string
	SpawnPoints []vector3.Vector3
}

var MState = &MatchState{
	Debug:       true,
	Presences:   make(Presences),
	Map:         "villa_gordon",
	SpawnPoints: nil,
}
