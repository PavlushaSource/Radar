package geom

import (
	"math/rand"

	"github.com/PavlushaSource/Radar/model/core/rnd"
)

type baseGeom struct {
	height   float64
	width    float64
	barriers []Barrier
	distance Distance
	rnd      *rand.Rand
	rndAsync rnd.RndAsync
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
	x := geom.rnd.Float64() * geom.width
	y := geom.rnd.Float64() * geom.height

	return newPoint(x, y)
}
