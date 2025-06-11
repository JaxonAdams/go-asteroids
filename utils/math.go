package utils

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func ComputeCentroid(points []rl.Vector2) rl.Vector2 {
	var sum rl.Vector2
	for _, p := range points {
		sum = rl.Vector2Add(sum, p)
	}
	return rl.Vector2Scale(sum, 1/float32(len(points)))
}
