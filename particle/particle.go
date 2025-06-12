package particle

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type IParticle interface {
	Update()
	Draw()
	IsDead() bool
}

type Particle struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Ttl      float32
}

type LineParticle struct {
	Rotation float32
	Length   float32
	Particle
}

type DotParticle struct {
	Particle
}

func (p *LineParticle) Update() {
	p.Position = rl.Vector2Add(p.Position, p.Velocity)
	p.Ttl -= rl.GetFrameTime()
}

func (p *LineParticle) Draw() {
	end := rl.Vector2{
		X: p.Position.X + p.Length*float32(math.Cos(float64(p.Rotation))),
		Y: p.Position.Y + p.Length*float32(math.Sin(float64(p.Rotation))),
	}
	rl.DrawLineV(p.Position, end, rl.White)
}

func (p *LineParticle) IsDead() bool {
	return p.Ttl <= 0
}
