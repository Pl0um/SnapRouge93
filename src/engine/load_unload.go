package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	// Chargement des textures du personnage
	e.LoadMenu = rl.LoadTexture("textures/img3.jpg")
	e.LoadPause = rl.LoadTexture("textures/img4.jpg")

}

func (e *Engine) Unload() {
	// On libère les textures chargées, le joueur, la map, les monstres, etc...
	rl.UnloadTexture(e.Player.Sprite)
	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}

	for _, Marchants := range e.Marchant {
		rl.UnloadTexture(Marchants.Sprite)
	}
	rl.UnloadTexture(e.LoadMenu)
	rl.UnloadTexture(e.LoadPause)
}
