package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_SIZE_X = 800
const WINDOW_SIZE_Y = 450
const SCALE = 40

type Player struct {
	pos   rl.Vector2
	shape []rl.Vector2
}

func computeCentroid(points []rl.Vector2) rl.Vector2 {
	var sum rl.Vector2
	for _, p := range points {
		sum = rl.Vector2Add(sum, p)
	}
	return rl.Vector2Scale(sum, 1/float32(len(points)))
}

func drawShape(pos rl.Vector2, scale float32, rotation float32, points []rl.Vector2) {
	centroid := computeCentroid(points)

	transformer := func(point rl.Vector2) rl.Vector2 {
		centered := rl.Vector2Subtract(point, centroid)
		scaled := rl.Vector2Scale(centered, scale)
		rotated := rl.Vector2Rotate(scaled, rotation)
		return rl.Vector2Add(rotated, pos)
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
		shape: []rl.Vector2{
			{X: 0.2, Y: 0.8},
			{X: 0.5, Y: 0.2},
			{X: 0.8, Y: 0.8},
			{X: 0.6, Y: 0.7},
			{X: 0.4, Y: 0.7},
		},
	}

	player.pos = rl.Vector2Add(
		rl.Vector2Scale(rl.Vector2{
			X: WINDOW_SIZE_X,
			Y: WINDOW_SIZE_Y,
		}, 0.5),
		computeCentroid(player.shape),
	)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		// Player Ship
		drawShape(
			player.pos,
			SCALE,
			float32(rl.GetTime())*rl.Pi,
			player.shape,
		)

		rl.EndDrawing()
	}
}
