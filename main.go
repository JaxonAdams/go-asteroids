package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_SIZE_X = 800
const WINDOW_SIZE_Y = 450
const SCALE = 40

type Player struct {
	pos rl.Vector2
}

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
	rl.InitWindow(WINDOW_SIZE_X, WINDOW_SIZE_Y, "ASTEROIDS")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	player := Player{
		pos: rl.Vector2Subtract(rl.Vector2Scale(rl.Vector2{X: WINDOW_SIZE_X, Y: WINDOW_SIZE_Y}, 0.5), rl.Vector2{X: SCALE, Y: SCALE}),
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		// Player Ship
		drawShape(
			player.pos,
			SCALE,
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
