package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Animation struct {
	Texture rl.Texture2D

	point rl.Vector2
	size  Size

	frameCount byte
}

func NewAnimation(texture rl.Texture2D, point rl.Vector2, size Size, frameCount byte) *Animation {
	self := &Animation{}

	self.point = point
	self.size = size
	self.frameCount = frameCount

	self.Texture = texture

	return self
}

func (self *Animation) Get(frame byte) rl.Rectangle {
	current := frame

	if current >= self.frameCount {
		current = current % self.frameCount
	}

	return rl.Rectangle{
		X:      self.point.X + float32(current)*self.size.Width,
		Y:      self.point.Y,
		Width:  self.size.Width,
		Height: self.size.Height,
	}
}
