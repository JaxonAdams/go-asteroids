package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const WINDOW_SIZE_X = 800
const WINDOW_SIZE_Y = 450
const SCALE = 40
const ROTATION_SPEED = 720.0

type Player struct {
	pos      rl.Vector2
	shape    []rl.Vector2
	rotation float32
}

var player Player

func main() {
	rl.InitWindow(WINDOW_SIZE_X, WINDOW_SIZE_Y, "ASTEROIDS")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	player.shape = []rl.Vector2{
		{X: -0.2, Y: 0.3},
		{X: 0.0, Y: -0.3},
		{X: 0.2, Y: 0.3},
		{X: 0.1, Y: 0.2},
		{X: -0.1, Y: 0.2},
	}

	player.pos = rl.Vector2Add(
		rl.Vector2Scale(rl.Vector2{
			X: WINDOW_SIZE_X,
			Y: WINDOW_SIZE_Y,
		}, 0.5),
		computeCentroid(player.shape),
	)

	for !rl.WindowShouldClose() {

		update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		draw()
		rl.EndDrawing()
	}
}

func update() {
	handleInput()
}

func draw() {
	// Player Ship
	drawShape(
		player.pos,
		SCALE,
		player.rotation,
		player.shape,
	)
}

func handleInput() {
	dt := rl.GetFrameTime()

	// Player Movement
	if rl.IsKeyDown(rl.KeyA) {
		player.rotation -= degreesToRadians(ROTATION_SPEED * dt)
	}

	if rl.IsKeyDown(rl.KeyD) {
		player.rotation += degreesToRadians(ROTATION_SPEED * dt)
	}
}

func drawShape(pos rl.Vector2, scale float32, rotation float32, points []rl.Vector2) {
	transformer := func(point rl.Vector2) rl.Vector2 {
		scaled := rl.Vector2Scale(point, scale)
		rotated := rl.Vector2Rotate(scaled, rotation)
		return rl.Vector2Add(rotated, pos)
	}

	for i := range len(points) {
		current := transformer(points[i])
		next := transformer(points[(i+1)%len(points)])

		rl.DrawLineV(current, next, rl.White)
	}
}

func computeCentroid(points []rl.Vector2) rl.Vector2 {
	var sum rl.Vector2
	for _, p := range points {
		sum = rl.Vector2Add(sum, p)
	}
	return rl.Vector2Scale(sum, 1/float32(len(points)))
}

func degreesToRadians(degrees float32) float32 {
	return degrees * (math.Pi / 180.0)
}
