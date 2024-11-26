package geom

// Point is a type interface represents point on the 2d field
type Point interface {
	X() float64
	Y() float64
	Copy(dst Point)
	set(x float64, y float64)
}

// point is a struct represents point on the 2d field
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

// Copy set this point coordinates to dst point
func (point *point) Copy(dst Point) {
	dst.set(point.x, point.y)
}

// set new coordinates to the point
func (point *point) set(x float64, y float64) {
	point.x = x
	point.y = y
}

// NewPoint returns new instance of the Point
func NewPoint(x float64, y float64) Point {
	point := new(point)
	point.set(x, y)
	return point
}
