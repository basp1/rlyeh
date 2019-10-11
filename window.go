package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	id int32

	active bool

	bounds  rl.Rectangle
	widgets []Widget
}

func NewWindow(bounds rl.Rectangle, widgets ...Widget) *Window {
	self := &Window{}

	self.id = nextId()
	self.active = true
	self.bounds = bounds

	self.widgets = []Widget{}
	for _, widget := range widgets {
		self.Add(widget)
	}

	return self
}

func (self *Window) GetBounds() rl.Rectangle {
	return self.bounds
}

func (self *Window) SetBounds(bounds rl.Rectangle) {
	self.bounds = bounds
}

func (self *Window) Add(widget Widget) {
	if 0 == widget.GetId() {
		widget.SetId(nextId())
	}

	self.widgets = append(self.widgets, widget)
}

func (self *Window) Update(dt float32) {
	if !self.IsActive() {
		return
	}

	for i := 0; i < len(self.widgets); i++ {
		widget := self.widgets[i]

		widget.Update(dt)
	}
}

func (self *Window) Draw() {
	if !self.IsActive() {
		return
	}

	for i := 0; i < len(self.widgets); i++ {
		widget := self.widgets[i]
		oldBounds := widget.GetBounds()
		dataSize := widget.GetDataSize()

		newBounds := rl.Rectangle{X: self.bounds.X, Y: self.bounds.Y,
			Width: dataSize.Width, Height: dataSize.Height}
		newBounds = fillBounds(self.bounds, newBounds, widget.GetFill())
		newBounds = alignBounds(self.bounds, newBounds, widget.GetAlign())
		newBounds = shrinkBounds(self.bounds, newBounds)

		if oldBounds != newBounds {
			widget.SetBounds(newBounds)
		}

		widget.Draw()
	}
}

func (self *Window) IsActive() bool {
	return self.active
}

func (self *Window) SetActive(value bool) {
	self.active = value
}
