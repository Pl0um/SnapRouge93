package engine

import (
	"fmt"
	"main/src/entity"
	"main/src/fight"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/Bloody Tears (Castlevania II Simons Quest OST) (Original Mix).mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INEXPLORATION
		rl.StopMusicStream(e.Music)
		e.Player.Inventory = append(e.Player.Inventory, e.Potion)
		e.Player.Inventory = append(e.Player.Inventory, e.Potion)
		e.Player.Inventory = append(e.Player.Inventory, e.Potion)

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
		e.InitItems()
		e.InitEntities()
		e.InitCamera()
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}


func (e *Engine) InCombatLogic() {
	//positionnement du joueur lors d'un combat
	for _, monster := range e.Monsters {
		e.Player.Position.X = monster.Position.X - 70
		e.Player.Position.Y = monster.Position.Y + 33
	}
	// pour fuir le combat
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INEXPLORATION
		fmt.Printf("je prends la fuite")
	}
	if e.Player.IsAlive == false {
		e.StateEngine = GAMEOVER
	}
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}
}


func (e *Engine) GAMEOVERLogic() {
	if !rl.IsMusicStreamPlaying(e.Music) {
        e.Music = rl.LoadMusicStream("sounds/music/Game Over (8-Bit Music).mp3")
        rl.PlayMusicStream(e.Music)
    }
    rl.UpdateMusicStream(e.Music)

	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateMenu = HOME
		e.InitItems()
		e.InitEntities()
		e.InitCamera()
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.InitItems()
		e.InitEntities()
		e.InitCamera()
		e.StateEngine = INEXPLORATION
	}
	if rl.IsKeyPressed(rl.KeyL) {
		e.IsRunning = false
	}
}
func (e *Engine) RemoveItemFromInventory(index int) {
	if index >= 0 && index < len(e.Player.Inventory) {
		e.Player.Inventory = append(e.Player.Inventory[:index], e.Player.Inventory[index+1:]...)
	}
}

func (e *Engine) InExplorationLogic() {

	if rl.IsKeyPressed(rl.KeyR) {
		if e.Player.Health+20 < 100 {
			e.Player.Health += 20
			e.RemoveItemFromInventory(0)
		} else if e.Player.Health+20 >= 100 {
			e.Player.Health = 100
			e.RemoveItemFromInventory(0)
		}
	}
	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
        e.Player.Speed = 5
    } else {
        e.Player.Speed = 2
    }


	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	if rl.IsKeyPressed(rl.KeyJ) {
		e.StateEngine = GAMEOVER
	}

	e.CheckCollisions()

	if !rl.IsMusicStreamPlaying(e.Music) {
        e.Music = rl.LoadMusicStream("sounds/music/Undertale OST 034 - Memory.mp3")
        rl.PlayMusicStream(e.Music)
    }
    rl.UpdateMusicStream(e.Music)
}

func (e *Engine) CheckCollisions() {
	e.AnimalCollisions()
	e.MonsterCollisions()
}

func (e *Engine) MonsterCollisions() {

	for _, monster := range e.Monsters {
		if monster.Position.X > e.Player.Position.X-50 &&
			monster.Position.X < e.Player.Position.X+80 &&
			monster.Position.Y > e.Player.Position.Y-60 &&
			monster.Position.Y < e.Player.Position.Y+30 {

			if monster.Name == "claude" {
				e.NormalTalk(monster, "Utilisez [Z], [Q], [S] et [D] pour vous déplacer, utilisez [Shift] pour sprinter \n  Appuyez sur [E] pour lancer un combat contre moi")
				if rl.IsKeyPressed(rl.KeyE) {
					e.StateEngine = INCOMBAT
					fight.Fight(e.Player, monster)
				}
			}
		}
	}
}

func (e *Engine) AnimalCollisions() {

	for _, Marchant := range e.Marchant {
		if Marchant.Position.X > e.Player.Position.X-50 &&
			Marchant.Position.X < e.Player.Position.X+80 &&
			Marchant.Position.Y > e.Player.Position.Y-60 &&
			Marchant.Position.Y < e.Player.Position.Y+30 {

			if Marchant.Name == "renard" {
				e.NormalTalk2(Marchant, "je vends 1 potion de soins contre 4 money. ")
				if rl.IsKeyPressed(rl.KeyE) {
					if e.Player.Money >= e.Potion.Price {
						e.Player.Money = e.Player.Money - e.Potion.Price
						e.Player.Inventory = append(e.Player.Inventory, e.Potion)
					} else {
						e.NormalTalk3(Marchant, "tu n'as pas assez d'argent")

					}
					//lancer un combat ?
				}
			}
		} else {
			//...
		}
	}
}

func (e *Engine) NormalTalk2(a entity.Marchant, sentence string) {
	e.RenderDialog2(a, sentence)
}
func (e *Engine) NormalTalk3(a entity.Marchant, sentence string) {
	e.RenderDialog3(a, sentence)
}
func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}

func (e *Engine) CipherTalk(m entity.Monster, sentence string) { //conversation codée
	var string1 string
	for _, char := range sentence {
		char = char + 2
		string1 = string1 + string(char)
	}
	e.RenderDialog(m, string1)
}
func(e *Engine) RobotTalk(m entity.Monster, sentence string) { //conversation en binaire
	var string1 string
	var b string
	for _, char := range sentence {
		x := int(char)
		b = ""
		for x / 2 != 0 {
			c := x % 2
			b = strconv.Itoa(c) + b
			x = x/2
		}
		b = "1" + b 
		for len(b) < 8 {
			b = "0" + b
		}
		string1 = string1 + b 
	}
	e.RenderDialog(m, b)
}
func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INEXPLORATION
	}
	if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyQ) {
		e.StateMenu = HOME
		e.InitItems()
		e.InitEntities()
		e.InitCamera()
	}

	if !rl.IsMusicStreamPlaying(e.Music) {
        e.Music = rl.LoadMusicStream("sounds/music/Undertale OST 034 - Memory.mp3")
        rl.PlayMusicStream(e.Music)
    }
    rl.UpdateMusicStream(e.Music)
}

