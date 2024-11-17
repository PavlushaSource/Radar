package geom

type Point interface {
	X() float64
	Y() float64
	Copy() Point
	set(x float64, y float64)
}

type point struct {
	x float64
	y float64
}

func (point *point) X() float64 {
	return point.x
}

func (point *point) Y() float64 {
	return point.y
}

func (original *point) Copy() Point {
	clon := new(point)
	clon.x = original.x
	clon.y = original.y
	return clon
}

func (point *point) set(x float64, y float64) {
	point.x = x
	point.y = y
}

func newPoint(x float64, y float64) Point {
	point := new(point)
	point.set(x, y)
	return point
}
