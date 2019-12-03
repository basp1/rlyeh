package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scrollbar struct {
	scrollable Scrollable
	layout     Layout

	state State
}

func NewScrollbar(fill Fill, scrollable Scrollable) *Scrollbar {
	self := &Scrollbar{}

	self.scrollable = scrollable
	self.state = Normal

	switch fill {
	case Vertical:
		self.layout = NewVBox(Right, Vertical)
		self.layout.Add(NewButton(Top, None, "-", func() { self.scrollable.Scroll(-1) }))
		self.layout.Add(NewLabel(Auto, Vertical, ""))
		self.layout.Add(NewButton(Bottom, None, "+", func() { self.scrollable.Scroll(1) }))
	case Horizontal:
		self.layout = NewHBox(Bottom, Horizontal)
		self.layout.Add(NewButton(Left, None, "-", func() { self.scrollable.Scroll(-1) }))
		self.layout.Add(NewLabel(Auto, Horizontal, ""))
		self.layout.Add(NewButton(Right, None, "+", func() { self.scrollable.Scroll(1) }))
	default:
		panic("fill should be 'Vertical' or 'Horizontal'")
	}

	return self
}

func (self *Scrollbar) GetId() int32 {
	return self.layout.GetId()
}

func (self *Scrollbar) SetId(id int32) {
	self.layout.SetId(id)
}

func (self *Scrollbar) GetParent() Widget {
	return self.layout.GetParent()
}

func (self *Scrollbar) SetParent(parent Widget) {
	self.layout.SetParent(parent)
}

func (self *Scrollbar) GetBounds() rl.Rectangle {
	return self.layout.GetBounds()
}

func (self *Scrollbar) SetBounds(bounds rl.Rectangle) {
	self.layout.SetBounds(bounds)
}

func (self *Scrollbar) GetAlign() Align {
	return self.layout.GetAlign()
}

func (self *Scrollbar) GetFill() Fill {
	return self.layout.GetFill()
}

func (self *Scrollbar) GetDataSize() Size {
	return self.layout.GetDataSize()
}

func (self *Scrollbar) Update(dt float32) {
	self.layout.Update(dt)
}

func (self *Scrollbar) Draw() {
	self.layout.Draw()
}
