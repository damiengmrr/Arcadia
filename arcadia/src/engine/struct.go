package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
)

type engine int

const (
	INGAME   engine = iota
	PAUSE    engine = iota
	GAMEOVER engine = iota
	WIN      engine = iota
)

type Engine struct {
	Player   entity.Player
	Monsters []entity.Monster
	//item     []item.Item

	Music       rl.Music
	MusicVolume float32

	Sprites map[string]rl.Texture2D

	Camera rl.Camera2D

	MapJSON MapJSON

	IsRunning   bool
	StateMenu   menu
	StateEngine engine
}

const (
	ChaseDistance = 113 // Distance à partir de laquelle le monstre commence à poursuivre le joueur
	RangeMonster  = 40  // distance attaque du monstre
	RangePlayer   = 45  // distance attaque du joueur
	TradeDistance = 70  // distance pour trade
)
