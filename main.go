package main

import (
	"math"
	"strconv"

	player "github.com/JaxonAdams/go-asteroids/Player"
	"github.com/JaxonAdams/go-asteroids/asteroid"
	"github.com/JaxonAdams/go-asteroids/audio"
	"github.com/JaxonAdams/go-asteroids/constants"
	"github.com/JaxonAdams/go-asteroids/particle"
	"github.com/JaxonAdams/go-asteroids/projectile"
	"github.com/JaxonAdams/go-asteroids/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var audioPlayer audio.AudioPlayer = audio.AudioPlayer{}

type GameState struct {
	PlayerShip     player.PlayerShip
	PlayerNumLives int
	PlayerScore    int
	AsteroidField  []*asteroid.Asteroid
	GameParticles  []particle.IParticle
	Projectiles    []*projectile.Projectile
}

func main() {
	rl.InitWindow(
		constants.WINDOW_SIZE_X,
		constants.WINDOW_SIZE_Y,
		"ASTEROIDS",
	)
	rl.InitAudioDevice()

	defer rl.CloseWindow()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)

	state := prepareGame()
	audioPlayer.Init()

	for !rl.WindowShouldClose() {

		update(&state)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		draw(state)
		rl.EndDrawing()
	}
}

func prepareGame() GameState {
	return prepareLevel([]*asteroid.Asteroid{}, 5, 0)
}

func prepareLevel(existingAsteroids []*asteroid.Asteroid, numLives, score int) GameState {
	var state GameState

	if len(existingAsteroids) > 0 {
		state.AsteroidField = existingAsteroids
	} else {
		state.AsteroidField = loadLevel()
	}

	state.PlayerShip = player.PlayerShip{}
	state.PlayerNumLives = numLives
	state.GameParticles = []particle.IParticle{}
	state.Projectiles = []*projectile.Projectile{}
	state.PlayerScore = score

	state.PlayerShip.Init()

	return state
}

func update(state *GameState) {
	updateAsteroids(state)
	updateParticles(state)
	updateProjectiles(state)

	if !state.PlayerShip.IsDead() {
		updatePlayerAlive(state)
	} else {
		updatePlayerDead(state)
	}
}

func draw(state GameState) {
	// Player Ship
	if !state.PlayerShip.IsDead() {
		player.DrawShip(
			state.PlayerShip.Position,
			state.PlayerShip.Rotation,
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
			constants.SCALE*a.GetSizeModifier(),
			a.Rotation,
			a.Shape,
		)
	}

	// Particles
	for _, pt := range state.GameParticles {
		pt.Draw()
	}

	// Projectiles
	for _, pr := range state.Projectiles {
		pr.Draw()
	}

	// Player Lives Remaining
	for i := range state.PlayerNumLives {
		player.DrawShip(
			rl.Vector2{
				X: float32(i*20) + constants.SCALE,
				Y: constants.SCALE,
			},
			0,
		)
	}

	// Player Score
	rl.DrawText(
		strconv.Itoa(state.PlayerScore),
		constants.WINDOW_SIZE_X-100,
		30,
		32,
		rl.White,
	)
}

func loadLevel() []*asteroid.Asteroid {
	numAsteroids := 10
	asteroids := make([]*asteroid.Asteroid, 0, numAsteroids)

	for range numAsteroids {
		size := utils.Rng.Int32N(int32(asteroid.NumAsteroidSizes))
		asteroids = append(asteroids, asteroid.New(asteroid.AsteroidSize(size)))
	}

	return asteroids
}

func updateAsteroids(state *GameState) {
	for _, a := range state.AsteroidField {
		a.Move()
	}
}

func updateParticles(state *GameState) {
	newParticles := (state.GameParticles)[:0]
	for _, pt := range state.GameParticles {
		pt.Update()
		if !pt.IsDead() {
			newParticles = append(newParticles, pt)
		}
	}
	state.GameParticles = newParticles
}

func updateProjectiles(state *GameState) {
	if len(state.Projectiles) == 0 {
		return
	}

	newProjectiles := (state.Projectiles)[:0]
	newAsteroids := []*asteroid.Asteroid{}
	hitIndices := map[int]bool{}

	for _, pr := range state.Projectiles {
		pr.Update()
		if pr.IsDead() {
			continue
		}

		collided := false
		for i, a := range state.AsteroidField {
			dist := rl.Vector2Distance(pr.Position, a.Position)
			asteroidRadius := float32(a.GetSize() * a.GetCollisionScale() * constants.SCALE)
			projectileRadius := pr.GetSize()

			if dist < (asteroidRadius + projectileRadius) {
				collided = true

				var pointsEarned int
				switch a.Size {
				case asteroid.BIG:
					pointsEarned = 50
				case asteroid.MEDIUM:
					pointsEarned = 100
				case asteroid.SMALL:
					pointsEarned = 200
				}
				state.PlayerScore += pointsEarned

				state.GameParticles = append(state.GameParticles, particle.CreateExplosion(a.Position, 12)...)
				newAsteroids = append(newAsteroids, a.Split()...)

				hitIndices[i] = true
				break
			}
		}

		if !collided {
			newProjectiles = append(newProjectiles, pr)
		}
	}

	for i, a := range state.AsteroidField {
		if !hitIndices[i] {
			newAsteroids = append(newAsteroids, a)
		}
	}

	state.Projectiles = newProjectiles
	state.AsteroidField = newAsteroids
}

func updatePlayerAlive(state *GameState) {
	state.PlayerShip.HandleInput()

	newAsteroids := []*asteroid.Asteroid{}
	hitIndices := map[int]bool{}

	for i, a := range state.AsteroidField {
		dist := rl.Vector2Distance(a.Position, state.PlayerShip.Position)
		asteroidRadius := float32(a.GetSize() * a.GetCollisionScale() * constants.SCALE)
		playerRadius := state.PlayerShip.GetCollisionRadius()

		if dist < (asteroidRadius + playerRadius) {
			state.PlayerShip.DeathTime = constants.PLAYER_DEATH_COOLDOWN
			state.PlayerNumLives -= 1

			// Particle Explosion
			state.GameParticles = append(
				state.GameParticles,
				particle.CreateExplosion(state.PlayerShip.Position, 30)...,
			)

			// Ship Explosion
			state.GameParticles = append(
				state.GameParticles,
				particle.CreateShipExplosion(state.PlayerShip.Position, 5)...,
			)

			// Split the asteroid
			newAsteroids = append(newAsteroids, a.Split()...)
			hitIndices[i] = true
			break
		}
	}

	for i, a := range state.AsteroidField {
		if !hitIndices[i] {
			newAsteroids = append(newAsteroids, a)
		}
	}
	state.AsteroidField = newAsteroids

	if state.PlayerShip.IsFiring {
		angle := state.PlayerShip.Rotation - math.Pi/2
		direction := rl.Vector2{
			X: float32(math.Cos(float64(angle))),
			Y: float32(math.Sin(float64(angle))),
		}
		speed := float32(4.0)

		state.Projectiles = append(state.Projectiles, &projectile.Projectile{
			Position: state.PlayerShip.Position,
			Velocity: rl.Vector2Scale(direction, speed),
			Size:     rl.Vector2{X: 2, Y: 2},
			Ttl:      2,
		})

		audioPlayer.PlayLazer()
	}
}

func updatePlayerDead(state *GameState) {
	state.PlayerShip.DeathTime -= float64(rl.GetFrameTime())
	state.PlayerShip.IsThrusting = false

	if state.PlayerShip.DeathTime <= 0.0 {
		var newState GameState
		if state.PlayerNumLives > 0 {
			newState = prepareLevel(state.AsteroidField, state.PlayerNumLives, state.PlayerScore)
		} else {
			newState = prepareGame()
		}
		state.PlayerShip = newState.PlayerShip
		state.PlayerNumLives = newState.PlayerNumLives
		state.PlayerScore = newState.PlayerScore
		state.AsteroidField = newState.AsteroidField
		state.GameParticles = newState.GameParticles
		state.Projectiles = newState.Projectiles
	}
}
