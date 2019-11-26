package rlyeh

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	notifications []*Notification = make([]*Notification, 0)
)

type Notification struct {
	Text     string
	Point    rl.Vector2
	FontSize int32
	Color    rl.Color
	Duration time.Duration
}

func (self *Notification) Draw() {
	rl.DrawText(self.Text, int32(self.Point.X), int32(self.Point.Y), self.FontSize, self.Color)
}

func Notify(duration time.Duration, format string, params ...interface{}) {
	app := GetApplication()
	if nil == app {
		return
	}

	text := fmt.Sprintf(format, params...)

	fontSize := int32(style.GlobalTextFontsize)
	color := style.LabelTextColor
	point := rl.Vector2{X: float32(app.Width) - float32(rl.MeasureText(text, fontSize)),
		Y: float32(app.Height) - float32(1+len(notifications))*float32(fontSize)}

	notifications = append(notifications, &Notification{text, point, fontSize, color, duration})
}
