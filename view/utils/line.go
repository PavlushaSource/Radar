package utils

type Line struct {
	StartX, StartY, EndX, EndY float64
}

func NewLine(startX, startY, endX, endY int) Line {
	return Line{
		StartX: float64(startX),
		StartY: float64(startY),
		EndX:   float64(endX),
		EndY:   float64(endY),
	}
}
