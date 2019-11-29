package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Combobox struct {
	label *Label

	dialog   *Dialog
	listview *ListView

	strings []string
	longest float32

	state State
}

func NewCombobox(align Align, fill Fill, items []string) *Combobox {
	self := &Combobox{}

	self.strings = items
	self.label = NewLabel(align, fill, items[0])
	self.label.BorderColor = style.ButtonDefaultBorderColor

	self.state = Normal

	max := 0
	for i, str := range items {
		if len(str) > len(items) {
			max = i
		}
	}

	self.longest = float32(rl.MeasureText("  "+self.strings[max], int32(style.GlobalTextFontsize)))
	self.label.MinWidth = 2 + self.longest

	self.dialog = NewDialog(rl.Rectangle{X: 0, Y: 0, Width: self.longest, Height: float32(len(items) * style.GlobalTextFontsize)}, "")
	self.dialog.Decoration = false

	count := len(items)
	if count > 9 {
		count = 9
	}
	self.listview = NewListView(items, count)
	if count == len(items) {
		self.listview.RemoveScrollbar()
	}
	self.listview.OnClick = func(item string) {
		self.label.Text = item
		self.dialog.Close()
	}
	self.dialog.Add(self.listview)
	self.dialog.Close()

	return self
}

func (self *Combobox) GetId() int32 {
	return self.label.id
}

func (self *Combobox) SetId(id int32) {
	self.label.SetId(id)
}

func (self *Combobox) GetParent() Widget {
	return self.label.GetParent()
}

func (self *Combobox) SetParent(parent Widget) {
	self.label.SetParent(parent)
}

func (self *Combobox) GetBounds() rl.Rectangle {
	return self.label.GetBounds()
}

func (self *Combobox) SetBounds(bounds rl.Rectangle) {
	self.label.SetBounds(bounds)
	b := self.dialog.GetBounds()
	b.X = bounds.X
	b.Y = bounds.Y - b.Height - bounds.Height
	b.Width = bounds.Width
	self.dialog.SetBounds(b)
}

func (self *Combobox) GetAlign() Align {
	return self.label.GetAlign()
}

func (self *Combobox) GetFill() Fill {
	return self.label.GetFill()
}

func (self *Combobox) GetDataSize() Size {
	return self.label.GetDataSize()
}

func (self *Combobox) Update(dt float32) {
	bounds := self.GetBounds()

	state := GetState(bounds)

	if self.dialog.IsActive() {
		self.dialog.Update(dt)
	} else if Pressed == self.state && Released == state {
		self.dialog.Open()
		self.Update(0)
	} else {
		self.label.Update(dt)
	}

	self.state = state
}

func (self *Combobox) Draw() {
	if self.dialog.IsActive() {
		self.dialog.Draw()
	}
	self.label.Draw()

}
