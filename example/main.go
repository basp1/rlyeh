package main

import (
	"math/rand"

	"github.com/basp1/rlyeh"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

func main() {
	app := rlyeh.NewApplication(SCREEN_WIDTH, SCREEN_HEIGHT, "Example")

	dialog := rlyeh.OkCancel("Message", func() {
		rlyeh.Notify(2, "Ok pressed...")
	})

	app.Add(dialog)
	app.Add(NewWindow(dialog))

	app.Run()

	app.Close()
}

func NewWindow(dialog *rlyeh.Dialog) *rlyeh.Window {
	window := rlyeh.NewWindow(rl.NewRectangle(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT))

	vbox := rlyeh.NewVBox(rlyeh.Auto, rlyeh.Both)
	vbox.Add(rlyeh.NewLabel(rlyeh.Auto, rlyeh.Both, "Label"))

	hbox := rlyeh.NewHBox(rlyeh.Center, rlyeh.None)
	hbox.Add(rlyeh.NewButton(rlyeh.Auto, rlyeh.None, "Open dialog", func() {
		dialog.Open()
	}))

	hbox.Add(rlyeh.NewButton(rlyeh.Auto, rlyeh.None, "Notify", func() {
		rlyeh.Notify(5, "Hello %s!", RandStringRunes(rand.Int()%10))
	}))

	hbox.Add(rlyeh.NewCheckbox(rlyeh.Auto, rlyeh.None))
	hbox.Add(rlyeh.NewLabel(rlyeh.Auto, rlyeh.None, "check"))

	hbox.Add(rlyeh.NewCombobox(rlyeh.Auto, rlyeh.None, []string{"one", "two", "three"}))
	hbox.Add(rlyeh.NewTextbox(rlyeh.Auto, rlyeh.None, 8))

	vbox.Add(hbox)

	window.Add(vbox)

	return window
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
