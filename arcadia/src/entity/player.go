package entity

import (
	"fmt"
	"main/src/item"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position             rl.Vector2
	Health               int
	MaxHealth            int
	Money                int
	Speed                float32
	TeleportCooldown     int
	Inventory            []item.Item
	IsAlive              bool
	Damage               int
	MaxSprintTime        int  // Durée maximale de sprint
	CurrentSprintTime    int  // Temps restant de sprint
	SprintCooldownTime   int  // Temps de cooldown après un sprint
	CurrentCooldown      int  // Compteur pour le cooldown
	IsSprinting          bool // Indique si le joueur est en train de sprinter
	HealthPlayerCooldown int  // cooldown du health le joueur
	Pquantity            int  // Ajout du champ pour garder la trace du moment où le monstre est mort
	CooldownPeriod       time.Duration
	Sprite               rl.Texture2D
}

func (p *Player) Attack(m *Monster) {
	currentTime := time.Now()

	// Réduction de la santé du monstre
	m.Health -= 1

	// Si la santé tombe à zéro ou en dessous, le monstre meurt
	if m.Health <= 0 {
		m.IsAlive = false
		m.LastDeadTime = currentTime
	}

	// Vérifier le temps écoulé depuis la mort et appliquer le cooldown
	if !m.IsAlive && currentTime.Sub(m.LastDeadTime) >= m.CooldownPeriod {
		//fmt.Println("Le cooldown est terminé, le monstre ressuscite.")
		//m.Health = m.MaxHealthM // Réinitialiser la santé du monstre
		//m.IsAlive = true
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
func (p *Player) ReceiveMoney(amount int) {
	p.Money += amount
}
