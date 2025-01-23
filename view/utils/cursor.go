package utils

type Cursor struct {
	StartX, StartY int
	EndX, EndY     int
	Pressed        bool
}

func NewCursor() Cursor {
	return Cursor{}
}
