package geom

type Point interface {
	X() float64
	Y() float64
	Copy(dst Point)
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

func (point *point) Copy(dst Point) {
	dst.set(point.x, point.y)
}

func (point *point) set(x float64, y float64) {
	point.x = x
	point.y = y
}

func NewPoint(x float64, y float64) Point {
	point := new(point)
	point.set(x, y)
	return point
}
