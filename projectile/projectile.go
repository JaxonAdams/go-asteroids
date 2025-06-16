package projectile

import rl "github.com/gen2brain/raylib-go/raylib"

type Projectile struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Ttl      float32
}

func (p *Projectile) Update() {
	p.Position = rl.Vector2Add(p.Position, p.Velocity)
	p.Ttl -= rl.GetFrameTime()
}

func (p *Projectile) Draw() {
	size := rl.Vector2{X: 2, Y: 2}
	centeredPos := rl.Vector2Subtract(p.Position, rl.Vector2{X: size.X / 2, Y: size.Y / 2})
	rl.DrawRectangleV(centeredPos, size, rl.White)
}

func (p *Projectile) IsDead() bool {
	return p.Ttl <= 0
}
