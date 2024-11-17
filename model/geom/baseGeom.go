package geom

import "math/rand"

type baseGeom struct {
	height   float64
	width    float64
	barriers []Barrier
	distance Distance
	rng      *rand.Rand
}

func (geom *baseGeom) Height() float64 {
	return geom.height
}

func (geom *baseGeom) Width() float64 {
	return geom.width
}

func (geom *baseGeom) Barriers() []Barrier {
	return geom.barriers
}
func (geom *baseGeom) Distance(first Point, second Point) float64 {
	return geom.distance(first, second, geom.barriers)
}

func (geom *baseGeom) NewPoint() Point {
	x := geom.rng.Float64() * geom.width
	y := geom.rng.Float64() * geom.height

	return newPoint(x, y)
}
