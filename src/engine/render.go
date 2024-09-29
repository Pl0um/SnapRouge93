package engine

import (
	"strconv"
	"main/src/entity"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.DrawTexture(e.LoadMenu, 0, 0, rl.White)
}

func (e *Engine) HomeRendering() {//affichage du menu Home
	rl.DrawTexture(e.LoadMenu, 0, 0, rl.White)

	rl.DrawText("Menu Principal", int32(rl.GetScreenWidth())/2-rl.MeasureText("Menu Principal", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.Red)
	rl.DrawText("[Entrer] pour jouer", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Entrer] pour jouer", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.Red)
	rl.DrawText("[Esc] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] pour quitter", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.Red)

}
func (e *Engine) GAMEOVERRendering() { // ecran de game over
	rl.ClearBackground(rl.Black)

	rl.DrawText("GAME OVER", int32(rl.GetScreenWidth())/2-rl.MeasureText("GAME OVER", 50)/2, int32(rl.GetScreenHeight())/2-200, 50, rl.Red)
	rl.DrawText("tu es mort", int32(rl.GetScreenWidth())/2-rl.MeasureText("tu es mort", 25)/2, int32(rl.GetScreenHeight())/2-100, 25, rl.Red)
	rl.DrawText("[Enter] pour rejouer", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] pour rejouer", 25)/2, int32(rl.GetScreenHeight())/2-40, 25, rl.Gray)
	rl.DrawText("[Esc] pour revenir au menu principal", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] pour revenir au menu principal", 25)/2, int32(rl.GetScreenHeight())/2-65, 25, rl.Gray)
	rl.DrawText("[L] pour RageQuit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[L] pour RageQuit", 25)/2, int32(rl.GetScreenHeight())/2-15, 25, rl.Gray)

	rl.EndDrawing()

}

func (e *Engine) InExplorationRendering() {//en exploration
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()
	e.RenderAnimal()
	e.RenderMonsters()
	e.RenderPlayer()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText("En Exploration", int32(rl.GetScreenWidth())/2-rl.MeasureText("En Exploration", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)

	rl.DrawText("Argent :", int32(rl.GetScreenWidth())/19-rl.MeasureText("Argent :", 40)/3, int32(rl.GetScreenHeight())/2-325, 35, rl.Yellow)
	rl.DrawText(strconv.Itoa(e.Player.Money), int32(rl.GetScreenWidth())/6-rl.MeasureText("Home Menu", 40)/4, int32(rl.GetScreenHeight())/2-325, 40, rl.Yellow)

	rl.DrawText("Pv :", int32(rl.GetScreenWidth())/23-rl.MeasureText("Pv :", 40)/2, int32(rl.GetScreenHeight())/2-375, 35, rl.Red)
	rl.DrawText(strconv.Itoa(e.Player.Health), int32(rl.GetScreenWidth())/7-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-375, 40, rl.Red)

	margeTop := int32(10)

	posX := int32(rl.GetScreenWidth())/2 - 600
    posY := margeTop


    rl.DrawRectangle(posX-250, posY +1000, int32(e.Player.Health), 15, rl.Green)
    if e.Player.Health < 100 {
        rl.DrawRectangle(posX-250, posY +1000, int32(200), 15, rl.White)
         rl.DrawRectangle(posX-250, posY +1000, int32(e.Player.Health), 15, rl.Yellow)
    }
    if e.Player.Health < 60 {
        rl.DrawRectangle(posX-250, posY +1000, int32(200), 15, rl.White)
         rl.DrawRectangle(posX-250, posY +1000, int32(e.Player.Health), 15, rl.Orange)
    }
    if e.Player.Health < 30 {
        rl.DrawRectangle(posX-250, posY +1000, int32(200), 15, rl.White)
         rl.DrawRectangle(posX-250, posY +1000, int32(e.Player.Health), 15, rl.Red)
    }
	
	text := fmt.Sprintf("potion restante : %d", len(e.Player.Inventory))
	rl.DrawText(text, 10, 70, 30, rl.RayWhite)


	rl.DrawText("FPS: "+strconv.Itoa(int(rl.GetFPS())), int32(rl.GetScreenWidth())/2+550, int32(rl.GetScreenHeight())/2-550, 40, rl.RayWhite)


}

func (e *Engine) InCombatRendering() { // ecran de combat
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera)
	e.RenderMap()
	e.RenderPlayer()
	e.RenderMonsters()

	rl.EndMode2D()

	rl.DrawText("En combat", int32(rl.GetScreenWidth())/2-rl.MeasureText("En combat", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.Black)
	rl.DrawText("appuyez sur Esc pour fuir", int32(rl.GetScreenWidth())/2-rl.MeasureText("appuyez sur Esc pour fuir", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.Black)

	rl.DrawText("Pv :", int32(rl.GetScreenWidth())/23-rl.MeasureText("Pv :", 40)/2, int32(rl.GetScreenHeight())/2-375, 35, rl.Red)
	rl.DrawText(strconv.Itoa(e.Player.Health), int32(rl.GetScreenWidth())/7-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-375, 40, rl.Red)

	margeTop := int32(10)

	posX := int32(rl.GetScreenWidth())/2 - 600
    posY := margeTop


    rl.DrawRectangle(posX-250, posY +1000, int32(e.Player.Health), 15, rl.Green)
    if e.Player.Health < 100 {
        rl.DrawRectangle(posX-250, posY +1000, int32(200), 15, rl.White)
         rl.DrawRectangle(posX-250, posY +1000, int32(e.Player.Health), 15, rl.Yellow)
    }
    if e.Player.Health < 60 {
        rl.DrawRectangle(posX-250, posY +1000, int32(200), 15, rl.White)
         rl.DrawRectangle(posX-250, posY +1000, int32(e.Player.Health), 15, rl.Orange)
    }
    if e.Player.Health < 30 {
        rl.DrawRectangle(posX-250, posY +1000, int32(200), 15, rl.White)
         rl.DrawRectangle(posX-250, posY +1000, int32(e.Player.Health), 15, rl.Red)
    }
	rl.DrawText("FPS: "+strconv.Itoa(int(rl.GetFPS())), int32(rl.GetScreenWidth())/2+550, int32(rl.GetScreenHeight())/2-550, 40, rl.RayWhite)

}

func (e *Engine) PauseRendering() {//affichage du menus pause
	rl.DrawTexture(e.LoadPause, 0, 0, rl.White)

	rl.DrawText("Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.Red)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.Red)
	rl.DrawText("[Q]/[A] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.Red)

	rl.EndDrawing()
}

func (e *Engine) RenderPlayer() {//afficher le joueur+
	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 70, 70),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 65, 65),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

}

func (e *Engine) RenderAnimal() {//afficher l'animal marchant
	for _, marchant := range e.Marchant {
		rl.DrawTexturePro(
			marchant.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(marchant.Position.X, marchant.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}

func (e *Engine) RenderMonsters() {//afficher les monstres
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
func (e *Engine) RenderDialog2(m entity.Marchant, sentence string) {//dialogue animal marchant
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X)+50,
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}
func (e *Engine) RenderDialog3(m entity.Marchant, sentence string) {//dialogue animal marchant
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X)+50,
		int32(m.Position.Y)+30,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}
func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X)+50,
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}
