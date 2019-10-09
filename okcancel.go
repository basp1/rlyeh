package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func OkCancel(text string, onOk func()) *Dialog {
	fontsize := int32(style[GlobalTextFontsize])
	width := float32(rl.MeasureText(text, fontsize))

	dialog := NewDialog(rl.NewRectangle(2*width, float32(5*fontsize), 0, 0))

	vbox := NewVBox(Auto, None)
	vbox.Add(NewLabel(Center, None, text))

	hbox := NewHBox(Center, None)
	hbox.Add(NewButton(Auto, None, "Ok", func() {
		onOk()
		dialog.Close()
	}))
	hbox.Add(NewButton(Auto, None, "Cancel", func() {
		dialog.Close()
	}))

	vbox.Add(hbox)

	dialog.Add(vbox)

	return dialog
}
