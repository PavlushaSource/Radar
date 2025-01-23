package utils

type Cursor struct {
	StartX, StartY int
	EndX, EndY     int
	Pressed        bool
}

func NewCursor() Cursor {
	return Cursor{}
}

func (c *Cursor) Reset() {
	c.StartX, c.StartY = 0, 0
	c.EndX, c.EndY = 0, 0
	c.Pressed = false
}
