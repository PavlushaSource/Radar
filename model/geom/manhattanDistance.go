package geom

import (
	"math"
)

// ManhattanDistance returns a Manhattan distance between to points
func ManhattanDistance(first Point, second Point, barriers []Barrier) float64 {
	// there are to ways to get achievable
	if !arePointsAchievable(first, second, barriers, minXManhattanAchievability) &&
		!arePointsAchievable(first, second, barriers, maxXManhattanAchievability) {
		return InfDistance
	}

	return math.Abs(first.X()-second.X()) + math.Abs(first.Y()-second.Y())
}

// minXManhattanAchievability returns true if lower (by X coord) Manhattan path between two points is achievability
func minXManhattanAchievability(first Point, second Point, barrier Barrier) bool {
	return math.Min(first.X(), second.X()) < math.Min(barrier.StartPoint().X(), barrier.FinishPoint().X()) &&
		math.Max(first.Y(), second.Y()) > math.Max(barrier.StartPoint().Y(), barrier.FinishPoint().Y())
}

// minXManhattanAchievability returns true if upper (by X coord) Manhattan path between two points is achievability
func maxXManhattanAchievability(first Point, second Point, barrier Barrier) bool {
	return math.Max(first.X(), second.X()) > math.Max(barrier.StartPoint().X(), barrier.FinishPoint().X()) &&
		math.Min(first.Y(), second.Y()) < math.Min(barrier.StartPoint().Y(), barrier.FinishPoint().Y())
}
