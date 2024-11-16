package geom

type Geom interface {
	Height() float64
	Width() float64
	Distance(first Point, second Point) float64

	NewPoint() Point
	MovePoint(point Point)
}
