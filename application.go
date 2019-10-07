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

	dialogs []*Dialog
	windows []Container
	modal   *Dialog
}

func NewApplication(width, height int, title string) *Application {
	self := &Application{}

	self.Width = width
	self.Height = height
	self.Title = title

	self.dialogs = make([]*Dialog, 0)
	self.windows = make([]Container, 0)
	self.modal = nil

	rand.Seed(time.Now().UTC().UnixNano())
	rl.InitWindow(int32(width), int32(height), title)
	rl.SetTargetFPS(60)

	application = self

	return self
}

func (self *Application) Add(container Container) {
	switch obj := container.(type) {
	case *Dialog:
		self.dialogs = append(self.dialogs, obj)
		break
	default:
		self.windows = append(self.windows, obj)
		break
	}
}

func (self *Application) Run() {
	for _, dialog := range self.dialogs {
		dialog.Close()
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

func (self *Application) update(dt float32) {
	for _, dialog := range self.dialogs {
		if dialog.Modal && dialog.IsOpen() {
			self.modal = dialog
		}
	}

	if nil != self.modal && self.modal.IsOpen() {
		self.modal.Update(dt)
	} else {
		for _, window := range self.windows {
			window.Update(dt)
		}
		for _, dialog := range self.dialogs {
			dialog.Update(dt)
		}
	}

	height := float32(self.Height)
	filtered := make([]*Notification, 0)
	for i := 0; i < len(notifications); i++ {
		notifications[i].Seconds -= dt
		if notifications[i].Seconds > 0 {
			height -= float32(style[GlobalTextFontsize])
			notifications[i].Point.Y = height
			filtered = append(filtered, notifications[i])
		}
	}
	notifications = filtered
}

func (self *Application) draw() {
	for _, window := range self.windows {
		window.Draw()
	}
	for _, dialog := range self.dialogs {
		if self.modal != dialog {
			dialog.Draw()
		}
	}

	for i := 0; i < len(notifications); i++ {
		rl.DrawText(notifications[i].Text,
			int32(notifications[i].Point.X), int32(notifications[i].Point.Y),
			notifications[i].FontSize, notifications[i].Color)
	}

	if nil != self.modal {
		self.modal.Draw()
	}
}
