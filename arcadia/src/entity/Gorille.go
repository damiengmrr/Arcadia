package entity

import (
	"fmt"
	"main/src/item"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Gorille struct {
	Name           string
	Position       rl.Vector2
	Health         int
	Damage         int
	Loot           []item.Item
	Worth          int //valeur en argent quand tué
	Speed          float32
	IsAlive        bool
	Cooldown       time.Duration // Durée du cooldown (ex: 2 secondes)
	LastAttackTime time.Time     // Dernière fois où le monstre a attaqué

	Sprite rl.Texture2D
}

func (m *Gorille) Attack(p *Player) {
	p.Health -= 20
}

func (m *Gorille) ToString() {
	fmt.Printf("Je suis un monstre avec %d points de vie\n", m.Health)
}
