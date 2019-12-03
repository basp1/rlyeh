package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type HBox struct {
	id     int32
	Parent Widget

	Bounds rl.Rectangle
	Align  Align
	Fill   Fill

	Widgets []Widget
}

func NewHBox(align Align, fill Fill, widgets ...Widget) *HBox {
	self := &HBox{}

	self.Align = align
	self.Fill = fill

	self.Widgets = []Widget{}
	for _, widget := range widgets {
		self.Add(widget)
	}

	return self
}

func (self *HBox) GetId() int32 {
	return self.id
}

func (self *HBox) SetId(id int32) {
	self.id = id
}

func (self *HBox) GetParent() Widget {
	return self.Parent
}

func (self *HBox) SetParent(parent Widget) {
	self.Parent = parent
}

func (self *HBox) GetBounds() rl.Rectangle {
	return self.Bounds
}

func (self *HBox) SetBounds(bounds rl.Rectangle) {
	self.Bounds = bounds
}

func (self *HBox) GetAlign() Align {
	return self.Align
}

func (self *HBox) GetFill() Fill {
	return self.Fill
}

func (self *HBox) Add(widget Widget) {
	if 0 == widget.GetId() {
		widget.SetId(nextId())
	}
	widget.SetParent(self)

	if nil == self.Widgets {
		self.Widgets = []Widget{}
	}

	self.Widgets = append(self.Widgets, widget)
}

func (self *HBox) GetDataSize() Size {
	var size Size

	for i := 0; i < len(self.Widgets); i++ {
		widget := self.Widgets[i]

		widgetSize := widget.GetDataSize()

		size.Width += widgetSize.Width
		if widgetSize.Height > size.Height {
			size.Height = widgetSize.Height
		}
	}

	size.Width += float32(int(style.GlobalPadding) * (1 + len(self.Widgets)))

	return size
}

func (self *HBox) Update(dt float32) {
	for i := 0; i < len(self.Widgets); i++ {
		widget := self.Widgets[i]
		widget.Update(dt)
	}

	x := self.Bounds.X + float32(style.GlobalPadding)
	bounds := self.Bounds

	fillers := 0
	for i := 0; i < len(self.Widgets); i++ {
		widget := self.Widgets[i]

		if Horizontal == widget.GetFill() || Both == widget.GetFill() {
			fillers += 1
		}
	}

	free := float32(0)
	if fillers > 0 {
		free = (bounds.Width - self.GetDataSize().Width) / float32(fillers)
	}

	for i := 0; i < len(self.Widgets); i++ {
		widget := self.Widgets[i]

		dataSize := widget.GetDataSize()
		oldBounds := widget.GetBounds()

		newBounds := rl.Rectangle{X: x, Y: bounds.Y, Width: dataSize.Width, Height: dataSize.Height}
		if Both == widget.GetFill() || Vertical == widget.GetFill() {
			newBounds.Y = bounds.Y
			newBounds.Height = bounds.Height
		}
		if Both == widget.GetFill() || Horizontal == widget.GetFill() {
			newBounds.X = x
			newBounds.Width = dataSize.Width + free
		}

		align := self.Align
		if Auto != widget.GetAlign() {
			align = widget.GetAlign()
		}
		if Top == align || Bottom == align || Center == align {
			newBounds = alignBounds(bounds, newBounds, widget.GetAlign())
		}

		newBounds.X = x
		newBounds = shrinkBounds(bounds, newBounds)

		if oldBounds != newBounds {
			widget.SetBounds(newBounds)
		}

		x += newBounds.Width + float32(style.GlobalPadding)
	}
}

func (self *HBox) Draw() {
	for i := 0; i < len(self.Widgets); i++ {
		widget := self.Widgets[i]

		widget.Draw()
	}
}

func (self *HBox) GetWidgets() []Widget {
	return self.Widgets
}

func (self *HBox) Clear() {
	self.Widgets = []Widget{}
}
