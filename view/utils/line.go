package utils

type Line struct {
	StartX, StartY, EndX, EndY float32
}

func NewLine(startX, startY, endX, endY int) Line {
	return Line{
		StartX: float32(startX),
		StartY: float32(startY),
		EndX:   float32(endX),
		EndY:   float32(endY),
	}
}
