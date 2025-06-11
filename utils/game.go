package utils

import (
	"github.com/JaxonAdams/go-asteroids/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawShape(pos rl.Vector2, scale float32, rotation float32, points []rl.Vector2) {
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

func ScreenWraparound(pos rl.Vector2) rl.Vector2 {
	newPos := pos

	if pos.X < 0 {
		newPos.X += constants.WINDOW_SIZE_X
	}
	if pos.X > constants.WINDOW_SIZE_X {
		newPos.X -= constants.WINDOW_SIZE_X
	}
	if pos.Y < 0 {
		newPos.Y += constants.WINDOW_SIZE_Y
	}
	if pos.Y > constants.WINDOW_SIZE_Y {
		newPos.Y -= constants.WINDOW_SIZE_Y
	}

	return newPos
}
