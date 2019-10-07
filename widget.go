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

type Container interface {
	Update(float32)
	Draw()
	Add(Widget)
}

type Size struct {
	Width  float32
	Height float32
}

func getState(bounds rl.Rectangle) State {
	state := Normal

	mousePoint := rl.GetMousePosition()

	if rl.CheckCollisionPointRec(mousePoint, bounds) {
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			state = Pressed
		} else if rl.IsMouseButtonReleased(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			state = Released
		} else {
			state = Focused
		}
	}

	return state
}

func shrinkBounds(parent rl.Rectangle, child rl.Rectangle) rl.Rectangle {
	xMax, yMax := parent.X+parent.Width, parent.Y+parent.Height

	child.X = max(parent.X, min(xMax, child.X))
	child.Y = max(parent.Y, min(yMax, child.Y))
	child.Width = min(xMax-child.X, child.Width)
	child.Height = min(yMax-child.Y, child.Height)

	return child
}

func alignBounds(parent rl.Rectangle, child rl.Rectangle, align Align) rl.Rectangle {
	xMax, yMax := parent.X+parent.Width, parent.Y+parent.Height

	switch align {
	case Left:
		child.X = parent.X
	case Top:
		child.Y = parent.Y
	case Right:
		child.X = xMax - child.Width
	case Bottom:
		child.Y = yMax - child.Height
	case Center:
		child.X = parent.X + (xMax-parent.X)/2 - child.Width/2
		child.Y = parent.Y + (yMax-parent.Y)/2 - child.Height/2
	}

	return child
}

func fillBounds(parent rl.Rectangle, child rl.Rectangle, fill Fill) rl.Rectangle {
	switch fill {
	case Horizontal:
		child.X = parent.X
		child.Width = parent.Width
	case Vertical:
		child.Y = parent.Y
		child.Height = parent.Height
	case Both:
		child.X = parent.X
		child.Y = parent.Y
		child.Width = parent.Width
		child.Height = parent.Height
	}

	return child
}
