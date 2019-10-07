package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Dialog struct {
	id int32

	widgets []Widget
	window  *Window

	dragState State
	dragPoint rl.Vector2

	Modal bool
}

func NewDialog(bounds rl.Rectangle, widgets ...Widget) *Dialog {
	self := &Dialog{}

	self.id = nextId()
	self.Modal = true

	bounds.Y += float32(style[GlobalBorderHeight])
	self.window = NewWindow(bounds)

	self.widgets = []Widget{}
	for _, widget := range widgets {
		self.Add(widget)
	}

	return self
}

func (self *Dialog) Add(widget Widget) {
	if 0 == widget.GetId() {
		widget.SetId(nextId())
	}

	self.widgets = append(self.widgets, widget)

	self.window.Add(widget)
}

func (self *Dialog) Update(dt float32) {
	if !self.window.Active {
		return
	}

	pressed := rl.IsMouseButtonDown(rl.MouseLeftButton)

	borderBounds := self.window.GetBounds()
	borderBounds.Y -= float32(style[GlobalBorderHeight])
	borderBounds.Height = float32(style[GlobalBorderHeight])

	borderState := getState(borderBounds)

	if !pressed {
		self.dragState = Normal
	} else if Pressed == borderState && Pressed != self.dragState {
		self.dragState = Pressed
		self.dragPoint = rl.GetMousePosition()
	} else if pressed && Pressed == self.dragState {
		bounds := self.window.GetBounds()
		mousePoint := rl.GetMousePosition()
		bounds.X += mousePoint.X - self.dragPoint.X
		bounds.Y += mousePoint.Y - self.dragPoint.Y
		self.dragPoint = mousePoint
		self.window.SetBounds(bounds)
	}

	self.window.Update(dt)
}

func (self *Dialog) GetDataSize() Size {
	var dataSize Size

	for i := 0; i < len(self.widgets); i++ {
		widget := self.widgets[i]
		widgetSize := widget.GetDataSize()
		if widgetSize.Width > dataSize.Width {
			dataSize.Width = widgetSize.Width
		}
		if widgetSize.Height > dataSize.Height {
			dataSize.Height = widgetSize.Height
		}
	}

	return dataSize
}

func (self *Dialog) Draw() {
	if !self.window.Visible {
		return
	}

	bounds := self.window.GetBounds()
	dataSize := self.GetDataSize()
	if bounds.Width < dataSize.Width {
		bounds.Width = dataSize.Width
	}
	if bounds.Height < dataSize.Height {
		bounds.Height = dataSize.Height
	}
	self.window.SetBounds(bounds)

	b := bounds.ToInt32()
	rl.DrawRectangle(b.X, b.Y-int32(style[GlobalBorderHeight]), b.Width, int32(style[GlobalBorderHeight]), GetColor(GlobalBorderColor))
	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, GetColor(GlobalBackgroundColor))
	rl.DrawRectangleLines(b.X, b.Y, b.Width, b.Height, GetColor(GlobalLinesColor))

	self.window.Draw()
}

func (self *Dialog) IsOpen() bool {
	return self.window.Active && self.window.Visible
}

func (self *Dialog) Open() {
	self.window.Active = true
	self.window.Visible = true
}

func (self *Dialog) Close() {
	self.window.Active = false
	self.window.Visible = false
}
