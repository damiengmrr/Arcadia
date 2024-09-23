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
		Position:             rl.Vector2{X: 480, Y: 4200},
		Health:               200,  // pv du joueur
		MaxHealth:            200,  // pv max du joueur
		Money:                1000, // argent du joueur
		Speed:                2,    // vitesse di joueur
		TeleportCooldown:     180,  // cooldown du dash
		Inventory:            []item.Item{},
		IsAlive:              true,
		Damage:               2,   // degat du joueur
		Pquantity:            1,   // quantiter de potion
		MaxSprintTime:        180, // temps de course max
		CurrentSprintTime:    300,
		SprintCooldownTime:   300, // cooldown de la course
		CurrentCooldown:      0,
		IsSprinting:          false,
		HealthPlayerCooldown: 0, // cooldown du soin
		Sprite:               e.Player.Sprite,
	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "claude",
		Position:       rl.Vector2{X: 480, Y: 4000},
		Health:         10,
		MaxHealthM:     10,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "gerard",
		Position:       rl.Vector2{X: 500, Y: 220},
		Health:         20,
		MaxHealthM:     20,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "gerard1",
		Position:       rl.Vector2{X: 248, Y: 2480},
		Health:         20,
		MaxHealthM:     20,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "gerard2",
		Position:       rl.Vector2{X: 364, Y: 2480},
		Health:         20,
		MaxHealthM:     20,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "gerard3",
		Position:       rl.Vector2{X: 480, Y: 3308},
		Health:         20,
		MaxHealthM:     20,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "gerard4",
		Position:       rl.Vector2{X: 596, Y: 2768},
		Health:         20,
		MaxHealthM:     20,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "gerard5",
		Position:       rl.Vector2{X: 712, Y: 2486},
		Health:         20,
		MaxHealthM:     20,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "gerard6",
		Position:       rl.Vector2{X: 828, Y: 2980},
		Health:         20,
		MaxHealthM:     20,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "gerard7",
		Position:       rl.Vector2{X: 944, Y: 3088},
		Health:         20,
		MaxHealthM:     20,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "gerard8",
		Position:       rl.Vector2{X: 1060, Y: 3264},
		Health:         20,
		MaxHealthM:     20,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		CooldownPeriod: 6 * time.Second,
		Cooldown:       1 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "gerard9",
		Position:       rl.Vector2{X: 1176, Y: 2676},
		Health:         20,
		MaxHealthM:     20,
		Damage:         2,
		Speed:          1.9,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	//sypder
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "spiedy1",
		Position:       rl.Vector2{X: 1456, Y: 1114},
		Health:         10,
		MaxHealthM:     20,
		Damage:         20,
		Speed:          2.75,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/spyder/spider.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "spiedy2",
		Position:       rl.Vector2{X: 1000, Y: 2026},
		Health:         10,
		MaxHealthM:     10,
		Damage:         20,
		Speed:          2.75,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/spyder/spider.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "spiedy3",
		Position:       rl.Vector2{X: 1684, Y: 1684},
		Health:         10,
		MaxHealthM:     10,
		Damage:         20,
		Speed:          2.75,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/spyder/spider.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "spiedy4",
		Position:       rl.Vector2{X: 1912, Y: 1342},
		Health:         10,
		MaxHealthM:     10,
		Damage:         20,
		Speed:          2.75,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/spyder/spider.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "spiedy5",
		Position:       rl.Vector2{X: 1228, Y: 1456},
		Health:         10,
		MaxHealthM:     10,
		Damage:         20,
		Speed:          2.75,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/spyder/spider.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "spiedy6",
		Position:       rl.Vector2{X: 2026, Y: 1798},
		Health:         10,
		MaxHealthM:     10,
		Damage:         20,
		Speed:          2.75,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/spyder/spider.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "spiedy7",
		Position:       rl.Vector2{X: 1342, Y: 1912},
		Health:         10,
		MaxHealthM:     10,
		Damage:         20,
		Speed:          2.75,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/spyder/spider.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "spiedy8",
		Position:       rl.Vector2{X: 1798, Y: 1000},
		Health:         10,
		MaxHealthM:     10,
		Damage:         20,
		Speed:          2.75,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/spyder/spider.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "spiedy9",
		Position:       rl.Vector2{X: 1114, Y: 1570},
		Health:         10,
		MaxHealthM:     10,
		Damage:         20,
		Speed:          2.75,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/spyder/spider.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:           "spiedy10",
		Position:       rl.Vector2{X: 1570, Y: 1228},
		Health:         10,
		MaxHealthM:     10,
		Damage:         20,
		Speed:          2.75,
		Loot:           []item.Item{},
		Cooldown:       1 * time.Second,
		CooldownPeriod: 6 * time.Second,
		Worth:          12,
		IsMove:         true,
		IsAlive:        true,
		Sprite:         rl.LoadTexture("textures/entities/spyder/spider.png"),
	})

	// fermier
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Fermier",
		Position: rl.Vector2{X: 480, Y: 4200},
		Health:   5,
		Damage:   5,
		Speed:    1.9,
		Loot:     []item.Item{},
		Cooldown: 1 * time.Second,
		Worth:    12,
		IsMove:   false,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Fermier/fermier.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Boss",
		Position: rl.Vector2{X: 4800, Y: 860},
		Health:   300,
		Damage:   45,
		Speed:    1,
		Loot:     []item.Item{},
		Cooldown: 3 * time.Second,
		Worth:    12,
		IsMove:   true,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Gorille/Gorille-Boss.png"),
	})
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
