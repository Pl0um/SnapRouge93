package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e Engine) Rendering() {
	rl.DrawTexture(rl.LoadTexture("textures/img3.jpg"), 0, 0, rl.White)
}

func (e Engine) HomeRendering() {
	rl.DrawTexture(rl.LoadTexture("textures/img3.jpg"), 0, 0, rl.White)

	rl.DrawText("Menu Principal", int32(rl.GetScreenWidth())/2-rl.MeasureText("Menu Principal", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.Red)
	rl.DrawText("Appuyez sur [Entrer] pour jouer", int32(rl.GetScreenWidth())/2-rl.MeasureText("Appuyez sur [Entrer] pour jouer", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.Red)
	rl.DrawText("Appuyez sur [Echap] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("Appuyez sur [Echap] pour quitter", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.Red)

}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)

}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(rl.LoadTexture("textures/img4.jpg"), 0, 0, rl.White)

	rl.DrawText("Menu Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("Menu Pause", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.Red)
	rl.DrawText("[P] ou [Echap] pour revenir au jeu", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] ou [Echap] pour revenir au jeu", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.Red)
	rl.DrawText("[Q]/[A] pour Quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Q]/[A] pour Quitter", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.Red)

}

func (e *Engine) RenderPlayer() {

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		rl.DrawTexturePro(
			monster.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}

func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}
