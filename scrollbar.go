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
		self.layout.Add(NewButton(Auto, Vertical, "=", nil))
		self.layout.Add(NewButton(Bottom, None, "+", func() { self.scrollable.Scroll(1) }))
	case Horizontal:
		self.layout = NewHBox(Bottom, Horizontal)
		self.layout.Add(NewButton(Left, None, "-", func() { self.scrollable.Scroll(-1) }))
		self.layout.Add(NewButton(Auto, Horizontal, "=", nil))
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

	bounds := self.GetBounds()
	button := self.layout.GetWidgets()[0]
	w := button.GetBounds().Width
	h := button.GetBounds().Height
	slider := self.layout.GetWidgets()[1]

	slider.SetBounds(rl.Rectangle{X: bounds.X,
		Y:      bounds.Y + h + (bounds.Height-3*h)*float32(1+self.scrollable.GetCurrent())/float32(self.scrollable.GetCount()),
		Width:  w,
		Height: h,
	})
}

func (self *Scrollbar) Draw() {
	self.layout.Draw()
}
