package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Combobox struct {
	id     int32
	Parent Widget

	Bounds rl.Rectangle
	Align  Align
	Fill   Fill

	strings []string
	active  int
	longest float32

	state    State
	unrolled bool
}

func NewCombobox(align Align, fill Fill, strings []string) *Combobox {
	self := &Combobox{}

	self.Align = align
	self.Fill = fill

	self.strings = strings

	self.active = 0
	self.state = Normal
	self.unrolled = false

	max := 0
	for i, str := range strings {
		if len(str) > len(strings[max]) {
			max = i
		}
	}

	self.longest = float32(rl.MeasureText("  "+self.strings[max], int32(style.GlobalTextFontsize)))

	return self
}

func (self *Combobox) GetId() int32 {
	return self.id
}

func (self *Combobox) SetId(id int32) {
	self.id = id
}

func (self *Combobox) GetParent() Widget {
	return self.Parent
}

func (self *Combobox) SetParent(parent Widget) {
	self.Parent = parent
}

func (self *Combobox) GetBounds() rl.Rectangle {
	return self.Bounds
}

func (self *Combobox) SetBounds(bounds rl.Rectangle) {
	self.Bounds = bounds
}

func (self *Combobox) GetAlign() Align {
	return self.Align
}

func (self *Combobox) GetFill() Fill {
	return self.Fill
}

func (self *Combobox) GetDataSize() Size {
	var size Size

	size.Width = self.longest
	size.Width += float32(style.ComboboxWidth)
	size.Width += float32(style.ComboboxPadding)

	if style.GlobalTextFontsize < style.ComboboxHeight {
		size.Height = float32(style.ComboboxHeight)
	} else {
		size.Height = float32(style.GlobalTextFontsize)
	}
	size.Height += float32(style.ComboboxPadding) / 2

	return size
}

func (self *Combobox) GetActive() int {
	return self.active
}

func (self *Combobox) Update(dt float32) {
	bounds := self.Bounds

	state := GetState(bounds)

	if !self.unrolled && Pressed != self.state && Pressed == state {
		self.unrolled = true
	}
	if self.unrolled && Pressed != self.state && Normal == state {
		self.unrolled = false
	}

	self.state = state
}

func (self *Combobox) Draw() {
	b := self.Bounds.ToInt32()

	if !self.unrolled {
		switch self.state {
		case Normal:
			rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.ComboboxDefaultBorderColor)
			rl.DrawRectangle(b.X+int32(style.ComboboxBorderWidth), b.Y+int32(style.ComboboxBorderWidth), b.Width-(2*int32(style.ComboboxBorderWidth)), b.Height-(2*int32(style.ComboboxBorderWidth)), style.ComboboxDefaultInsideColor)
			rl.DrawText(self.strings[self.active], b.X+((b.Width/2)-(rl.MeasureText(self.strings[self.active], int32(style.GlobalTextFontsize))/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), style.ComboboxDefaultTextColor)
			break
		case Focused:
			rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.ComboboxHoverBorderColor)
			rl.DrawRectangle(b.X+int32(style.ComboboxBorderWidth), b.Y+int32(style.ComboboxBorderWidth), b.Width-(2*int32(style.ComboboxBorderWidth)), b.Height-(2*int32(style.ComboboxBorderWidth)), style.ComboboxHoverInsideColor)
			rl.DrawText(self.strings[self.active], b.X+((b.Width/2)-(rl.MeasureText(self.strings[self.active], int32(style.GlobalTextFontsize))/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), style.ComboboxHoverTextColor)
			break
		case Pressed:
			rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.ComboboxPressedBorderColor)
			rl.DrawRectangle(b.X+int32(style.ComboboxBorderWidth), b.Y+int32(style.ComboboxBorderWidth), b.Width-(2*int32(style.ComboboxBorderWidth)), b.Height-(2*int32(style.ComboboxBorderWidth)), style.ComboboxPressedInsideColor)
			rl.DrawText(self.strings[self.active], b.X+((b.Width/2)-(rl.MeasureText(self.strings[self.active], int32(style.GlobalTextFontsize))/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), style.ComboboxPressedTextColor)
			break
		default:
			break
		}
	} else {
		itemCount := len(self.strings)

		for i := 0; i < itemCount; i++ {
			switch self.state {
			case Normal:
				rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.ComboboxDefaultBorderColor)
				rl.DrawRectangle(b.X+int32(style.ComboboxBorderWidth), b.Y+int32(style.ComboboxBorderWidth), b.Width-(2*int32(style.ComboboxBorderWidth)), b.Height-(2*int32(style.ComboboxBorderWidth)), style.ComboboxDefaultInsideColor)
				rl.DrawText(self.strings[i], b.X+((b.Width/2)-(rl.MeasureText(self.strings[i], int32(style.GlobalTextFontsize))/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), style.ComboboxDefaultTextColor)
				break
			case Focused:
				rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.ComboboxHoverBorderColor)
				rl.DrawRectangle(b.X+int32(style.ComboboxBorderWidth), b.Y+int32(style.ComboboxBorderWidth), b.Width-(2*int32(style.ComboboxBorderWidth)), b.Height-(2*int32(style.ComboboxBorderWidth)), style.ComboboxHoverInsideColor)
				rl.DrawText(self.strings[i], b.X+((b.Width/2)-(rl.MeasureText(self.strings[i], int32(style.GlobalTextFontsize))/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), style.ComboboxHoverTextColor)
				break
			case Pressed:
				rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.ComboboxPressedBorderColor)
				rl.DrawRectangle(b.X+int32(style.ComboboxBorderWidth), b.Y+int32(style.ComboboxBorderWidth), b.Width-(2*int32(style.ComboboxBorderWidth)), b.Height-(2*int32(style.ComboboxBorderWidth)), style.ComboboxPressedInsideColor)
				rl.DrawText(self.strings[i], b.X+((b.Width/2)-(rl.MeasureText(self.strings[i], int32(style.GlobalTextFontsize))/2)), b.Y+((b.Height/2)-(int32(style.GlobalTextFontsize)/2)), int32(style.GlobalTextFontsize), style.ComboboxPressedTextColor)
				break
			default:
				break
			}

			b.Y -= b.Height
		}
	}
}
