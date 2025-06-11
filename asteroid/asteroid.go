package asteroid

import (
	"math"
	"math/rand/v2"
	"time"

	"github.com/JaxonAdams/go-asteroids/constants"
	"github.com/JaxonAdams/go-asteroids/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AsteroidSize int

type Asteroid struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Speed    float32
	Rotation float32
	Shape    []rl.Vector2
	Size     AsteroidSize
}

const (
	BIG AsteroidSize = iota
	MEDIUM
	SMALL
)

var s = rand.NewPCG(42, uint64(time.Now().Unix()))
var rng = rand.New(s)

func (a *Asteroid) Move() {
	dt := rl.GetFrameTime()
	velocity := rl.Vector2Scale(a.Velocity, a.Speed*dt)
	a.Position = rl.Vector2Add(a.Position, velocity)
	a.Position = utils.ScreenWraparound(a.Position)
}

func New(size AsteroidSize) *Asteroid {

	pos := getRandPosition()
	rot := getRandRotation()
	vel := getRandVelocity(rot)

	shape := getRandShape(size)

	return &Asteroid{
		Position: pos,
		Velocity: vel,
		Speed:    100.0,
		Rotation: rot,
		Shape:    shape,
		Size:     size,
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

func getRandShape(size AsteroidSize) []rl.Vector2 {
	var points []rl.Vector2

	var sizeMultiplier float32
	switch size {
	case BIG:
		sizeMultiplier = 1.2
	case MEDIUM:
		sizeMultiplier = 0.7
	case SMALL:
		sizeMultiplier = 0.2
	}

	numPoints := rng.Int32N(4) + 7 // 7–10 points
	baseAngle := 2 * math.Pi / float64(numPoints)

	for i := int32(0); i < numPoints; i++ {
		// Add small angle jitter for non-uniform spacing
		angleJitter := rng.Float64()*baseAngle*0.4 - baseAngle*0.2
		angle := float64(i)*baseAngle + angleJitter

		// Allow deeper valleys by lowering the minimum radius
		radius := 0.4 + rng.Float64()*0.6*float64(sizeMultiplier)

		// Polar to Cartesian
		x := float32(math.Cos(angle) * radius)
		y := float32(math.Sin(angle) * radius)

		points = append(points, rl.Vector2{X: x, Y: y})
	}

	return points
}
