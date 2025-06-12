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

func CalculateAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
