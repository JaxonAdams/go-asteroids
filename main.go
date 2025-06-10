package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawShape(pos rl.Vector2, scale float32, points []rl.Vector2) {
	transformer := func(point rl.Vector2) rl.Vector2 {
		return rl.Vector2Add(rl.Vector2Scale(point, scale), pos)
	}

	for i := range len(points) {
		current := transformer(points[i])
		next := transformer(points[(i+1)%len(points)])

		rl.DrawLineV(current, next, rl.White)
	}
}

func main() {
	rl.InitWindow(800, 450, "ASTEROIDS")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		// Player Ship
		drawShape(
			rl.Vector2{X: 100, Y: 100},
			40,
			[]rl.Vector2{
				{X: 0.2, Y: 0.8},
				{X: 0.5, Y: 0.2},
				{X: 0.8, Y: 0.8},
				{X: 0.6, Y: 0.7},
				{X: 0.4, Y: 0.7},
			},
		)

		rl.EndDrawing()
	}
}
