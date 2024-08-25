package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	// Chargement des textures du personnage
	e.Player.Sprite = rl.LoadTexture("textures/entities/soldier/Soldier-Idle.png")

	// Chargement des textures de la carte
	e.Sprites["Plant"] = rl.LoadTexture("textures/map/tilesets/TX Plant.png")
	e.Sprites["PlantShadow"] = rl.LoadTexture("textures/map/tilesets/TX Shadow Plant.png")
}

func (e *Engine) Unload() {
	rl.UnloadTexture(e.Player.Sprite)

	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}

}
