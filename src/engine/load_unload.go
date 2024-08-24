package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {

	e.Sprites.Plants = rl.LoadTexture("textures/map/tilesets/TX Plant.png")
	e.Sprites.PlantsShadow = rl.LoadTexture("textures/map/tilesets/TX Shadow Plant.png")
	/* e.Sprites = append(e.Sprites,
		Sprite{
			Name:   "grass",
			Sprite: rl.LoadTexture("texture/map/tilesets/TXPlant.png"),
		},
	) */

	fmt.Println("====================================")
	fmt.Println(e.Sprites.Plants.ID)
}

func (e *Engine) Unload() {
	rl.UnloadTexture(e.Sprites.Plants)
	rl.UnloadTexture(e.Sprites.PlantsShadow)
	/*
		 	for i := 0; i < len(e.Sprites); i++ {
				rl.UnloadTexture(e.Sprites[i].Sprite)
			}
	*/
}
