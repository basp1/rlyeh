package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ListView struct {
	layout   *HBox
	onSelect func(string)

	items []string

	count   int
	upper   int
	current int
}

func NewListView(items []string, count int, onSelect func(string)) *ListView {
	self := &ListView{}

	self.layout = NewHBox(Auto, Both)
	self.onSelect = onSelect

	self.items = items

	self.count = count
	self.upper = 0
	self.current = 0

	vbox := NewVBox(Auto, None)

	for i := 0; i < count; i++ {
		vbox.Add(NewLabel(Left, None, ""))
	}

	self.layout.Add(vbox)

	if len(items) > count {
		scroll := NewVBox(Auto, Vertical)
		scroll.Add(NewButton(Top, None, "-", func() {
			if self.upper > 0 {
				self.upper -= 1
				self.current -= 1
			}
		}))
		scroll.Add(NewLabel(Auto, Vertical, ""))
		scroll.Add(NewButton(Top, None, "+", func() {
			if self.upper < (len(self.items) - self.count) {
				self.current += 1
				self.upper += 1
			}
		}))
		self.layout.Add(scroll)
	}

	return self
}

func (self *ListView) GetSelected() string {
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

	for i := 0; i < self.count; i++ {
		label := labels[i].(*Label)

		j := self.upper + i
		if j < len(self.items) {
			switch GetState(label.GetBounds()) {
			case Released:
				self.current = j
			}
		}
	}

	for i := 0; i < self.count; i++ {
		label := labels[i].(*Label)

		text := ""

		j := self.upper + i
		if j < len(self.items) {
			text = self.items[j]

			if j == self.current {
				self.current = j
				label.BackgroundColor = style.ListviewSelectedBackgroundColor
				label.TextColor = style.ListviewSelectedTextColor

				if nil != self.onSelect {
					self.onSelect(text)
				}
			} else {
				label.BackgroundColor = style.GlobalBackgroundColor
				label.TextColor = style.ListviewTextColor
			}
		}

		label.Text = text
	}

	self.layout.Update(dt)
}

func (self *ListView) Draw() {
	self.layout.Draw()
}
