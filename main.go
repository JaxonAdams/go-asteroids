package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "ASTEROIDS")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		// Player Ship
		rl.DrawLine(100, 300, 150, 200, rl.White)
		rl.DrawLine(150, 200, 200, 300, rl.White)
		rl.DrawLine(100, 300, 120, 280, rl.White)
		rl.DrawLine(200, 300, 180, 280, rl.White)
		rl.DrawLine(120, 280, 180, 280, rl.White)

		rl.EndDrawing()
	}
}
