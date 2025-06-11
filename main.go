package main

import (
	player "github.com/JaxonAdams/go-asteroids/Player"
	"github.com/JaxonAdams/go-asteroids/asteroid"
	"github.com/JaxonAdams/go-asteroids/constants"
	"github.com/JaxonAdams/go-asteroids/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var p player.PlayerShip
var a *asteroid.Asteroid // TODO: remove me

func main() {
	rl.InitWindow(
		constants.WINDOW_SIZE_X,
		constants.WINDOW_SIZE_Y,
		"ASTEROIDS",
	)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	p.Init()

	a = asteroid.New()

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
	utils.DrawShape(
		p.Position,
		constants.SCALE,
		p.Rotation,
		p.Shape,
	)

	shouldDrawTail := int32(rl.GetTime()*20)%2 == 0
	if p.IsThrusting && shouldDrawTail {
		// Player Thrust Tail
		utils.DrawShape(
			p.Position,
			constants.SCALE,
			p.Rotation,
			p.TailShape,
		)

	}

	// Asteroids
	utils.DrawShape(
		a.Position,
		constants.SCALE,
		a.Rotation,
		a.Shape,
	)
}
