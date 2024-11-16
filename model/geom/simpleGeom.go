package geom

import (
	"math/rand"
	"time"
)

type simpleGeom struct {
	baseGeom
}

func (geom *simpleGeom) MovePoint(point Point) {
	x := geom.rng.Float64() * geom.width
	y := geom.rng.Float64() * geom.height

	point.set(x, y)
}

func NewSimpleGeom(height float64, width float64, barriers []Barrier, distance Distance) Geom {
	geom := new(simpleGeom)
	geom.height = height
	geom.width = width
	geom.barriers = barriers
	geom.distance = distance
	geom.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	return geom
}
