package engine

import (
	"fmt"
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = RUNNING

		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-07-Simon_s-In-There-Somewhere.mp3")
		rl.PlayMusicStream(e.Music)
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) RunningLogic() {
	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
	}

	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	// Collision avec monstres
	for _, monstre := range e.Monsters {
		if collision(monstre, e.Player) {
			fmt.Println("collision !")
		}
	}
	fmt.Println("=============")
	fmt.Println(e.Monsters["claude"].Position.X, e.Monsters["claude"].Position.Y)
	fmt.Println(e.Player.Position.X, e.Player.Position.Y)

	//Musique
	rl.UpdateMusicStream(e.Music)
}

func collision(monster entity.Monster, player entity.Player) bool {
	if monster.Position.X > player.Position.X-20 &&
		monster.Position.X < player.Position.X+20 &&
		monster.Position.Y > player.Position.Y-20 &&
		monster.Position.Y < player.Position.Y+20 {

		return true
	} else {
		return false
	}
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = RUNNING
	}
	if rl.IsKeyPressed(rl.KeyA) {
		e.StateMenu = HOME

		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
		rl.PlayMusicStream(e.Music)
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}
