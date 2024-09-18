package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position         rl.Vector2
	Health           int
	Money            int
	Speed            float32
	TeleportCooldown int
	Inventory        []item.Item
	IsAlive bool
	MaxSprintTime         int  // Durée maximale de sprint
    CurrentSprintTime     int  // Temps restant de sprint
    SprintCooldownTime    int  // Temps de cooldown après un sprint
    CurrentCooldown       int  // Compteur pour le cooldown
    IsSprinting           bool // Indique si le joueur est en train de sprinter
	HealthPlayerCooldown  int  // cooldown du health le joueur
	Sprite rl.Texture2D
}

func (p Player) Attack(m*Monster) {
    m.Health -= 1
    if m.Health <= 0 {
        m.IsAlive = false
    }
}

func (p *Player) ToString() {
	fmt.Printf(`
	Joueur:
		Vie: %d,
		Argent: %d,
		Inventaire: %+v
	
	\n`, p.Health, p.Money, p.Inventory)
}
