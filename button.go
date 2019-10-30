package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	id   int32
	Text string

	Parent Widget

	Bounds rl.Rectangle
	Align  Align
	Fill   Fill
	State  State

	OnClick func()
}

func NewButton(align Align, fill Fill, text string, onClick func()) *Button {
	self := &Button{}

	self.Align = align
	self.Fill = fill

	self.Text = text

	self.OnClick = onClick

	return self
}

func (self *Button) GetId() int32 {
	return self.id
}

func (self *Button) SetId(id int32) {
	self.id = id
}

func (self *Button) GetParent() Widget {
	return self.Parent
}

func (self *Button) SetParent(parent Widget) {
	self.Parent = parent
}

func (self *Button) GetBounds() rl.Rectangle {
	return self.Bounds
}

func (self *Button) SetBounds(bounds rl.Rectangle) {
	self.Bounds = bounds
}

func (self *Button) GetAlign() Align {
	return self.Align
}

func (self *Button) GetFill() Fill {
	return self.Fill
}

func (self *Button) GetDataSize() Size {
	var size Size

	size.Width = float32(rl.MeasureText(self.Text, int32(style.GlobalTextFontsize)))
	size.Height = float32(style.GlobalTextFontsize)

	size.Width += float32(style.ButtonTextPadding)
	size.Height += float32(style.ButtonTextPadding) / 2

	return size
}

func (self *Button) Update(dt float32) {
	state := GetState(self.Bounds)

	if Released == state && Pressed == self.State && nil != self.OnClick {
		self.OnClick()
	}

	self.State = state
}

func (self *Button) Draw() {
	b := self.Bounds.ToInt32()
	state := self.State
	text := self.Text

	switch state {
	case Released:
		rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.ButtonDefaultBorderColor)
		rl.DrawRectangle(b.X+int32(style.ButtonBorderWidth), b.Y+int32(style.ButtonBorderWidth), b.Width-(2*int32(style.ButtonBorderWidth)), b.Height-(2*int32(style.ButtonBorderWidth)), style.ButtonDefaultInsideColor)
		rl.DrawText(text, b.X+((b.Width/2)-(rl.MeasureText(text, int32(style.GlobalTextFontsize))/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), style.ButtonDefaultTextColor)
	case Normal:
		rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.ButtonDefaultBorderColor)
		rl.DrawRectangle(b.X+int32(style.ButtonBorderWidth), b.Y+int32(style.ButtonBorderWidth), b.Width-(2*int32(style.ButtonBorderWidth)), b.Height-(2*int32(style.ButtonBorderWidth)), style.ButtonDefaultInsideColor)
		rl.DrawText(text, b.X+((b.Width/2)-(rl.MeasureText(text, int32(style.GlobalTextFontsize))/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), style.ButtonDefaultTextColor)
		break

	case Focused:
		rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.ButtonHoverBorderColor)
		rl.DrawRectangle(b.X+int32(style.ButtonBorderWidth), b.Y+int32(style.ButtonBorderWidth), b.Width-(2*int32(style.ButtonBorderWidth)), b.Height-(2*int32(style.ButtonBorderWidth)), style.ButtonHoverInsideColor)
		rl.DrawText(text, b.X+((b.Width/2)-(rl.MeasureText(text, int32(style.GlobalTextFontsize))/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), style.ButtonHoverTextColor)
		break

	case Pressed:
		rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.ButtonPressedBorderColor)
		rl.DrawRectangle(b.X+int32(style.ButtonBorderWidth), b.Y+int32(style.ButtonBorderWidth), b.Width-(2*int32(style.ButtonBorderWidth)), b.Height-(2*int32(style.ButtonBorderWidth)), style.ButtonPressedInsideColor)
		rl.DrawText(text, b.X+((b.Width/2)-(rl.MeasureText(text, int32(style.GlobalTextFontsize))/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), style.ButtonPressedTextColor)
		break

	default:
		break
	}
}
