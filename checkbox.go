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

	size.Width = float32(style[LabelBorderWidth]) + float32(style[GlobalTextFontsize])
	size.Height = size.Width

	return size
}

func (self *Checkbox) Update(dt float32) {
	state := GetState(self.Bounds)
	if Released == state {
		self.Checked = !self.Checked
	}
	self.state = state
}

func (self *Checkbox) Draw() {
	b := self.Bounds.ToInt32()
	borderWidth := int32(style[LabelBorderWidth])

	switch self.state {
	case Normal:
		rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, GetColor(CheckboxDefaultBorderColor))
		rl.DrawRectangle(b.X+borderWidth, b.Y+borderWidth, b.Width-(2*borderWidth), b.Height-(2*borderWidth), GetColor(CheckboxDefaultInsideColor))
		break
	case Focused:
		rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, GetColor(CheckboxHoverBorderColor))
		rl.DrawRectangle(b.X+borderWidth, b.Y+borderWidth, b.Width-(2*borderWidth), b.Height-(2*borderWidth), GetColor(CheckboxHoverInsideColor))
		break
	case Pressed:
		rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, GetColor(CheckboxClickBorderColor))
		rl.DrawRectangle(b.X+borderWidth, b.Y+borderWidth, b.Width-(2*borderWidth), b.Height-(2*borderWidth), GetColor(CheckboxClickInsideColor))
		break
	default:
		break
	}

	if self.Checked {
		rl.DrawRectangle(b.X+int32(style[CheckboxInsideWidth]), b.Y+int32(style[CheckboxInsideWidth]), b.Width-(2*int32(style[CheckboxInsideWidth])), b.Height-(2*int32(style[CheckboxInsideWidth])), rl.GetColor(int32(style[CheckboxDefaultActiveColor])))
	}
}
