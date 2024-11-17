package geom

import "math"

func EuclideanDistance(first Point, second Point, barriers []Barrier) float64 {
	if !arePointsAchievable(first, second, barriers, euclideanAchievability) {
		return InfDistance
	}

	return math.Sqrt(math.Pow(first.X()-second.X(), 2) + math.Pow(first.Y()-second.Y(), 2))
}

func euclideanAchievability(first Point, second Point, barrier Barrier) bool {
	var barrierLine = barrier.LineSegment

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

	return !(projectionAchievability(first.X(), second.X(), barrierLine.StartPoint.X(), barrierLine.FinishPoint.X()) &&
		projectionAchievability(first.Y(), second.Y(), barrierLine.StartPoint.Y(), barrierLine.FinishPoint.Y()) &&
		getTriangleArea(first, second, barrierLine.StartPoint)*getTriangleArea(first, second, barrierLine.FinishPoint) <= Eps &&
		getTriangleArea(barrierLine.StartPoint, barrierLine.FinishPoint, first)*getTriangleArea(barrierLine.StartPoint, barrierLine.FinishPoint, second) <= Eps)
}
