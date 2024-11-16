package geom

type Barrier struct {
	lineSegment LineSegment
}

func IsLineSegmentIntersectWithBarrier(ls LineSegment, barriers []Barrier) bool {
	var flag = false
	for _, barrier := range barriers {
		flag = flag || IsLineSegmentsIntersect(ls, barrier.lineSegment)
	}

	return flag
}
