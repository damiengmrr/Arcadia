package engine

import (
	"main/src/entity"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/marche.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateMenu = SETTINGS
	}

}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) InGameLogic() {
	// Mouvement
	if e.Player.TeleportCooldown > 0 {
		e.Player.TeleportCooldown-- // Diminue le cooldown de 1 à chaque frame
	}
	if e.Player.HealthPlayerCooldown > 0 {
        e.Player.HealthPlayerCooldown--
    }
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
		rl.LoadMusicStream("sounds/music/marche.mp3")
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
		rl.LoadMusicStream("sounds/music/marche.mp3")
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
		rl.LoadMusicStream("sounds/music/marche.mp3")
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
		rl.LoadMusicStream("sounds/music/marche.mp3")
	}

	// dash du joueur
	if (rl.IsKeyPressed(rl.KeyW) || rl.IsKeyDown(rl.KeyUp)) && rl.IsKeyDown(rl.KeySpace) {
		e.TeleportPlayer1()
		time.Sleep(500)
	}
	if (rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown)) && rl.IsKeyDown(rl.KeySpace) {
		e.TeleportPlayer2()
		time.Sleep(500)
	}
	if (rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft)) && rl.IsKeyDown(rl.KeySpace) {
		e.TeleportPlayer3()
		time.Sleep(500)
	}
	if (rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)) && rl.IsKeyDown(rl.KeySpace) {
		e.TeleportPlayer4()
		time.Sleep(500)
	}

	if rl.IsKeyDown(rl.KeyF) {
        e.HealthPlayer()
    }

	if (rl.IsKeyDown(rl.KeyRightShift) || rl.IsKeyDown(rl.KeyLeftShift)) && e.Player.CurrentCooldown == 0 {
        // Vérifier si le joueur a encore du temps de sprint
        if e.Player.CurrentSprintTime > 0 {
            e.Player.Speed = 4 // Vitesse de sprint
            e.Player.IsSprinting = true
            e.Player.CurrentSprintTime-- // Réduire le temps de sprint restant
        } else {
            // Si le temps de sprint est épuisé, démarrer le cooldown
            e.Player.IsSprinting = false
            e.Player.Speed = 2.0                                   // Revenir à la vitesse normale
            e.Player.CurrentCooldown = e.Player.SprintCooldownTime // Commencer le cooldown
        }
    } else {
        // Si le joueur arrête de sprinter ou si le sprint est désactivé
        e.Player.Speed = 2.0 // Vitesse normale
        e.Player.IsSprinting = false
    }

    // Gestion du cooldown du sprint
    if e.Player.CurrentCooldown > 0 {
        e.Player.CurrentCooldown-- // Décrémenter le cooldown à chaque frame
        if e.Player.CurrentCooldown == 0 {
            // Une fois le cooldown terminé, réinitialiser le temps de sprint
            e.Player.CurrentSprintTime = e.Player.MaxSprintTime
        }
    }

    // Régénération progressive du temps de sprint si le joueur ne sprinte pas
    if !e.Player.IsSprinting && e.Player.CurrentCooldown == 0 && e.Player.CurrentSprintTime < e.Player.MaxSprintTime {
        e.Player.CurrentSprintTime++ // Régénérer progressivement le temps de sprint
    }

    // Facultatif : Régénération progressive du temps de sprint si le joueur ne sprinte pas
    if !e.Player.IsSprinting && e.Player.CurrentCooldown == 0 && e.Player.CurrentSprintTime < e.Player.MaxSprintTime {
        e.Player.CurrentSprintTime++ // Régénérer progressivement le temps de sprint
    }

	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	// TEST DEGATS
		if e.Player.Health <= 0 {
			e.Player.IsAlive = false
		}
	
	if !e.Player.IsAlive {
		e.StateEngine = GAMEOVER
	}

	e.CheckCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-07-Simon_s-In-There-Somewhere.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	// paramètre in game
	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateMenu = SETTINGS
	}

}

func (e *Engine) CheckCollisions() {

	e.MonsterCollisions()
	// Mise à jour des monstres
	e.UpdateMonsters()
}

func (e *Engine) MonsterCollisions() {
    for i, monster := range e.Monsters {
        if monster.IsAlive {
            if monster.Position.X > e.Player.Position.X-20 &&
                monster.Position.X < e.Player.Position.X+20 &&
                monster.Position.Y > e.Player.Position.Y-20 &&
                monster.Position.Y < e.Player.Position.Y+20 {
                if monster.Health > 0 && e.Player.Health > 0 {
                    e.NormalTalk(monster, "Bonjour")
                    if time.Since(monster.LastAttackTime) > monster.Cooldown {
                        // Le monstre attaque
                        e.Monsters[i].Attack(&e.Player)
                        // Met à jour le dernier moment de l'attaque
                        e.Monsters[i].LastAttackTime = time.Now()
                        //e.Monsters[i].Attack(&e.Player)
                    }
                    if rl.IsKeyPressed(rl.KeyE) {
                        //fight.Fight(e.Player, monster)
                        e.Player.Attack(&e.Monsters[i])
                    }
                }
            }
			
        }
    }
}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyA) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}
func (e *Engine) UpdateMonsters() {
	for i, monster := range e.Monsters {
		if monster.IsAlive {
			distance := rl.Vector2Distance(e.Player.Position, monster.Position)

			if distance <= ChaseDistance {
				direction := rl.Vector2Subtract(e.Player.Position, monster.Position)
				direction = rl.Vector2Normalize(direction)
				monster.Position = rl.Vector2Add(monster.Position, rl.Vector2Scale(direction, monster.Speed))
				e.Monsters[i] = monster // Met à jour le monstre dans la liste
			}
		}
	}
}

// TeleportePlayer téléporte le joueur à une nouvelle position
func (e *Engine) TeleportPlayer1() {
	if e.Player.TeleportCooldown <= 0 {
		newX := e.Player.Position.X
		newY := e.Player.Position.Y - 100 // Z
		time.Sleep(500)

		e.Player.Position = rl.Vector2{X: newX, Y: newY}
		e.Player.TeleportCooldown = 180
	}
}
func (e *Engine) TeleportPlayer2() {
	if e.Player.TeleportCooldown <= 0 {
		newX := e.Player.Position.X
		newY := e.Player.Position.Y + 100 // S
		e.Player.Position = rl.Vector2{X: newX, Y: newY}
		e.Player.TeleportCooldown = 180
	}
}
func (e *Engine) TeleportPlayer3() {
	if e.Player.TeleportCooldown <= 0 {
		newX := e.Player.Position.X - 100 // Q
		newY := e.Player.Position.Y

		e.Player.Position = rl.Vector2{X: newX, Y: newY}
		e.Player.TeleportCooldown = 180
	}
}

func (e *Engine) TeleportPlayer4() {
	if e.Player.TeleportCooldown <= 0 {
		newX := e.Player.Position.X + 100 //D
		newY := e.Player.Position.Y

		e.Player.Position = rl.Vector2{X: newX, Y: newY}
		e.Player.TeleportCooldown = 180
	}

}

func (e *Engine) GameOverLogic() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.reset()
	}
}

func (e *Engine) reset() {
	e.Player.Health = 200
	e.Player.IsAlive = true
	e.StateMenu = PLAY
	e.StateEngine = INGAME
	e.Player.Position = rl.Vector2{X: 370, Y: 280}
		
}

func (e *Engine) HealthPlayer() {
    if e.Player.HealthPlayerCooldown <= 0 && e.Player.Health <= 150 {
        e.Player.Health += 50

        e.Player.HealthPlayerCooldown = 300
    }

}
