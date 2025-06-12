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
	Position  rl.Vector2
	Velocity  rl.Vector2
	Speed     float32
	Rotation  float32
	Shape     []rl.Vector2
	Size      AsteroidSize
	AvgRadius float64
}

const (
	BIG AsteroidSize = iota
	MEDIUM
	SMALL
	NumAsteroidSizes
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

	shape, avgRadius := getRandShape(size)

	return &Asteroid{
		Position:  pos,
		Velocity:  vel,
		Speed:     100.0,
		Rotation:  rot,
		Shape:     shape,
		Size:      size,
		AvgRadius: avgRadius,
	}
}

func (a Asteroid) GetSize() float64 {
	return a.AvgRadius
}

func (a Asteroid) GetCollisionScale() float64 {
	switch a.Size {
	case BIG:
		return 0.8
	case MEDIUM:
		return 0.6
	case SMALL:
		return 0.6
	}

	return 1.0
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

func getRandShape(size AsteroidSize) ([]rl.Vector2, float64) {
	var points []rl.Vector2

	var sizeMultiplier float32
	switch size {
	case BIG:
		sizeMultiplier = 2.0
	case MEDIUM:
		sizeMultiplier = 1.0
	case SMALL:
		sizeMultiplier = 0.5
	}

	numPoints := rng.Int32N(4) + 7 // 7–10 points
	baseAngle := 2 * math.Pi / float64(numPoints)

	var allPointsRadius []float64
	for i := int32(0); i < numPoints; i++ {
		// Add small angle jitter for non-uniform spacing
		angleJitter := rng.Float64()*baseAngle*0.4 - baseAngle*0.2
		angle := float64(i)*baseAngle + angleJitter

		radius := 0.4 + rng.Float64()*0.6*float64(sizeMultiplier)
		allPointsRadius = append(allPointsRadius, radius)

		// Polar to Cartesian
		x := float32(math.Cos(angle) * radius)
		y := float32(math.Sin(angle) * radius)

		points = append(points, rl.Vector2{X: x, Y: y})
	}

	return points, utils.CalculateAverage(allPointsRadius)
}
