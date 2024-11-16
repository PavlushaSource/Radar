package geom

import (
	"math"
)

type Distance func(first Point, second Point) float64

func EuclideanDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.X()-p2.X(), 2) + math.Pow(p1.Y()-p2.Y(), 2))
}

func ManhattanDistance(p1 Point, p2 Point) float64 {
	return math.Abs(p1.X()-p2.X()) + math.Abs(p1.Y()-p2.Y())
}

func CurvilinearDistance(p1 Point, p2 Point) float64 {
	// We need to choose a function to implement it.
	return 0
}
