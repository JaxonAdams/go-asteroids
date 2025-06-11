package asteroid

import (
	"math"
	"math/rand/v2"
	"time"

	"github.com/JaxonAdams/go-asteroids/constants"
	"github.com/JaxonAdams/go-asteroids/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var s = rand.NewPCG(42, uint64(time.Now().Unix()))
var rng = rand.New(s)

type Asteroid struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Speed    float32
	Rotation float32
	Shape    []rl.Vector2
}

func (a *Asteroid) Move() {
	dt := rl.GetFrameTime()
	velocity := rl.Vector2Scale(a.Velocity, a.Speed*dt)
	a.Position = rl.Vector2Add(a.Position, velocity)
	a.Position = utils.ScreenWraparound(a.Position)
}

func New() *Asteroid {

	pos := getRandPosition()
	rot := getRandRotation()
	vel := getRandVelocity(rot)

	shape := []rl.Vector2{
		{X: 0.0, Y: 1.0},
		{X: 1.0, Y: 1.0},
		{X: 1.0, Y: 0.0},
		{X: 0.0, Y: 0.0},
	}

	return &Asteroid{
		Position: pos,
		Velocity: vel,
		Speed:    100.0,
		Rotation: rot,
		Shape:    shape,
	}
}

func getRandPosition() rl.Vector2 {
	return rl.Vector2{
		X: rng.Float32() * constants.WINDOW_SIZE_X,
		Y: rng.Float32() * constants.WINDOW_SIZE_Y,
	}
}

func getRandRotation() float32 {
	return rng.Float32() * (2 * math.Pi)
}

func getRandVelocity(rotation float32) rl.Vector2 {
	return rl.Vector2{
		X: float32(math.Cos(float64(rotation))),
		Y: float32(math.Sin(float64(rotation))),
	}
}
