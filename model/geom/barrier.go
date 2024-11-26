package geom

// Barrier interface represent a Barrier segment on the field
type Barrier interface {
	StartPoint() Point
	FinishPoint() Point
	set(start Point, finish Point)
}

// barrier struct represent a barrier segment on the field
type barrier struct {
	startPoint  Point
	finishPoint Point
}

// StartPoint returns a start point of the barrier segment
func (barrier *barrier) StartPoint() Point {
	return barrier.startPoint
}

// FinishPoint returns a finish point of the barrier segment
func (barrier *barrier) FinishPoint() Point {
	return barrier.finishPoint
}

// set new barrier segment
func (barrier *barrier) set(start Point, finish Point) {
	barrier.startPoint = start
	barrier.finishPoint = finish
}

// NewBarrier return a new instance of the Barrier
func NewBarrier(start Point, finish Point) Barrier {
	barrier := new(barrier)
	barrier.set(start, finish)
	return barrier
}
