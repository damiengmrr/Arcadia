package engine

import (
	"main/src/entity"
	"main/src/item"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 800
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMap("textures/map/tilesets/map.json")

}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:         rl.Vector2{X: 480, Y: 4200},
		Health:           200,
		Money:            1000,
		Speed:            2,
		TeleportCooldown: 180,
		Inventory:        []item.Item{},
		IsAlive:          true,
		MaxSprintTime:         180, // 5 secondes de sprint à 60 FPS
        CurrentSprintTime:     300, // Démarrer avec le plein de temps de sprint
        SprintCooldownTime:    300, // 3 secondes de cooldown
        CurrentCooldown:       0,   // Pas de cooldown au départ
        IsSprinting:           false,
		HealthPlayerCooldown:  0,
		Sprite: e.Player.Sprite,

	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 480, Y: 4000},
		Health:   20,
		Damage:   5,
		Speed:    1.9,
		Loot:     []item.Item{},
		Cooldown: 1 * time.Second,
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "gerard",
		Position: rl.Vector2{X: 500, Y: 220},
		Health:   20,
		//MaxHealth: 20,
		Damage:   5,
		Speed:    1.9,
		Loot:     []item.Item{},
		Cooldown: 1 * time.Second,
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	//sypder
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "spiedy",
		Position: rl.Vector2{X: 600, Y: 220},
		Health:   50,
		//MaxHealth: 20,
		Damage:   5,
		Speed:    1.9,
		Loot:     []item.Item{},
		Cooldown: 1 * time.Second,
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/spyder/spyder-Idle.png"),
	})


		// Gorille
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Goren",
		Position: rl.Vector2{X: 700, Y: 320},
		Health:   50,
		//MaxHealth: 20,
		Damage:   5,
		Speed:    1.9,
		Loot:     []item.Item{},
		Cooldown: 1 * time.Second,
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Gorille/Gorille-Idle.png"),
	})

	e.Player.Money = 12
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()
	e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")

	rl.PlayMusicStream(e.Music)
}
