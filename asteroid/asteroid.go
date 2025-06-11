package asteroid

import (
	"math/rand/v2"
	"time"

	"github.com/JaxonAdams/go-asteroids/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var s = rand.NewPCG(42, uint64(time.Now().Unix()))
var rng = rand.New(s)

type Asteroid struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Speed    rl.Vector2
	Rotation float32
	Shape    []rl.Vector2
}

func New() *Asteroid {
	pos := rl.Vector2{
		X: rng.Float32() * constants.WINDOW_SIZE_X,
		Y: rng.Float32() * constants.WINDOW_SIZE_Y,
	}

	shape := []rl.Vector2{
		{X: 0.0, Y: 1.0},
		{X: 1.0, Y: 1.0},
		{X: 1.0, Y: 0.0},
		{X: 0.0, Y: 0.0},
	}

	return &Asteroid{
		Position: pos,
		Shape:    shape,
	}
}
