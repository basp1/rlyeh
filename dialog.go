package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Dialog struct {
	id int32

	window *Window

	dragState State
	dragPoint rl.Vector2

	modal bool

	Title      string
	Decoration bool
}

func NewDialog(bounds rl.Rectangle, title string) *Dialog {
	self := &Dialog{}

	self.id = nextId()
	self.modal = true

	self.Title = title
	self.Decoration = true

	bounds.Y += float32(style.DialogTitleFontsize)
	self.window = NewWindow(bounds)

	app := GetApplication()
	if nil != app {
		app.Add(self)
	}

	return self
}

func (self *Dialog) GetBounds() rl.Rectangle {
	return self.window.GetBounds()
}

func (self *Dialog) SetBounds(bounds rl.Rectangle) {
	self.window.SetBounds(bounds)
}

func (self *Dialog) Add(widget Widget) {
	if 0 == widget.GetId() {
		widget.SetId(nextId())
	}

	self.window.Add(widget)
}

func (self *Dialog) Clear() {
	self.window.widgets = self.window.widgets[:0]
}

func (self *Dialog) Update(dt float32) {
	if !self.window.IsActive() {
		return
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		self.window.SetActive(false)
		return
	}

	pressed := rl.IsMouseButtonDown(rl.MouseLeftButton)

	if self.Decoration {
		borderBounds := self.window.GetBounds()
		borderBounds.Y -= float32(style.DialogTitleFontsize)
		borderBounds.Height = float32(style.DialogTitleFontsize)

		borderState := GetState(borderBounds)

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
	}

	self.window.Update(dt)
}

func (self *Dialog) GetDataSize() Size {
	var dataSize Size

	for i := 0; i < len(self.window.widgets); i++ {
		widget := self.window.widgets[i]
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
	if !self.window.IsActive() {
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

	if self.Decoration {
		rl.DrawRectangle(b.X, b.Y-int32(style.DialogTitleFontsize), b.Width, int32(style.DialogTitleFontsize), style.DialogTitleBackgroundColor)
		rl.DrawText(self.Title, b.X+1, b.Y-int32(style.DialogTitleFontsize), int32(style.DialogTitleFontsize), style.DialogTitleTextColor)
	}
	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, style.GlobalBackgroundColor)
	rl.DrawRectangleLines(b.X, b.Y, b.Width, b.Height, style.GlobalLinesColor)

	self.window.Draw()
}

func (self *Dialog) IsActive() bool {
	return self.window.IsActive()
}

func (self *Dialog) SetActive(value bool) {
	self.window.SetActive(value)
}

func (self *Dialog) Close() {
	self.SetActive(false)
}

func (self *Dialog) Open() {
	self.SetActive(true)
}

func (self *Dialog) IsModal() bool {
	return self.modal
}

func (self *Dialog) IsMovable() bool {
	return true
}
