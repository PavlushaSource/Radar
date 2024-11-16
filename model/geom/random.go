package geom

type randomGeom struct {
	height     float64
	width      float64
	distance   Distance
	randomCore int
}

func (geom *randomGeom) Height() float64 {
	return geom.height
}

func (geom *randomGeom) Width() float64 {
	return geom.width
}

func (geom *randomGeom) Distance(first Point, second Point) float64 {
	return geom.distance(first, second)
}

func (geom *randomGeom) MovePoint(point Point) {
	point.set(0, 0)
}

func (geom *randomGeom) NewPoint() Point {
	return newPoint(0, 0)
}

func NewRandomGeom(height float64, width float64, distance Distance) Geom {
	geom := new(randomGeom)
	geom.height = height
	geom.width = width
	geom.distance = distance
	geom.randomCore = 52
	return geom
}
