package structs

import "github.com/deeean/go-vector/vector3"

type VehicleHandle int

type Quaternion struct {
	X, Y, Z, W float64
}

type TeardownPlayer struct {
	Position vector3.Vector3
	Rotation Quaternion
	Health   float32
}

type VehicleInstance struct {
	Position vector3.Vector3
	Rotation Quaternion
	Driver   string
	Health   float32
}

type Presences map[string]*TeardownPlayer
type Vehicles map[VehicleHandle]*VehicleInstance

type ServerConfig struct {
	Name       	string      	`yaml:"name"`
	Gamemode    string      	`yaml:"gamemode"`
	Debug       bool        	`yaml:"debug"`
	Map		 	string      	`yaml:"map"`
	SpawnPoints [][]float64 	`yaml:"spawn_points"`
}

type MatchState struct {
	Debug       bool
	Presences   Presences
	Vehicles    Vehicles
	Map         string
	SpawnPoints []vector3.Vector3
}

var MState = &MatchState{
	Debug:       true,
	Presences:   make(Presences),
	Vehicles:    make(Vehicles),
	Map:         "",
	SpawnPoints: nil,
}
