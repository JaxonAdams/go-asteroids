package player

import (
	"math"

	"github.com/JaxonAdams/go-asteroids/constants"
	"github.com/JaxonAdams/go-asteroids/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var PlayerShape []rl.Vector2 = []rl.Vector2{
	{X: -0.2, Y: 0.3},
	{X: 0.0, Y: -0.3},
	{X: 0.2, Y: 0.3},
	{X: 0.1, Y: 0.2},
	{X: -0.1, Y: 0.2},
}

func DrawShip(pos rl.Vector2, rotation float32) {
	utils.DrawShape(
		pos,
		constants.SCALE,
		rotation,
		PlayerShape,
	)
}

type PlayerShip struct {
	Position    rl.Vector2
	Velocity    rl.Vector2
	Forward     rl.Vector2
	Shape       []rl.Vector2
	TailShape   []rl.Vector2
	Rotation    float32
	IsThrusting bool
	IsFiring    bool
	DeathTime   float64
}

func (p *PlayerShip) Init() {
	p.Shape = PlayerShape
	p.TailShape = []rl.Vector2{
		{X: 0.1, Y: 0.2},
		{X: 0.0, Y: 0.5},
		{X: -0.1, Y: 0.2},
	}

	p.Position = rl.Vector2Add(
		rl.Vector2Scale(rl.Vector2{
			X: constants.WINDOW_SIZE_X,
			Y: constants.WINDOW_SIZE_Y,
		}, 0.5),
		utils.ComputeCentroid(p.Shape),
	)
}

func (p *PlayerShip) HandleInput() {
	dt := rl.GetFrameTime()

	// Player Movement
	if rl.IsKeyDown(rl.KeyA) {
		p.Rotation -= constants.ROTATION_SPEED * dt
	}

	if rl.IsKeyDown(rl.KeyD) {
		p.Rotation += constants.ROTATION_SPEED * dt
	}

	if rl.IsKeyDown(rl.KeyW) {
		p.IsThrusting = true

		p.Forward = rl.Vector2{
			X: float32(math.Cos(float64(p.Rotation - math.Pi/2))),
			Y: float32(math.Sin(float64(p.Rotation - math.Pi/2))),
		}

		acceleration := rl.Vector2Scale(p.Forward, constants.PLAYER_SPEED*dt)

		p.Velocity = rl.Vector2Add(p.Velocity, acceleration)
	} else {
		p.IsThrusting = false
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		p.IsFiring = true
	} else {
		p.IsFiring = false
	}

	// Environmental Effects
	p.Velocity = rl.Vector2Scale(p.Velocity, 1.0-constants.DRAG)
	p.Position = rl.Vector2Add(p.Position, p.Velocity)
	p.Position = utils.ScreenWraparound(p.Position)
}

func (p PlayerShip) IsDead() bool {
	return p.DeathTime > 0
}

func (p PlayerShip) GetCollisionRadius() float32 {
	return constants.SCALE * 0.5
}
