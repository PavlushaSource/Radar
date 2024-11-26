package geom

// Geom is a type interface representing all geometric methods needs for the engine to update dogs positions and status
type Geom interface {
	Height() float64
	Width() float64
	Barriers() []Barrier

	Distance(first Point, second Point) float64

	NewRandomPoint() Point
	MovePoint(point Point)
}
