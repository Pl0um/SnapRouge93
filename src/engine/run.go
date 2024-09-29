package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) Run() {
	rl.SetTargetFPS(60)

	for engine.IsRunning {

		rl.BeginDrawing()

		switch engine.StateMenu {
		case HOME:
			engine.HomeRendering()
			engine.HomeLogic()

		case SETTINGS:
			engine.SettingsLogic()

		case PLAY:
			switch engine.StateEngine {

			case INEXPLORATION:
				engine.InExplorationRendering() // en jeu mais pas en combat
				engine.InExplorationLogic()

			case INCOMBAT:
				engine.InCombatRendering() // en jeu et en combat
				engine.InCombatLogic()

			case PAUSE:
				engine.PauseRendering()
				engine.PauseLogic()

			case GAMEOVER:
				engine.GAMEOVERRendering() // le Game Over
				engine.GAMEOVERLogic()
				//...
			}
		}

		rl.EndDrawing()
	}
}
