package rlyeh

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewSaveFileDialog(path string, callback func(item string)) *Dialog {
	items := getFileList(path)

	dialog := NewDialog(rl.Rectangle{100, 100, 0, 0}, "Save file...")

	textbox := NewTextbox(Auto, Horizontal, 20)

	vbox := NewVBox(Auto, Both)

	listview := NewListView(items, 10)
	listview.OnClick = func(item string) {
		textbox.Text = item
	}
	listview.OnDoubleClick = func(item string) {
		name := path + "/" + item
		fi, err := os.Stat(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			path = name
			listview.SetItems(getFileList(path))
			break
		}
	}

	vbox.Add(listview)

	vbox.Add(textbox)

	hbox := NewHBox(Right, None)
	hbox.Add(NewButton(Auto, None, "Save", func() {
		if "" != textbox.Text && nil != callback {
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
	items := getFileList(path)

	dialog := NewDialog(rl.Rectangle{100, 100, 0, 0}, "Open file...")

	vbox := NewVBox(Auto, Both)

	listview := NewListView(items, 10)
	listview.OnDoubleClick = func(item string) {
		name := path + item
		fi, err := os.Stat(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			path = name
			listview.SetItems(getFileList(path))
			break
		}
	}

	vbox.Add(listview)

	hbox := NewHBox(Right, None)
	hbox.Add(NewButton(Auto, None, "Open", func() {
		item := listview.GetCurrentItem()
		if nil != callback {
			callback(item)
		}
		dialog.Close()
	}))
	hbox.Add(NewButton(Auto, None, "Cancel", func() {
		dialog.Close()
	}))

	vbox.Add(hbox)

	dialog.Add(vbox)

	return dialog
}

func getFileList(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	items := []string{"../"}

	for _, f := range files {
		if f.IsDir() {
			items = append(items, f.Name()+"/")
		}
	}

	for _, f := range files {
		if !f.IsDir() {
			items = append(items, f.Name())
		}
	}

	return items
}
