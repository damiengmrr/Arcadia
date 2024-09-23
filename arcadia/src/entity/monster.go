package entity

import (
	"fmt"
	"main/src/item"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Monster struct {
	Name           string
	Position       rl.Vector2
	Health         int
	Pquantity      int
	Damage         int
	MaxHealthM     int
	Loot           []item.Item
	Worth          int //valeur en argent quand tué
	Speed          float32
	IsAlive        bool
	Cooldown       time.Duration // Durée du cooldown (ex: 2 secondes)
	CooldownPeriod time.Duration
	LastAttackTime time.Time // Dernière fois où le monstre a attaqué*
	IsMove         bool      // bouger ou non
	LastDeadTime   time.Time

	Sprite rl.Texture2D
}

func (m *Monster) Attack(p *Player) {
	p.Health -= m.Damage
	if m.Health <= 0 {

	}
}

func (m *Monster) ToString() {
	fmt.Printf("Je suis un monstre avec %d points de vie\n", m.Health)
}
