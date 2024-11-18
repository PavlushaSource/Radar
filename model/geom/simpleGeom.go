package geom

import (
	"math/rand"
	"time"

	"github.com/PavlushaSource/Radar/model/core/rnd"
)

type simpleGeom struct {
	baseGeom
}

func (geom *simpleGeom) MovePoint(point Point) {
	x := geom.rndAsync.Float64ByFloat64(point.X()) * geom.width
	y := geom.rndAsync.Float64ByFloat64(point.Y()) * geom.height

	point.set(x, y)
}

func NewSimpleGeom(height float64, width float64, barriers []Barrier, distance Distance, rndAsync rnd.RndAsync) Geom {
	geom := new(simpleGeom)
	geom.height = height
	geom.width = width
	geom.barriers = barriers
	geom.distance = distance
	geom.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	geom.rndAsync = rndAsync
	return geom
}
