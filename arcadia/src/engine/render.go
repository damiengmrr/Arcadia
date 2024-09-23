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
	// chargement de l'image de fond du menu
	rl.DrawTexture(rl.LoadTexture("textures/map/tilesets/menu.png"), 0, 0, rl.White)

	// texte du menu
	rl.DrawText("ISILDOR", int32(rl.GetScreenWidth())/2-rl.MeasureText("ISILDOR", 80)/2, int32(rl.GetScreenHeight())/2-200, 80, rl.DarkGreen)
	rl.DrawText("Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Menu", 40)/2, int32(rl.GetScreenHeight())/2-100, 40, rl.Gray)
	rl.DrawText("[Enter] Pour jouer", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] Pour jouer", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.Gray)
	rl.DrawText("[Esc] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] pour quitter", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.Gray)
}

func (e *Engine) InGameRendering() { // rendu ingame
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

	// affichage de la barre de vie du joueur
	rl.DrawRectangle(posX-30, posY+700, int32(200), 15, rl.White)             // barre de vie initial
	rl.DrawRectangle(posX-30, posY+700, int32(e.Player.Health), 15, rl.Green) // idem
	if e.Player.Health < 100 {                                                // condition
		rl.DrawRectangle(posX-30, posY+700, int32(200), 15, rl.White)
		rl.DrawRectangle(posX-30, posY+700, int32(e.Player.Health), 15, rl.Yellow) // changement de couleur si la vie < 100
	}
	if e.Player.Health < 60 {
		rl.DrawRectangle(posX-30, posY+700, int32(200), 15, rl.White)
		rl.DrawRectangle(posX-30, posY+700, int32(e.Player.Health), 15, rl.Orange) // changement de couleur si la vie < 60
	}
	if e.Player.Health < 30 {
		rl.DrawRectangle(posX-30, posY+700, int32(200), 15, rl.White)
		rl.DrawRectangle(posX-30, posY+700, int32(e.Player.Health), 15, rl.Red) // changement de couleur si la vie < 30
	}
	rl.DrawText(strconv.Itoa(e.Player.Health), posX-70, posY+699, int32(20), rl.White) // affichage textuel du la vie en temps réel

	// cooldown du dash in game
	rl.DrawRectangle(posX-30, posY+680, int32(180), 5, rl.White)                                 // barre du cooldown de dash
	rl.DrawText(strconv.Itoa(e.Player.TeleportCooldown), posX-70, posY+675, int32(20), rl.White) // affichage textuel du cooldown dash en temps réel
	rl.DrawRectangle(posX-30, posY+680, int32(e.Player.TeleportCooldown), 5, rl.Red)             // fond de couleur de la barre de cooldown du dash

	rl.DrawTexture(rl.LoadTexture("textures/map/tilesets/minimap.png"), 1165, 0, rl.White) // affichage de la carte du jeux en haut a droite du jeux

	rl.DrawText(strconv.Itoa(int(rl.GetFPS())), posX-70, posY-50, int32(20), rl.White) // affichage en temps réel des FPS en haut a gauche du jeux
}

func (e *Engine) PauseRendering() { // rendu du menu pause
	// image de fond du menu pause
	rl.DrawTexture(rl.LoadTexture("textures/map/tilesets/zzzzz.png"), 0, 0, rl.White)
	// texte du menu pause
	rl.DrawText("ISILDOR", int32(rl.GetScreenWidth())/2-rl.MeasureText("ISILDOR", 80)/2, int32(rl.GetScreenHeight())/2-350, 80, rl.Orange)
	rl.DrawText("Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("Pause", 60)/2, int32(rl.GetScreenHeight())/2-150, 60, rl.DarkGray)
	rl.DrawText("[Esc] pour reprendre", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] pour reprendre", 30)/2, int32(rl.GetScreenHeight())/2, 30, rl.Black)
	rl.DrawText("[Q]/[A] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Q]/[A] pour quitter", 30)/2, int32(rl.GetScreenHeight())/2+100, 30, rl.Black)
}

func (e *Engine) GameOverRendering() { // rendu du GameOver
	// couleur de fond
	rl.ClearBackground(rl.Black)
	// image de fond
	rl.DrawTexture(rl.LoadTexture("textures/map/tilesets/gameOver.jpg"), 0, 0, rl.White)

}

func (e *Engine) RenderPlayer() { // rendu joueur

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

}

func (e *Engine) RenderMonsters() { // rendu monstre

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

func (e *Engine) RenderDialog(m entity.Monster, sentence string) { // rendu du dialog
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

func (e *Engine) RenderHealth() { // rendu de la vie des monstres
	rl.BeginMode2D(e.Camera)

	for _, monster := range e.Monsters {
		distance := rl.Vector2Distance(e.Player.Position, monster.Position) // affichage à une certaine distance
		if distance <= ChaseDistance {                                      // condition pour la distance de chasse
			if monster.IsAlive && monster.IsMove { // condition permettant d'etre sur du monstre
				// affichage de la barre de vie des monstres
				rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+30, int32(20), 5, rl.DarkBrown)
				rl.DrawText(strconv.Itoa(monster.Health), int32(monster.Position.X)+25, int32(monster.Position.Y)+35, int32(3), rl.White) // affichage du nombre de vie
				rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+30, int32(monster.Health), 5, rl.Red)
				rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+30, int32(20), 5, rl.DarkBrown)
				rl.DrawText(strconv.Itoa(monster.Health), int32(monster.Position.X)+25, int32(monster.Position.Y)+35, int32(3), rl.White) // affichage du nombre de vie
				rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+30, int32(monster.Health), 5, rl.Red)
			}
		}
	}
	rl.EndMode2D()
}
func (e *Engine) WinRendering() { // rendu de la Win du jeux
	// chargement de l'image de fond ainsi que la couleur
	rl.ClearBackground(rl.White)
	rl.DrawTexture(rl.LoadTexture("textures/map/tilesets/win.jpeg"), 380, 0, rl.White)

	// texte du menu
	rl.DrawText("WIN", int32(rl.GetScreenWidth())/2-rl.MeasureText("WIN", 80)/2, int32(rl.GetScreenHeight())/2+300, 80, rl.DarkGreen)
	rl.DrawText("ISILDOR", int32(rl.GetScreenWidth())/2-rl.MeasureText("[ISILDOR", 30)/2, int32(rl.GetScreenHeight())/2+100, 30, rl.Gold)
	rl.DrawText("[Esc] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] pour quitter", 20)/2, int32(rl.GetScreenHeight())/2+200, 20, rl.Gray)
}
