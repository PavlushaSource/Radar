package utils

type Line struct {
	StartX, StartY, EndX, EndY float64
}

func NewLine(WindowX, WindowY, startX, startY, endX, endY int) Line {
	return Line{
		StartX: float64(getBetweenBoundary(0, WindowX, startX)),
		StartY: float64(getBetweenBoundary(0, WindowY, startY)),
		EndX:   float64(getBetweenBoundary(0, WindowX, endX)),
		EndY:   float64(getBetweenBoundary(0, WindowY, endY)),
	}
}

func getBetweenBoundary(a, b, x int) int {
	return min(max(a, x), b)
}
