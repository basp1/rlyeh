package rlyeh

import (
	"fmt"

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
}

func NewCombobox(align Align, fill Fill, strings []string) *Combobox {
	self := &Combobox{}

	self.Align = align
	self.Fill = fill

	self.strings = strings

	self.active = 0

	max := 0
	for i, str := range strings {
		if len(str) > len(strings[max]) {
			max = i
		}
	}
	self.longest = float32(rl.MeasureText("  "+self.strings[max], int32(style[GlobalTextFontsize])))

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
	size.Height = float32(style[GlobalTextFontsize])

	size.Width += float32(style[ComboboxWidth])
	size.Height += float32(style[ComboboxHeight])

	size.Width += float32(style[ComboboxPadding])
	size.Height += float32(style[ComboboxPadding]) / 2

	return size
}

func (self *Combobox) GetActive() int {
	return self.active
}

func (self *Combobox) Update(dt float32) {
}

func (self *Combobox) Draw() {
	bounds := self.Bounds
	bounds.Width -= float32(style[ComboboxWidth])
	bounds.Height -= float32(style[ComboboxHeight])
	bounds.Width -= float32(style[ComboboxPadding])
	bounds.Height -= float32(style[ComboboxPadding]) / 2

	b := bounds.ToInt32()
	state := Normal

	clicked := false
	click := rl.NewRectangle(bounds.X+bounds.Width+float32(style[ComboboxPadding]), bounds.Y, float32(style[ComboboxWidth]), float32(style[ComboboxHeight]))
	c := click.ToInt32()

	mousePoint := rl.GetMousePosition()

	itemCount := len(self.strings)
	for i := 0; i < itemCount; i++ {
		if i != self.active {
			continue
		}

		if rl.CheckCollisionPointRec(mousePoint, bounds) || rl.CheckCollisionPointRec(mousePoint, click) {
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				state = Pressed
			} else if rl.IsMouseButtonReleased(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				clicked = true
			} else {
				state = Focused
			}
		}

		switch state {
		case Normal:
			rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, GetColor(ComboboxDefaultBorderColor))
			rl.DrawRectangle(b.X+int32(style[ComboboxBorderWidth]), b.Y+int32(style[ComboboxBorderWidth]), b.Width-(2*int32(style[ComboboxBorderWidth])), b.Height-(2*int32(style[ComboboxBorderWidth])), GetColor(ComboboxDefaultInsideColor))

			rl.DrawRectangle(c.X, c.Y, c.Width, c.Height, GetColor(ComboboxDefaultBorderColor))
			rl.DrawRectangle(c.X+int32(style[ComboboxBorderWidth]), c.Y+int32(style[ComboboxBorderWidth]), c.Width-(2*int32(style[ComboboxBorderWidth])), c.Height-(2*int32(style[ComboboxBorderWidth])), GetColor(ComboboxDefaultInsideColor))
			rl.DrawText(fmt.Sprintf("%d/%d", self.active+1, itemCount), c.X+((c.Width/2)-(rl.MeasureText(fmt.Sprintf("%d/%d", self.active+1, itemCount), int32(style[GlobalTextFontsize]))/2)), c.Y+((c.Height/2)-(int32(style[GlobalTextFontsize])/2)), int32(style[GlobalTextFontsize]), GetColor(ComboboxDefaultListTextColor))
			rl.DrawText(self.strings[i], b.X+((b.Width/2)-(rl.MeasureText(self.strings[i], int32(style[GlobalTextFontsize]))/2)), b.Y+((b.Height/2)-(int32(style[GlobalTextFontsize])/2)), int32(style[GlobalTextFontsize]), GetColor(ComboboxDefaultTextColor))
			break
		case Focused:
			rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, GetColor(ComboboxHoverBorderColor))
			rl.DrawRectangle(b.X+int32(style[ComboboxBorderWidth]), b.Y+int32(style[ComboboxBorderWidth]), b.Width-(2*int32(style[ComboboxBorderWidth])), b.Height-(2*int32(style[ComboboxBorderWidth])), GetColor(ComboboxHoverInsideColor))

			rl.DrawRectangle(c.X, c.Y, c.Width, c.Height, GetColor(ComboboxHoverBorderColor))
			rl.DrawRectangle(c.X+int32(style[ComboboxBorderWidth]), c.Y+int32(style[ComboboxBorderWidth]), c.Width-(2*int32(style[ComboboxBorderWidth])), c.Height-(2*int32(style[ComboboxBorderWidth])), GetColor(ComboboxHoverInsideColor))
			rl.DrawText(fmt.Sprintf("%d/%d", self.active+1, itemCount), c.X+((c.Width/2)-(rl.MeasureText(fmt.Sprintf("%d/%d", self.active+1, itemCount), int32(style[GlobalTextFontsize]))/2)), c.Y+((c.Height/2)-(int32(style[GlobalTextFontsize])/2)), int32(style[GlobalTextFontsize]), GetColor(ComboboxHoverListTextColor))
			rl.DrawText(self.strings[i], b.X+((b.Width/2)-(rl.MeasureText(self.strings[i], int32(style[GlobalTextFontsize]))/2)), b.Y+((b.Height/2)-(int32(style[GlobalTextFontsize])/2)), int32(style[GlobalTextFontsize]), GetColor(ComboboxHoverTextColor))
			break
		case Pressed:
			rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, GetColor(ComboboxPressedBorderColor))
			rl.DrawRectangle(b.X+int32(style[ComboboxBorderWidth]), b.Y+int32(style[ComboboxBorderWidth]), b.Width-(2*int32(style[ComboboxBorderWidth])), b.Height-(2*int32(style[ComboboxBorderWidth])), GetColor(ComboboxPressedInsideColor))

			rl.DrawRectangle(c.X, c.Y, c.Width, c.Height, GetColor(ComboboxPressedListBorderColor))
			rl.DrawRectangle(c.X+int32(style[ComboboxBorderWidth]), c.Y+int32(style[ComboboxBorderWidth]), c.Width-(2*int32(style[ComboboxBorderWidth])), c.Height-(2*int32(style[ComboboxBorderWidth])), GetColor(ComboboxPressedListInsideColor))
			rl.DrawText(fmt.Sprintf("%d/%d", self.active+1, itemCount), c.X+((c.Width/2)-(rl.MeasureText(fmt.Sprintf("%d/%d", self.active+1, itemCount), int32(style[GlobalTextFontsize]))/2)), c.Y+((c.Height/2)-(int32(style[GlobalTextFontsize])/2)), int32(style[GlobalTextFontsize]), GetColor(ComboboxPressedListTextColor))
			rl.DrawText(self.strings[i], b.X+((b.Width/2)-(rl.MeasureText(self.strings[i], int32(style[GlobalTextFontsize]))/2)), b.Y+((b.Height/2)-(int32(style[GlobalTextFontsize])/2)), int32(style[GlobalTextFontsize]), GetColor(ComboboxPressedTextColor))
			break
		default:
			break
		}

		if clicked {
			rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, GetColor(ComboboxPressedBorderColor))
			rl.DrawRectangle(b.X+int32(style[ComboboxBorderWidth]), b.Y+int32(style[ComboboxBorderWidth]), b.Width-(2*int32(style[ComboboxBorderWidth])), b.Height-(2*int32(style[ComboboxBorderWidth])), GetColor(ComboboxPressedInsideColor))

			rl.DrawRectangle(c.X, c.Y, c.Width, c.Height, GetColor(ComboboxPressedListBorderColor))
			rl.DrawRectangle(c.X+int32(style[ComboboxBorderWidth]), c.Y+int32(style[ComboboxBorderWidth]), c.Width-(2*int32(style[ComboboxBorderWidth])), c.Height-(2*int32(style[ComboboxBorderWidth])), GetColor(ComboboxPressedListInsideColor))
			rl.DrawText(fmt.Sprintf("%d/%d", self.active+1, itemCount), c.X+((c.Width/2)-(rl.MeasureText(fmt.Sprintf("%d/%d", self.active+1, itemCount), int32(style[GlobalTextFontsize]))/2)), c.Y+((c.Height/2)-(int32(style[GlobalTextFontsize])/2)), int32(style[GlobalTextFontsize]), GetColor(ComboboxPressedListTextColor))
			rl.DrawText(self.strings[i], b.X+((b.Width/2)-(rl.MeasureText(self.strings[i], int32(style[GlobalTextFontsize]))/2)), b.Y+((b.Height/2)-(int32(style[GlobalTextFontsize])/2)), int32(style[GlobalTextFontsize]), GetColor(ComboboxPressedTextColor))
		}
	}

	if rl.CheckCollisionPointRec(mousePoint, bounds) || rl.CheckCollisionPointRec(mousePoint, click) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			self.active++
			if self.active >= itemCount {
				self.active = 0
			}
		}
	}
}
