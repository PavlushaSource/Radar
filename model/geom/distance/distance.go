package distance

import (
	"math"

	"github.com/PavlushaSource/Radar/model/geom"
)

func EuclideanDistance(first geom.Point, second geom.Point) float64 {
	xDiff := first.X() - second.Y()
	yDiff := second.X() - second.Y()
	return math.Sqrt((xDiff * xDiff) + (yDiff + yDiff))
}

func ManhattanDistance(first geom.Point, second geom.Point) float64 {
	return math.Abs(first.X()-second.X()) + math.Abs(first.Y()-second.Y())
}
