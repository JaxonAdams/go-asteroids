package projectile

import rl "github.com/gen2brain/raylib-go/raylib"

type Projectile struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Size     rl.Vector2
	Ttl      float32
}

func (p *Projectile) Update() {
	p.Position = rl.Vector2Add(p.Position, p.Velocity)
	p.Ttl -= rl.GetFrameTime()
}

func (p *Projectile) Draw() {
	centeredPos := rl.Vector2Subtract(p.Position, rl.Vector2{X: p.Size.X / 2, Y: p.Size.Y / 2})
	rl.DrawRectangleV(centeredPos, p.Size, rl.White)
}

func (p *Projectile) IsDead() bool {
	return p.Ttl <= 0
}

func (p Projectile) GetSize() float32 {
	return p.Size.X / 2
}
