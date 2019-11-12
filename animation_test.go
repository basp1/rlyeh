package rlyeh

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/stretchr/testify/assert"
)

func TestAnimation(T *testing.T) {
	a := NewAnimation(rl.Texture2D{}, rl.Vector2{X: 0, Y: 0}, Size{Width: 32, Height: 32}, 2)

	frame := a.Get(0)

	assert.Equal(T, float32(0), frame.X)
	assert.Equal(T, float32(0), frame.Y)

	frame = a.Get(1)

	assert.Equal(T, float32(32), frame.X)
	assert.Equal(T, float32(0), frame.Y)

	frame = a.Get(2)

	assert.Equal(T, float32(0), frame.X)
	assert.Equal(T, float32(0), frame.Y)

	frame = a.Get(3)

	assert.Equal(T, float32(32), frame.X)
	assert.Equal(T, float32(0), frame.Y)
}
