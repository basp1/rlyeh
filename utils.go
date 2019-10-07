package rlyeh

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func nextId() int32 {
	return rl.GetRandomValue(1, math.MaxInt32)
}

func max(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}

func min(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}
