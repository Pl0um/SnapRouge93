package entity

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type  Marchant struct {
	Name     string
	Position rl.Vector2
	Health   int
	Damage   int
	Worth    int //valeur en argent quand tué

	IsAlive bool

	Sprite rl.Texture2D
}

func (a *Marchant) Attack(p *Player) {
	p.Health -= 1
}



func (a *Marchant) ToString() {
	fmt.Printf("je vends des choses")
}
