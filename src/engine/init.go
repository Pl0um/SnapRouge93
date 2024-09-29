package engine

import (
	"main/src/entity"
	"main/src/item"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1200
)



func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	// Initialisation des composants du jeu
	e.InitItems()
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMap("textures/map/tilesets/map.json")

}
func (e *Engine) InitEntities() {

	e.Marchant = append(e.Marchant, entity.Marchant {
		Name: "renard",
		Position: rl.Vector2{X: 300, Y: 1000},
		Health: 5,
		Damage: 5,
		Worth: 1,

		IsAlive: true,
		Sprite: rl.LoadTexture("textures/entities/soldier/Soldier-Idle.png"),
	})

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 100, Y: 770},
		Health:    100,
		Money :    1,
		Speed:     2,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite: rl.LoadTexture("textures/entities/Knight/IDLE.png"),
	
	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 250, Y: 785},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")

	rl.PlayMusicStream(e.Music)
}

func (e *Engine) InitItems() {
	Potion := item.Item {
		Name:   "potion",
		Price:   3,
	}
	e.Player.Inventory = append(e.Player.Inventory, Potion)
}