package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Blue)

	rl.EndDrawing()
}

func (e *Engine) HomeRendering() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Blue)

	rl.DrawText("Home Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[Enter] to Play", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] to Play", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Esc] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

	rl.EndDrawing()
}

func (e *Engine) RunningRendering() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera)

	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)

	e.RenderMap()
	e.RenderPlayer()

	rl.EndDrawing()
}

func (e *Engine) PauseRendering() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Red)

	rl.DrawText("Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q]/[A] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

	rl.EndDrawing()
}

func (e *Engine) RenderPlayer() {

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100), //source rectangle
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

	rl.EndMode2D()
}
