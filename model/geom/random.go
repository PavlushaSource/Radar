package geom

type randomGeom struct {
	length     float64
	width      float64
	distance   Distance
	randomCore int
}

func (geom *randomGeom) Length() float64 {
	return geom.length
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

func NewRandomGeom(length float64, width float64, distance Distance) Geom {
	geom := new(randomGeom)
	geom.length = length
	geom.width = width
	geom.distance = distance
	geom.randomCore = 52
	return geom
}
