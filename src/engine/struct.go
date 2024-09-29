package engine

import (
	"main/src/entity"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
)

type engine int

const (
	INEXPLORATION  engine = iota
	PAUSE    engine = iota
	GAMEOVER engine = iota
	INCOMBAT engine = iota
)

type Engine struct {
	Potion  item.Item
	Player   entity.Player
	Monsters []entity.Monster
	Marchant []entity.Marchant

	Music       rl.Music
	MusicVolume float32

	Sprites map[string]rl.Texture2D
	LoadMenu rl.Texture2D
	LoadPause rl.Texture2D


	Camera rl.Camera2D

	MapJSON MapJSON

	IsRunning   bool
	StateMenu   menu
	StateEngine engine
}
