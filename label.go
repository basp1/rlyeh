package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	id   int32
	Text string

	Parent Widget

	Bounds rl.Rectangle
	Align  Align
	Fill   Fill
}

func NewLabel(align Align, fill Fill, text string) *Label {
	self := &Label{}

	self.Align = align
	self.Fill = fill

	self.Text = text

	return self
}

func (self *Label) GetId() int32 {
	return self.id
}

func (self *Label) SetId(id int32) {
	self.id = id
}

func (self *Label) GetParent() Widget {
	return self.Parent
}

func (self *Label) SetParent(parent Widget) {
	self.Parent = parent
}

func (self *Label) GetBounds() rl.Rectangle {
	return self.Bounds
}

func (self *Label) SetBounds(bounds rl.Rectangle) {
	self.Bounds = bounds
}

func (self *Label) GetAlign() Align {
	return self.Align
}

func (self *Label) GetFill() Fill {
	return self.Fill
}

func (self *Label) GetDataSize() Size {
	var size Size

	size.Width = float32(rl.MeasureText(self.Text, int32(style.GlobalTextFontsize)))
	size.Height = float32(style.GlobalTextFontsize)

	size.Width += float32(style.LabelTextPadding)
	size.Height += float32(style.LabelTextPadding) / 2

	return size
}

func (self *Label) Update(dt float32) {
}

func (self *Label) Draw() {
	textColor := style.LabelTextColor
	border := rl.NewColor(0, 0, 0, 0)
	inner := rl.NewColor(0, 0, 0, 0)

	b := self.Bounds.ToInt32()
	textWidth := rl.MeasureText(self.Text, int32(style.GlobalTextFontsize))

	rl.DrawRectangleRec(self.Bounds, border)
	rl.DrawRectangle(b.X+int32(style.LabelBorderWidth), b.Y+int32(style.LabelBorderWidth), b.Width-(2*int32(style.LabelBorderWidth)), b.Height-(2*int32(style.LabelBorderWidth)), inner)
	rl.DrawText(self.Text, b.X+((b.Width/2)-(textWidth/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), textColor)
}
