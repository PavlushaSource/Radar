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

	movedPoint := geom.LimitPointMovement(point, NewPoint(x, y))
	movedPoint = geom.CorrectMovingAchievable(point, movedPoint, geom.newSimpleRandomPoint)

	point.set(movedPoint.X(), movedPoint.Y())
}

func (geom *simpleGeom) newSimpleRandomPoint(_ Point) Point {
	return geom.NewRandomPoint()
}

func NewSimpleGeom(
	height float64,
	width float64,
	barriers []Barrier,
	maxMoveDistance float64,
	distance Distance,
	rndAsync rnd.RndAsync,
) Geom {
	geom := new(simpleGeom)

	geom.height = height
	geom.width = width
	geom.barriers = barriers
	geom.distance = distance
	geom.maxMoveDistance = maxMoveDistance
	geom.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	geom.rndAsync = rndAsync

	return geom
}
