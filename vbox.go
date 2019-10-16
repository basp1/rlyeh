package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type VBox struct {
	id     int32
	Parent Widget

	Bounds rl.Rectangle
	Align  Align
	Fill   Fill

	Widgets []Widget
}

func NewVBox(align Align, fill Fill, widgets ...Widget) *VBox {
	self := &VBox{}

	self.Align = align
	self.Fill = fill

	self.Widgets = []Widget{}
	for _, widget := range widgets {
		self.Add(widget)
	}

	return self
}

func (self *VBox) GetId() int32 {
	return self.id
}

func (self *VBox) SetId(id int32) {
	self.id = id
}

func (self *VBox) GetParent() Widget {
	return self.Parent
}

func (self *VBox) SetParent(parent Widget) {
	self.Parent = parent
}

func (self *VBox) GetBounds() rl.Rectangle {
	return self.Bounds
}

func (self *VBox) SetBounds(bounds rl.Rectangle) {
	self.Bounds = bounds
}

func (self *VBox) GetAlign() Align {
	return self.Align
}

func (self *VBox) GetFill() Fill {
	return self.Fill
}

func (self *VBox) Add(widget Widget) {
	if 0 == widget.GetId() {
		widget.SetId(nextId())
	}
	widget.SetParent(self)

	if nil == self.Widgets {
		self.Widgets = []Widget{}
	}

	self.Widgets = append(self.Widgets, widget)
}

func (self *VBox) GetDataSize() Size {
	var size Size

	for i := 0; i < len(self.Widgets); i++ {
		widget := self.Widgets[i]

		widgetSize := widget.GetDataSize()

		if widgetSize.Width > size.Width {
			size.Width = widgetSize.Width
		}
		size.Height += widgetSize.Height
	}

	size.Height += float32(int(style[GlobalPadding]) * (1 + len(self.Widgets)))

	return size
}

func (self *VBox) Update(dt float32) {
	for i := 0; i < len(self.Widgets); i++ {
		widget := self.Widgets[i]
		widget.Update(dt)
	}
}

func (self *VBox) Draw() {
	y := self.Bounds.Y + float32(style[GlobalPadding])
	bounds := self.Bounds

	fillers := 0
	for i := 0; i < len(self.Widgets); i++ {
		widget := self.Widgets[i]

		if Vertical == widget.GetFill() || Both == widget.GetFill() {
			fillers += 1
		}
	}

	free := float32(0)
	if fillers > 0 {
		free = (bounds.Height - self.GetDataSize().Height) / float32(fillers)
	}

	for i := 0; i < len(self.Widgets); i++ {
		widget := self.Widgets[i]

		dataSize := widget.GetDataSize()
		oldBounds := widget.GetBounds()

		newBounds := rl.Rectangle{X: bounds.X, Y: y, Width: dataSize.Width, Height: dataSize.Height}
		if Both == widget.GetFill() || Horizontal == widget.GetFill() {
			newBounds.X = bounds.X
			newBounds.Width = bounds.Width
		}
		if Both == widget.GetFill() || Vertical == widget.GetFill() {
			newBounds.Y = y
			newBounds.Height = dataSize.Height + free
		}

		align := self.Align
		if Auto != widget.GetAlign() {
			align = widget.GetAlign()
		}
		if Left == align || Right == align || Center == align {
			newBounds = alignBounds(bounds, newBounds, widget.GetAlign())
		}

		newBounds.Y = y
		newBounds = shrinkBounds(bounds, newBounds)

		if oldBounds != newBounds {
			widget.SetBounds(newBounds)
		}

		widget.Draw()

		y += newBounds.Height + float32(style[GlobalPadding])
	}
}

func (self *VBox) GetWidgets() []Widget {
	return self.Widgets
}
