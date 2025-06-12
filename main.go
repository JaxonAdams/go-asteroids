package main

import (
	player "github.com/JaxonAdams/go-asteroids/Player"
	"github.com/JaxonAdams/go-asteroids/asteroid"
	"github.com/JaxonAdams/go-asteroids/constants"
	"github.com/JaxonAdams/go-asteroids/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand/v2"
	"time"
)

var p player.PlayerShip
var s = rand.NewPCG(42, uint64(time.Now().Unix()))
var rng = rand.New(s)

func main() {
	rl.InitWindow(
		constants.WINDOW_SIZE_X,
		constants.WINDOW_SIZE_Y,
		"ASTEROIDS",
	)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	p.Init()

	asteroids := loadLevel()

	for !rl.WindowShouldClose() {

		update(asteroids)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		draw(asteroids)
		rl.EndDrawing()
	}
}

func update(asteroids []asteroid.Asteroid) {
	p.HandleInput()
	for i := range asteroids {
		asteroids[i].Move()
	}
}

func draw(asteroids []asteroid.Asteroid) {
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
	for _, a := range asteroids {
		utils.DrawShape(
			a.Position,
			constants.SCALE,
			a.Rotation,
			a.Shape,
		)
	}
}

func loadLevel() []asteroid.Asteroid {
	numAsteroids := 10
	asteroids := make([]asteroid.Asteroid, 0, numAsteroids)

	for range numAsteroids {
		size := rng.Int32N(int32(asteroid.NumAsteroidSizes))
		asteroids = append(asteroids, *asteroid.New(asteroid.AsteroidSize(size)))
	}

	return asteroids
}
