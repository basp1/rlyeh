package rlyeh

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ListView struct {
	OnClick       func(string)
	OnDoubleClick func(string)

	layout *HBox

	items []string

	count   int
	upper   int
	current int

	clicked int
	state   State
	time    time.Duration
}

func NewListView(items []string, count int) *ListView {
	self := &ListView{}

	self.layout = NewHBox(Auto, Both)

	self.count = count

	vbox := NewVBox(Auto, Both)

	for i := 0; i < count; i++ {
		vbox.Add(NewLabel(Left, None, ""))
	}

	self.layout.Add(vbox)

	scroll := NewVBox(Right, Vertical)
	scroll.Add(NewButton(Top, None, "-", func() {
		if self.upper > 0 {
			self.upper -= 1
			self.current -= 1
		} else if self.current > 0 {
			self.current -= 1
		}
	}))
	scroll.Add(NewLabel(Auto, Vertical, ""))
	scroll.Add(NewButton(Top, None, "+", func() {
		if self.upper < (len(self.items) - self.count) {
			self.current += 1
			self.upper += 1
		} else if self.current < (len(self.items) - 1) {
			self.current += 1
		}
	}))
	self.layout.Add(scroll)

	self.SetItems(items)

	return self
}

func (self *ListView) SetItems(items []string) {
	self.upper = 0
	self.current = 0
	self.state = Normal
	self.clicked = -1
	self.items = items

	labels := self.layout.GetWidgets()[0].(Layout).GetWidgets()
	for i := len(items); i < self.count; i++ {
		label := labels[i].(*Label)
		label.Text = ""
	}
}

func (self *ListView) GetCurrent() string {
	return self.items[self.current]
}

func (self *ListView) GetId() int32 {
	return self.layout.GetId()
}

func (self *ListView) SetId(id int32) {
	self.layout.SetId(id)
}

func (self *ListView) GetParent() Widget {
	return self.layout.GetParent()
}

func (self *ListView) SetParent(parent Widget) {
	self.layout.SetParent(parent)
}

func (self *ListView) GetBounds() rl.Rectangle {
	return self.layout.GetBounds()
}

func (self *ListView) SetBounds(bounds rl.Rectangle) {
	self.layout.SetBounds(bounds)
}

func (self *ListView) GetAlign() Align {
	return self.layout.GetAlign()
}

func (self *ListView) GetFill() Fill {
	return self.layout.GetFill()
}

func (self *ListView) GetDataSize() Size {
	return self.layout.GetDataSize()
}

func (self *ListView) Update(dt float32) {
	vbox := self.layout.GetWidgets()[0].(*VBox)
	labels := vbox.GetWidgets()

	self.time += time.Duration(1e9 * dt)

	for i := 0; i < self.count; i++ {
		label := labels[i].(*Label)

		j := self.upper + i
		if j < len(self.items) {
			text := self.items[j]
			label.Text = text

			switch GetState(label.GetBounds()) {
			case Released:

				if Pressed == self.state && j == self.current {
					self.time = 0
					self.clicked = j
					if nil != self.OnClick {
						self.OnClick(text)
					}
				}

				self.current = j
				self.state = Released
				break
			case Pressed:
				if j == self.clicked && j == self.current {
					if (200*time.Millisecond) > self.time && nil != self.OnDoubleClick {
						self.OnDoubleClick(text)
						self.clicked = -1
					}
				}

				self.current = j
				self.state = Pressed
				break
			}

			if self.current == j {
				label.BackgroundColor = style.ListviewSelectedBackgroundColor
				label.TextColor = style.ListviewSelectedTextColor
			} else {
				label.BackgroundColor = style.GlobalBackgroundColor
				label.TextColor = style.ListviewTextColor
			}
		}
	}

	self.layout.Update(dt)
}

func (self *ListView) Draw() {
	self.layout.Draw()
}
