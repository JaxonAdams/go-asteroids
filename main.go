package main

import (
	player "github.com/JaxonAdams/go-asteroids/Player"
	"github.com/JaxonAdams/go-asteroids/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var p player.PlayerShip

func main() {
	rl.InitWindow(
		constants.WINDOW_SIZE_X,
		constants.WINDOW_SIZE_Y,
		"ASTEROIDS",
	)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	p.Init()

	for !rl.WindowShouldClose() {

		update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		draw()
		rl.EndDrawing()
	}
}

func update() {
	p.HandleInput()
}

func draw() {
	// Player Ship
	drawShape(
		p.Position,
		constants.SCALE,
		p.Rotation,
		p.Shape,
	)

	shouldDrawTail := int32(rl.GetTime()*20)%2 == 0
	if p.IsThrusting && shouldDrawTail {
		// Player Thrust Tail
		drawShape(
			p.Position,
			constants.SCALE,
			p.Rotation,
			p.TailShape,
		)
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
