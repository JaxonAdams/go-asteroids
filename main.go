package main

import (
	"math"
	"math/rand/v2"
	"time"

	player "github.com/JaxonAdams/go-asteroids/Player"
	"github.com/JaxonAdams/go-asteroids/asteroid"
	"github.com/JaxonAdams/go-asteroids/constants"
	"github.com/JaxonAdams/go-asteroids/particle"
	"github.com/JaxonAdams/go-asteroids/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState struct {
	PlayerShip    player.PlayerShip
	AsteroidField []*asteroid.Asteroid
	GameParticles []particle.IParticle
}

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

	state := prepareLevel([]*asteroid.Asteroid{})

	for !rl.WindowShouldClose() {

		update(&state)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		draw(state)
		rl.EndDrawing()
	}
}

func prepareLevel(existingAsteroids []*asteroid.Asteroid) GameState {
	var state GameState

	if len(existingAsteroids) > 0 {
		state.AsteroidField = existingAsteroids
	} else {
		state.AsteroidField = loadLevel()
	}

	state.PlayerShip = player.PlayerShip{}
	state.GameParticles = []particle.IParticle{}

	state.PlayerShip.Init()

	return state
}

func update(state *GameState) {
	// Update asteroids
	for _, a := range state.AsteroidField {
		a.Move()
	}

	// Update particles
	newParticles := (state.GameParticles)[:0]
	for _, pt := range state.GameParticles {
		pt.Update()
		if !pt.IsDead() {
			newParticles = append(newParticles, pt)
		}
	}
	state.GameParticles = newParticles

	if !state.PlayerShip.IsDead() {
		state.PlayerShip.HandleInput()

		for _, a := range state.AsteroidField {
			// Check for collision
			dist := rl.Vector2Distance(a.Position, state.PlayerShip.Position)
			asteroidRadius := float32(a.GetSize() * a.GetCollisionScale() * constants.SCALE)
			playerRadius := state.PlayerShip.GetCollisionRadius()

			if dist < (asteroidRadius + playerRadius) {
				state.PlayerShip.DeathTime = constants.PLAYER_DEATH_COOLDOWN

				// Create death particles
				for range 5 {
					angle := 2 * math.Pi * rng.Float64()
					newParticle := &particle.LineParticle{
						Rotation: 2 * math.Pi * rng.Float32(),
						Length:   constants.SCALE * (0.6 + (0.4 * rng.Float32())),
						Particle: particle.Particle{
							Position: rl.Vector2Add(
								state.PlayerShip.Position,
								rl.Vector2{X: rng.Float32() * 3, Y: rng.Float32() * 3},
							),
							Velocity: rl.Vector2Scale(
								rl.Vector2{X: float32(math.Cos(angle)), Y: float32(math.Sin(angle))},
								2.0*rng.Float32(),
							),
							Ttl: 3.0 + rng.Float32(),
						},
					}
					state.GameParticles = append(state.GameParticles, newParticle)
				}
			}
		}
	} else {
		state.PlayerShip.DeathTime -= float64(rl.GetFrameTime())
		state.PlayerShip.IsThrusting = false

		if state.PlayerShip.DeathTime <= 0.0 {
			newState := prepareLevel(state.AsteroidField)
			state.PlayerShip = newState.PlayerShip
			state.AsteroidField = newState.AsteroidField
			state.GameParticles = newParticles
		}
	}
}

func draw(state GameState) {
	// Player Ship
	if !state.PlayerShip.IsDead() {
		utils.DrawShape(
			state.PlayerShip.Position,
			constants.SCALE,
			state.PlayerShip.Rotation,
			state.PlayerShip.Shape,
		)
	}

	shouldDrawTail := int32(rl.GetTime()*20)%2 == 0
	if state.PlayerShip.IsThrusting && shouldDrawTail {
		// Player Thrust Tail
		utils.DrawShape(
			state.PlayerShip.Position,
			constants.SCALE,
			state.PlayerShip.Rotation,
			state.PlayerShip.TailShape,
		)

	}

	// Asteroids
	for _, a := range state.AsteroidField {
		utils.DrawShape(
			a.Position,
			constants.SCALE,
			a.Rotation,
			a.Shape,
		)
	}

	// Particles
	for _, pt := range state.GameParticles {
		pt.Draw()
	}
}

func loadLevel() []*asteroid.Asteroid {
	numAsteroids := 10
	asteroids := make([]*asteroid.Asteroid, 0, numAsteroids)

	for range numAsteroids {
		size := rng.Int32N(int32(asteroid.NumAsteroidSizes))
		asteroids = append(asteroids, asteroid.New(asteroid.AsteroidSize(size)))
	}

	return asteroids
}
