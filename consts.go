package rlyeh

type Align int32

const (
	Auto Align = iota
	Top
	Bottom
	Left
	Right
	Center
)

type State int32

const (
	Normal State = iota
	Focused
	Pressed
	Released
	Disabled
)

type Fill int32

const (
	None Fill = iota
	Vertical
	Horizontal
	Both
)
