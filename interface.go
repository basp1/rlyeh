package rlyeh

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Widget interface {
	GetId() int32
	SetId(int32)

	GetParent() Widget
	SetParent(Widget)

	GetBounds() rl.Rectangle
	SetBounds(rl.Rectangle)
	GetDataSize() Size

	GetAlign() Align
	GetFill() Fill

	Update(float32)
	Draw()
}

type Form interface {
	Update(float32)
	Draw()

	Add(Widget)

	IsActive() bool
	SetActive(bool)

	IsModal() bool
	IsMovable() bool
}

type Layout interface {
	Widget

	Add(Widget)
	GetWidgets() []Widget
}
