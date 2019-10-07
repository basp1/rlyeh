package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Image struct {
	id int32

	Parent Widget
	Bounds rl.Rectangle

	Align Align
	Fill  Fill

	texture   rl.Texture2D
	rectangle rl.Rectangle
}

func NewImage(align Align, fill Fill, texture rl.Texture2D, rectangle *rl.Rectangle) *Image {
	self := &Image{}

	self.texture = texture

	self.Align = align
	self.Fill = fill

	if nil != rectangle {
		self.rectangle = *rectangle
	} else {
		self.rectangle.X = 0
		self.rectangle.Y = 0
		self.rectangle.Width = float32(texture.Width)
		self.rectangle.Height = float32(texture.Height)
	}

	return self
}

func (self *Image) GetId() int32 {
	return self.id
}

func (self *Image) SetId(id int32) {
	self.id = id
}

func (self *Image) GetParent() Widget {
	return self.Parent
}

func (self *Image) SetParent(parent Widget) {
	self.Parent = parent
}

func (self *Image) GetBounds() rl.Rectangle {
	return self.Bounds
}

func (self *Image) SetBounds(bounds rl.Rectangle) {
	self.Bounds = bounds
}

func (self *Image) GetAlign() Align {
	return self.Align
}

func (self *Image) GetFill() Fill {
	return self.Fill
}

func (self *Image) GetDataSize() Size {
	var size Size

	size.Width = self.rectangle.Width
	size.Height = self.rectangle.Height

	return size
}

func (self *Image) Update(dt float32) {
}

func (self *Image) Draw() {
	sourceRec := self.rectangle
	destRec := self.Bounds
	origin := rl.Vector2{}
	rotation := float32(0)
	color := rl.Blank

	rl.DrawTexturePro(self.texture, sourceRec, destRec, origin, rotation, color)
}
