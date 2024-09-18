package engine

import (
	"main/src/entity"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.ClearBackground(rl.White)
}

func (e *Engine) HomeRendering() {
	rl.DrawTexture(rl.LoadTexture("textures/map/tilesets/menu.png"), 0, 0, rl.White)

	rl.DrawText("ARCADIA", int32(rl.GetScreenWidth())/2-rl.MeasureText("ARCADIA", 80)/2, int32(rl.GetScreenHeight())/2-200, 80, rl.DarkGreen)
	rl.DrawText("Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Menu", 40)/2, int32(rl.GetScreenHeight())/2-100, 40, rl.Gray)
	rl.DrawText("[Enter] Pour jouer", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] Pour jouer", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.Gray)
	rl.DrawText("[Esc] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] pour quitter", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.Gray)
	rl.DrawText("[TAB] paramètres", int32(rl.GetScreenWidth())/2-rl.MeasureText("[TAB] paramètres", 20)/2, int32(rl.GetScreenHeight())/2+50, 20, rl.Gray)
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.SkyBlue)

	rl.BeginMode2D(e.Camera)

	e.RenderMap()
	e.RenderMonsters()
	e.RenderPlayer()
	e.RenderHealth()
	rl.EndMode2D()

	// Taille et positionnement des textes
	fontSize := int32(40)
	marginTop := int32(10)
	// Position de l'interface à gauche de l'écran
	posX := int32(rl.GetScreenWidth())/2 - 600
	posY := marginTop

	// Affichage de l'argent
	rl.DrawText("Argent :", posX, posY, fontSize, rl.White)
	rl.DrawText(strconv.Itoa(e.Player.Money), posX+200, posY, fontSize, rl.White)

	// Espacement entre les éléments
	posY += fontSize + 10

	rl.DrawRectangle(posX-30, posY+700, int32(200), 15, rl.White)
	rl.DrawRectangle(posX-30, posY+700, int32(e.Player.Health), 15, rl.Green)
	if e.Player.Health < 100 {
		rl.DrawRectangle(posX-30, posY+700, int32(200), 15, rl.White)
		rl.DrawRectangle(posX-30, posY+700, int32(e.Player.Health), 15, rl.Yellow)
	}
	if e.Player.Health < 60 {
		rl.DrawRectangle(posX-30, posY+700, int32(200), 15, rl.White)
		rl.DrawRectangle(posX-30, posY+700, int32(e.Player.Health), 15, rl.Orange)
	}
	if e.Player.Health < 30 {
		rl.DrawRectangle(posX-30, posY+700, int32(200), 15, rl.White)
		rl.DrawRectangle(posX-30, posY+700, int32(e.Player.Health), 15, rl.Red)
	}
	rl.DrawText(strconv.Itoa(e.Player.Health), posX-70, posY+699, int32(20), rl.White)

	// cooldown du dash in game
	rl.DrawRectangle(posX-30, posY+680, int32(180), 5, rl.White)
	rl.DrawText(strconv.Itoa(e.Player.TeleportCooldown), posX-70, posY+675, int32(20), rl.White)
	rl.DrawRectangle(posX-30, posY+680, int32(e.Player.TeleportCooldown), 5, rl.Red)
}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(rl.LoadTexture("textures/map/tilesets/fond_pause.png"), -20, 0, rl.White)
	// texte du menu pause
	rl.DrawText("Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("Pause", 45)/2, int32(rl.GetScreenHeight())/2-150, 60, rl.DarkGray)
	rl.DrawText("[Esc] pour reprendre", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] pour reprendre", 25)/2, int32(rl.GetScreenHeight())/2, 30, rl.Black)
	rl.DrawText("[Q]/[A] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Q]/[A] pour quitter", 25)/2, int32(rl.GetScreenHeight())/2+100, 30, rl.Black)
}

func (e *Engine) GameOverRendering() {
	rl.DrawTexture(rl.LoadTexture("textures/map/tilesets/fond_pause.png"), -20, 0, rl.White)
	// texte du menu pause
	rl.DrawText("Game Over", int32(rl.GetScreenWidth())/2-rl.MeasureText("Pause", 45)/2, int32(rl.GetScreenHeight())/2-150, 60, rl.DarkGray)
	rl.DrawText("[Esc] pour reprendre", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] pour reprendre", 25)/2, int32(rl.GetScreenHeight())/2, 30, rl.Black)
	rl.DrawText("[Q]/[A] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Q]/[A] pour quitter", 25)/2, int32(rl.GetScreenHeight())/2+100, 30, rl.Black)

}

func (e *Engine) SettingsRendering() {
	rl.ClearBackground(rl.White)

	rl.DrawText("Paramètres", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paramètres", 60)/2, int32(rl.GetScreenHeight())/2-150, 60, rl.DarkGray)
}

func (e *Engine) RenderPlayer() {

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

}

func (e *Engine) RenderMonsters() {

	for _, monster := range e.Monsters {
		if monster.IsAlive {
			rl.DrawTexturePro(
				monster.Sprite,
				rl.NewRectangle(0, 0, 100, 100),
				rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White,
			)
		}
	}

}

func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}

func (e *Engine) RenderHealth() {
	rl.BeginMode2D(e.Camera)

	for _, monster := range e.Monsters {
		distance := rl.Vector2Distance(e.Player.Position, monster.Position)
		if distance <= ChaseDistance {
			if monster.IsAlive {
				rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+30, int32(20), 5, rl.DarkBrown)
				rl.DrawText(strconv.Itoa(monster.Health), int32(monster.Position.X)+25, int32(monster.Position.Y)+35, int32(3), rl.White)
				rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+30, int32(monster.Health), 5, rl.Red)
				rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+30, int32(20), 5, rl.DarkBrown)
				rl.DrawText(strconv.Itoa(monster.Health), int32(monster.Position.X)+25, int32(monster.Position.Y)+35, int32(3), rl.White)
				rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+30, int32(monster.Health), 5, rl.Red)
			}
		}
	}
	rl.EndMode2D()
}
