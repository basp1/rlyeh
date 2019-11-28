package rlyeh

import (
	"io/ioutil"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewSaveFileDialog(path string, callback func(item string)) *Dialog {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	items := []string{}
	for _, f := range files {
		items = append(items, f.Name())
	}

	dialog := NewDialog(rl.Rectangle{100, 100, 0, 0})

	textbox := NewTextbox(Auto, Horizontal, 20)

	vbox := NewVBox(Auto, Both)

	listview := NewListView(items, 10)
	listview.OnClick = func(item string) {
		textbox.Text = item
	}

	vbox.Add(listview)

	vbox.Add(textbox)

	hbox := NewHBox(Right, None)
	hbox.Add(NewButton(Auto, None, "Save", func() {
		if "" != textbox.Text {
			callback(textbox.Text)
		}
	}))
	hbox.Add(NewButton(Auto, None, "Cancel", func() {
		dialog.Close()
	}))

	vbox.Add(hbox)

	dialog.Add(vbox)

	return dialog
}

func NewOpenFileDialog(path string, callback func(item string)) *Dialog {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	items := []string{}
	for _, f := range files {
		items = append(items, f.Name())
	}

	dialog := NewDialog(rl.Rectangle{100, 100, 0, 0})

	vbox := NewVBox(Auto, Both)

	listview := NewListView(items, 10)

	vbox.Add(listview)

	hbox := NewHBox(Right, None)
	hbox.Add(NewButton(Auto, None, "Open", func() {
		item := listview.GetCurrent()
		callback(item)
		dialog.Close()
	}))
	hbox.Add(NewButton(Auto, None, "Cancel", func() {
		dialog.Close()
	}))

	vbox.Add(hbox)

	dialog.Add(vbox)

	return dialog
}
