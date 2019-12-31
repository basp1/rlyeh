package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Checkbox struct {
	id   int32
	Text string

	Parent Widget

	Bounds rl.Rectangle
	Align  Align
	Fill   Fill

	Checked bool
	state   State
}

func NewCheckbox(align Align, fill Fill) *Checkbox {
	self := &Checkbox{}

	self.Align = align
	self.Fill = fill

	self.Checked = false
	self.state = Normal

	return self
}

func (self *Checkbox) GetId() int32 {
	return self.id
}

func (self *Checkbox) SetId(id int32) {
	self.id = id
}

func (self *Checkbox) GetParent() Widget {
	return self.Parent
}

func (self *Checkbox) SetParent(parent Widget) {
	self.Parent = parent
}

func (self *Checkbox) GetBounds() rl.Rectangle {
	return self.Bounds
}

func (self *Checkbox) SetBounds(bounds rl.Rectangle) {
	self.Bounds = bounds
}

func (self *Checkbox) GetAlign() Align {
	return self.Align
}

func (self *Checkbox) GetFill() Fill {
	return self.Fill
}

func (self *Checkbox) GetDataSize() Size {
	var size Size

	size.Width = float32(style.LabelBorderWidth) + float32(style.GlobalTextFontsize)
	size.Height = size.Width

	return size
}

func (self *Checkbox) Update(dt float32) {
	state := GetState(self.Bounds)
	if Released == state && Released != self.state {
		self.Checked = !self.Checked
	}
	self.state = state
}

func (self *Checkbox) Draw() {
	b := self.Bounds.ToInt32()

	borderWidth := int32(style.LabelBorderWidth)

	if self.Checked {
		rl.DrawRectangleLines(b.X, b.Y, b.Width, b.Height, style.CheckboxDefaultBorderColor)

		rl.DrawLineEx(rl.Vector2{self.Bounds.X + 2*float32(borderWidth), self.Bounds.Y + 2*float32(borderWidth)},
			rl.Vector2{self.Bounds.X + self.Bounds.Width - 2*float32(borderWidth), self.Bounds.Y + self.Bounds.Height - 2*float32(borderWidth)},
			float32(borderWidth), style.CheckboxDefaultActiveColor)
		rl.DrawLineEx(rl.Vector2{self.Bounds.X + self.Bounds.Width - 2*float32(borderWidth), self.Bounds.Y + 2*float32(borderWidth)},
			rl.Vector2{self.Bounds.X + 2*float32(borderWidth), self.Bounds.Y + self.Bounds.Height - 2*float32(borderWidth)},
			float32(borderWidth), style.CheckboxDefaultActiveColor)
	} else {
		rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.CheckboxDefaultBorderColor)
		rl.DrawRectangle(b.X+borderWidth, b.Y+borderWidth, b.Width-(2*borderWidth), b.Height-(2*borderWidth), style.CheckboxDefaultInsideColor)

	}
}
