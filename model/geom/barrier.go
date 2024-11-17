package geom

type Barrier interface {
	StartPoint() Point
	FinishPoint() Point
	set(start Point, finish Point)
}

type barrier struct {
	startPoint  Point
	finishPoint Point
}

func (barrier *barrier) StartPoint() Point {
	return barrier.startPoint
}

func (barrier *barrier) FinishPoint() Point {
	return barrier.finishPoint
}

func (barrier *barrier) set(start Point, finish Point) {
	barrier.startPoint = start
	barrier.finishPoint = finish
}

func NewBarrier(start Point, finish Point) Barrier {
	barrier := new(barrier)
	barrier.set(start, finish)
	return barrier
}
