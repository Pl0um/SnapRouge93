package fight

import (
	"main/src/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type fight int

const (
	PLAYER_TURN  fight = iota
	MONSTER_TURN fight = iota
)

func Fight(player entity.Player, monster entity.Monster) {
	for { 
		if rl.IsKeyPressed(rl.KeyE) {
			player.Attack(&monster)
			if monster.Health <= 0 {
				monster.IsAlive = false
				break

			}
		}


	}
}
