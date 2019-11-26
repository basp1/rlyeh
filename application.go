package rlyeh

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	application *Application = nil
)

func GetApplication() *Application {
	return application
}

type Application struct {
	Width  int
	Height int
	Title  string

	movables []Form
	forms    []Form
	modal    Form

	options map[string]interface{}
}

func NewApplication(width, height int, title string) *Application {
	self := &Application{}

	self.Width = width
	self.Height = height
	self.Title = title

	self.movables = make([]Form, 0)
	self.forms = make([]Form, 0)
	self.modal = nil

	self.options = make(map[string]interface{})

	rand.Seed(time.Now().UTC().UnixNano())

	rl.InitWindow(int32(width), int32(height), title)
	rl.SetTargetFPS(60)
	rl.SetExitKey(-1)

	application = self

	return self
}

func (self *Application) GetOption(name string) interface{} {
	value, ok := self.options[name]
	if !ok {
		return nil
	} else {
		return value
	}
}

func (self *Application) SetOption(name string, value interface{}) {
	self.options[name] = value
}

func (self *Application) RemoveOption(name string) {
	delete(self.options, name)
}

func (self *Application) Add(form Form) {
	if form.IsMovable() {
		self.movables = append(self.movables, form)
	} else {
		self.forms = append(self.forms, form)
	}
}

func findForm(forms []Form, form Form) int {
	index := -1

	for i, x := range forms {
		if x == form {
			index = i
			break
		}
	}

	return index
}

func deleteForm(forms []Form, index int) []Form {
	if index < len(forms)-1 {
		copy(forms[index:], forms[index+1:])
	}
	forms[len(forms)-1] = nil
	forms = forms[:len(forms)-1]

	return forms
}

func (self *Application) Remove(form Form) {
	if form.IsMovable() {
		index := findForm(self.movables, form)
		if index >= 0 {
			self.movables = deleteForm(self.movables, index)
		}
	} else {
		index := findForm(self.forms, form)
		if index >= 0 {
			self.forms = deleteForm(self.forms, index)
		}
	}
}

func (self *Application) Run() {
	for _, movable := range self.movables {
		movable.SetActive(false)
	}

	frameTime := time.Now().UTC().UnixNano()

	for !rl.WindowShouldClose() {
		var dt float32

		time := time.Now().UTC().UnixNano()
		if 0 == frameTime {
			dt = 0
		} else {
			dt = float32(time-frameTime) / 1e9
		}
		frameTime = time

		self.update(dt)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		self.draw()

		rl.EndDrawing()
	}
}

func (self *Application) Close() {
	rl.CloseWindow()
}

func (self *Application) GetStyle() *Style {
	return style
}

func (self *Application) update(dt float32) {
	for _, movable := range self.movables {
		if movable.IsModal() && movable.IsActive() {
			self.modal = movable
		}
	}

	if nil != self.modal && self.modal.IsActive() {
		self.modal.Update(dt)
	} else {
		for _, form := range self.forms {
			if form.IsActive() {
				form.Update(dt)
			}
		}
		for _, movable := range self.movables {
			if movable.IsActive() {
				movable.Update(dt)
			}
		}
	}

	height := float32(self.Height)
	filtered := make([]*Notification, 0)
	for i := 0; i < len(notifications); i++ {
		notifications[i].Duration -= time.Nanosecond * time.Duration(1e9*dt)
		if notifications[i].Duration > 0 {
			height -= float32(style.GlobalTextFontsize)
			notifications[i].Point.Y = height
			filtered = append(filtered, notifications[i])
		}
	}
	notifications = filtered
}

func (self *Application) draw() {
	for _, form := range self.forms {
		form.Draw()
	}
	for _, movable := range self.movables {
		if self.modal != movable {
			movable.Draw()
		}
	}

	for i := 0; i < len(notifications); i++ {
		notifications[i].Draw()
	}

	if nil != self.modal {
		self.modal.Draw()
	}
}
