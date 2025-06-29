package particle

import (
	"math"
	"math/rand/v2"

	"github.com/JaxonAdams/go-asteroids/constants"
	"github.com/JaxonAdams/go-asteroids/utils"
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

func (p *DotParticle) Update() {
	p.Position = rl.Vector2Add(p.Position, p.Velocity)
	p.Ttl -= rl.GetFrameTime()
}

func (p *DotParticle) Draw() {
	rl.DrawPixelV(p.Position, rl.White)
}

func (p *DotParticle) IsDead() bool {
	return p.Ttl <= 0
}

func CreateExplosion(position rl.Vector2, count int) []IParticle {
	var particles []IParticle
	for range count {
		angle := rand.Float32() * 2 * math.Pi
		speed := 1.0 + rand.Float32()*2.0

		particles = append(particles, &DotParticle{
			Particle: Particle{
				Position: position,
				Velocity: rl.Vector2{
					X: speed * float32(math.Cos(float64(angle))),
					Y: speed * float32(math.Sin(float64(angle))),
				},
				Ttl: 0.6 + rand.Float32()*0.4,
			},
		})
	}
	return particles
}

func CreateShipExplosion(position rl.Vector2, count int) []IParticle {
	var particles []IParticle
	for range count {
		angle := 2 * math.Pi * utils.Rng.Float64()
		newParticle := &LineParticle{
			Rotation: 2 * math.Pi * utils.Rng.Float32(),
			Length:   constants.SCALE * (0.6 + (0.4 * utils.Rng.Float32())),
			Particle: Particle{
				Position: rl.Vector2Add(
					position,
					rl.Vector2{X: utils.Rng.Float32() * 3, Y: utils.Rng.Float32() * 3},
				),
				Velocity: rl.Vector2Scale(
					rl.Vector2{X: float32(math.Cos(angle)), Y: float32(math.Sin(angle))},
					2.0*utils.Rng.Float32(),
				),
				Ttl: 3.0 + utils.Rng.Float32(),
			},
		}
		particles = append(particles, newParticle)
	}
	return particles
}
