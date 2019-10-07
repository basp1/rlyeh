package rlyeh

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Textbox struct {
	id   int32
	Text string

	Parent Widget

	Bounds rl.Rectangle
	Align  Align
	Fill   Fill

	framesCounter int
	state         State

	runeCount int
	text      string
}

func NewTextbox(align Align, fill Fill, runeCount int) *Textbox {
	self := &Textbox{}

	self.framesCounter = 0
	self.state = Normal

	self.Align = align
	self.Fill = fill

	self.runeCount = runeCount
	self.text = ""

	return self
}

func (self *Textbox) GetId() int32 {
	return self.id
}

func (self *Textbox) SetId(id int32) {
	self.id = id
}

func (self *Textbox) GetParent() Widget {
	return self.Parent
}

func (self *Textbox) SetParent(parent Widget) {
	self.Parent = parent
}

func (self *Textbox) GetBounds() rl.Rectangle {
	return self.Bounds
}

func (self *Textbox) SetBounds(bounds rl.Rectangle) {
	self.Bounds = bounds
}

func (self *Textbox) GetAlign() Align {
	return self.Align
}

func (self *Textbox) GetFill() Fill {
	return self.Fill
}

func (self *Textbox) GetDataSize() Size {
	var size Size

	size.Width = float32(self.runeCount) * float32(rl.MeasureText("A", int32(style[GlobalTextFontsize])))
	size.Height = float32(style[GlobalTextFontsize])

	size.Width += float32(style[LabelTextPadding])
	size.Height += float32(style[LabelTextPadding]) / 2

	return size
}

func (self *Textbox) Update(dt float32) {
	letter := int32(-1)
	inBounds := rl.CheckCollisionPointRec(rl.GetMousePosition(), self.Bounds)

	if inBounds {
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			self.state = Pressed
		} else if Pressed != self.state {
			self.state = Focused
		}
	} else {
		if self.state == Focused {
			self.state = Normal
		}
		if self.state == Pressed && rl.IsMouseButtonDown(rl.MouseLeftButton) {
			self.state = Normal
		}
	}

	if inBounds && Focused == self.state && rl.IsMouseButtonDown(rl.MouseLeftButton) {
		self.state = Pressed
	}

	if Pressed == self.state {
		self.framesCounter++

		letter = rl.GetKeyPressed()
		if letter != -1 {
			if letter >= 32 && letter < 127 {
				self.text = fmt.Sprintf("%s%c", self.text, letter)
			}
		}

		if rl.IsKeyPressed(rl.KeyBackspace) {
			if len(self.text) > 0 {
				self.text = self.text[:len(self.text)-1]
			}
		}
	}
}

func (self *Textbox) Draw() {
	bounds := self.Bounds

	boundsInside := bounds
	boundsInside.X += float32(style[TextboxBorderWidth])
	boundsInside.Y += float32(style[TextboxBorderWidth])
	boundsInside.Width -= float32(style[TextboxBorderWidth]) * 2
	boundsInside.Height -= float32(style[TextboxBorderWidth]) * 2

	textPointX := int32(bounds.X) + 2
	textPointY := int32(bounds.Y) + int32(style[TextboxBorderWidth]) + int32(bounds.Height/2) - int32(style[TextboxTextFontsize])/2

	runeWidth := float32(rl.MeasureText("A", int32(style[GlobalTextFontsize])))
	runeCount := len(self.text)
	truncated := self.text[runeCount-int(min(float32(runeCount), boundsInside.Width/runeWidth)):]

	switch self.state {
	case Normal:
		rl.DrawRectangleRec(bounds, GetColor(TextboxBorderColor))
		rl.DrawRectangleRec(boundsInside, GetColor(TextboxInsideColor))
		rl.DrawText(truncated, textPointX, textPointY, int32(style[TextboxTextFontsize]), GetColor(TextboxTextColor))
		break

	case Focused:
		rl.DrawRectangleRec(bounds, GetColor(TextboxActiveBorderColor))
		rl.DrawRectangleRec(boundsInside, GetColor(TextboxInsideColor))
		rl.DrawText(truncated, textPointX, textPointY, int32(style[TextboxTextFontsize]), GetColor(TextboxTextColor))
		break

	case Pressed:
		rl.DrawRectangleRec(bounds, GetColor(TextboxActiveBorderColor))
		rl.DrawRectangleRec(boundsInside, GetColor(TextboxInsideColor))
		rl.DrawText(truncated, textPointX, textPointY, int32(style[TextboxTextFontsize]), GetColor(TextboxTextColor))

		if (self.framesCounter/20)%2 == 0 {
			rl.DrawRectangle(int32(bounds.X)+4+rl.MeasureText(truncated, int32(style[GlobalTextFontsize])),
				int32(bounds.Y+2),
				1,
				int32(bounds.Height-4), GetColor(TextboxLineColor))
		}
		break
	default:
		break
	}
}
