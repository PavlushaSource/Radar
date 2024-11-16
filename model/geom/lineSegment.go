package geom

type LineSegment struct {
	startPoint  Point
	finishPoint Point
}

// TODO: move to model consts
const Eps = 1e-9

func IsLineSegmentsIntersect(ls1 LineSegment, ls2 LineSegment) bool {
	// checks projections intersection
	var isLineSegmentsProjectionIntersect = func(ls1Start, ls1Finish, ls2Start, ls2Finish float64) bool {
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
		return 0.5*(p2.X()-p1.X())*(p3.Y()-p1.Y()) - (p2.Y()-p1.Y())*(p3.X()-p1.X())
	}

	return isLineSegmentsProjectionIntersect(ls1.startPoint.X(), ls1.finishPoint.X(), ls2.startPoint.X(), ls2.finishPoint.X()) &&
		isLineSegmentsProjectionIntersect(ls1.startPoint.Y(), ls1.finishPoint.Y(), ls2.startPoint.Y(), ls2.finishPoint.Y()) &&
		getTriangleArea(ls1.startPoint, ls1.finishPoint, ls2.startPoint)*getTriangleArea(ls1.startPoint, ls1.finishPoint, ls2.finishPoint) <= Eps &&
		getTriangleArea(ls2.startPoint, ls2.finishPoint, ls1.startPoint)*getTriangleArea(ls2.startPoint, ls2.finishPoint, ls1.finishPoint) <= Eps
}
