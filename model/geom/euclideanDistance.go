package geom

import "math"

// EuclideanDistance returns a Euclidean distance between to points
func EuclideanDistance(first Point, second Point, barriers []Barrier) float64 {
	if !arePointsAchievable(first, second, barriers, euclideanAchievability) {
		return InfDistance
	}

	xDiff := first.X() - second.X()
	yDiff := first.Y() - second.Y()

	return math.Sqrt(xDiff*xDiff + yDiff*yDiff)
}

// euclideanAchievability returns true if the segment between the points does not intersect with the barrier
func euclideanAchievability(first Point, second Point, barrier Barrier) bool {
	// checks projections intersection
	var projectionAchievability = func(ls1Start, ls1Finish, ls2Start, ls2Finish float64) bool {
		if ls1Start > ls1Finish {
			ls1Start, ls1Finish = ls1Finish, ls1Start
		}

		if ls2Start > ls2Finish {
			ls2Start, ls2Finish = ls2Finish, ls2Start
		}

		return max(ls1Start, ls2Start) <= min(ls1Finish, ls2Finish)
	}

	// returns triangle area
	var getTriangleArea = func(p1 Point, p2 Point, p3 Point) float64 {
		return 0.5 * ((p2.X()-p1.X())*(p3.Y()-p1.Y()) - (p2.Y()-p1.Y())*(p3.X()-p1.X()))
	}

	return !(projectionAchievability(first.X(), second.X(), barrier.StartPoint().X(), barrier.FinishPoint().X()) &&
		projectionAchievability(first.Y(), second.Y(), barrier.StartPoint().Y(), barrier.FinishPoint().Y()) &&
		getTriangleArea(first, second, barrier.StartPoint())*getTriangleArea(first, second, barrier.FinishPoint()) <= Eps &&
		getTriangleArea(barrier.StartPoint(), barrier.FinishPoint(), first)*getTriangleArea(barrier.StartPoint(), barrier.FinishPoint(), second) <= Eps)
}
