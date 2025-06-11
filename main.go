package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const WINDOW_SIZE_X = 800
const WINDOW_SIZE_Y = 450
const SCALE = 40
const PLAYER_SPEED = 20.0
const ROTATION_SPEED = 1.5 * (2 * math.Pi)
const DRAG = 0.01

type Player struct {
	Position    rl.Vector2
	Velocity    rl.Vector2
	Shape       []rl.Vector2
	TailShape   []rl.Vector2
	Rotation    float32
	IsThrusting bool
}

var player Player

func main() {
	rl.InitWindow(WINDOW_SIZE_X, WINDOW_SIZE_Y, "ASTEROIDS")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	player.Shape = []rl.Vector2{
		{X: -0.2, Y: 0.3},
		{X: 0.0, Y: -0.3},
		{X: 0.2, Y: 0.3},
		{X: 0.1, Y: 0.2},
		{X: -0.1, Y: 0.2},
	}

	player.TailShape = []rl.Vector2{
		{X: 0.1, Y: 0.2},
		{X: 0.0, Y: 0.5},
		{X: -0.1, Y: 0.2},
	}

	player.Position = rl.Vector2Add(
		rl.Vector2Scale(rl.Vector2{
			X: WINDOW_SIZE_X,
			Y: WINDOW_SIZE_Y,
		}, 0.5),
		computeCentroid(player.Shape),
	)

	for !rl.WindowShouldClose() {

		update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		draw()
		rl.EndDrawing()
	}
}

func update() {
	handleInput()
}

func draw() {
	// Player Ship
	drawShape(
		player.Position,
		SCALE,
		player.Rotation,
		player.Shape,
	)

	shouldDrawTail := int32(rl.GetTime()*20)%2 == 0
	if player.IsThrusting && shouldDrawTail {
		// Player Thrust Tail
		drawShape(
			player.Position,
			SCALE,
			player.Rotation,
			player.TailShape,
		)
	}
}

func handleInput() {
	dt := rl.GetFrameTime()

	// Player Movement
	if rl.IsKeyDown(rl.KeyA) {
		player.Rotation -= ROTATION_SPEED * dt
	}

	if rl.IsKeyDown(rl.KeyD) {
		player.Rotation += ROTATION_SPEED * dt
	}

	if rl.IsKeyDown(rl.KeyW) {
		player.IsThrusting = true

		forward := rl.Vector2{
			X: float32(math.Cos(float64(player.Rotation - math.Pi/2))),
			Y: float32(math.Sin(float64(player.Rotation - math.Pi/2))),
		}

		acceleration := rl.Vector2Scale(forward, PLAYER_SPEED*dt)

		player.Velocity = rl.Vector2Add(player.Velocity, acceleration)
	} else {
		player.IsThrusting = false
	}

	// Environmental Effects
	player.Velocity = rl.Vector2Scale(player.Velocity, 1.0-DRAG)
	player.Position = rl.Vector2Add(player.Position, player.Velocity)
	player.Position = screenWraparound(player.Position)
}

func drawShape(pos rl.Vector2, scale float32, rotation float32, points []rl.Vector2) {
	transformer := func(point rl.Vector2) rl.Vector2 {
		scaled := rl.Vector2Scale(point, scale)
		rotated := rl.Vector2Rotate(scaled, rotation)
		return rl.Vector2Add(rotated, pos)
	}

	for i := range len(points) {
		current := transformer(points[i])
		next := transformer(points[(i+1)%len(points)])

		rl.DrawLineV(current, next, rl.White)
	}
}

func computeCentroid(points []rl.Vector2) rl.Vector2 {
	var sum rl.Vector2
	for _, p := range points {
		sum = rl.Vector2Add(sum, p)
	}
	return rl.Vector2Scale(sum, 1/float32(len(points)))
}

func screenWraparound(pos rl.Vector2) rl.Vector2 {
	newPos := pos

	if pos.X < 0 {
		newPos.X += WINDOW_SIZE_X
	}
	if pos.X > WINDOW_SIZE_X {
		newPos.X -= WINDOW_SIZE_X
	}
	if pos.Y < 0 {
		newPos.Y += WINDOW_SIZE_Y
	}
	if pos.Y > WINDOW_SIZE_Y {
		newPos.Y -= WINDOW_SIZE_Y
	}

	return newPos
}
