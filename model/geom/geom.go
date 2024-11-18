package geom

type Geom interface {
	Height() float64
	Width() float64
	Barriers() []Barrier

	Distance(first Point, second Point) float64

	NewRandomPoint() Point
	MovePoint(point Point)
}
