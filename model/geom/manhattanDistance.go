package geom

import (
	"math"
)

func ManhattanDistance(first Point, second Point, barriers []Barrier) float64 {
	// there are to ways to get achievable
	if !arePointsAchievable(first, second, barriers, minXManhattanAchievability) &&
		!arePointsAchievable(first, second, barriers, maxXManhattanAchievability) {
		return InfDistance
	}

	return math.Abs(first.X()-second.X()) + math.Abs(first.Y()-second.Y())
}

func minXManhattanAchievability(first Point, second Point, barrier Barrier) bool {
	var barrierLine = barrier.LineSegment

	return math.Min(first.X(), second.X()) < math.Min(barrierLine.StartPoint.X(), barrierLine.FinishPoint.X()) &&
		math.Max(first.Y(), second.Y()) > math.Max(barrierLine.StartPoint.Y(), barrierLine.FinishPoint.Y())
}

func maxXManhattanAchievability(first Point, second Point, barrier Barrier) bool {
	var barrierLine = barrier.LineSegment

	return math.Max(first.X(), second.X()) > math.Max(barrierLine.StartPoint.X(), barrierLine.FinishPoint.X()) &&
		math.Min(first.Y(), second.Y()) < math.Min(barrierLine.StartPoint.Y(), barrierLine.FinishPoint.Y())
}
