package geom

import (
	"math/rand"
	"time"

	"github.com/PavlushaSource/Radar/model/core/rnd"
)

// simpleGeom is a struct for simple geometry
// In simpleGeom, the points can move to a random point
type simpleGeom struct {
	baseGeom
}

// MovePoint moves the point to other coordinates that meet the constraints of the field
func (geom *simpleGeom) MovePoint(point Point) {
	movedPoint := geom.newSimpleRandomPoint(point)
	movedPoint = geom.LimitPointMovement(point, movedPoint)
	movedPoint = geom.CorrectMovingAchievable(point, movedPoint, geom.newSimpleRandomPoint)

	point.set(movedPoint.X(), movedPoint.Y())
}

// newSimpleRandomPoint generates new random point, taking into account the coordinates of the old point
func (geom *simpleGeom) newSimpleRandomPoint(point Point) Point {
	x := geom.rndAsync.Float64ByFloat64(point.X()*float64(geom.rndCounter)) * geom.width
	geom.IncrementRndCounter()

	y := geom.rndAsync.Float64ByFloat64(point.Y()*float64(geom.rndCounter)) * geom.height
	geom.IncrementRndCounter()

	return NewPoint(x, y)
}

// NewSimpleGeom creates new instance of the simple geometry
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
