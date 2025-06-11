package utils

import (
	"github.com/JaxonAdams/go-asteroids/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
