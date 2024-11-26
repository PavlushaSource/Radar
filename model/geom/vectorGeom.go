package geom

import (
	"math"
	"math/rand"
	"time"

	"github.com/PavlushaSource/Radar/model/core/rnd"
)

const maxPointMovingDegree = 359

// vectorGeom is a struct for vector geometry
// In vectorGeom, the points move straight, deviating by some angle
type vectorGeom struct {
	baseGeom
}

// MovePoint moves the point to other coordinates that meet the constraints of the field
func (geom *vectorGeom) MovePoint(point Point) {
	movedPoint := geom.newVectorRandomPoint(point)
	movedPoint = geom.LimitPointMovement(point, movedPoint)
	movedPoint = geom.CorrectMovingAchievable(point, movedPoint, geom.newVectorRandomPoint)

	point.set(movedPoint.X(), movedPoint.Y())
}

// newVectorRandomPoint generates new random point, taking into account the coordinates of the old point
func (geom *vectorGeom) newVectorRandomPoint(point Point) Point {
	distance := geom.rndAsync.Float64ByFloat64(point.X()*float64(geom.rndCounter)) * math.Max(geom.width, geom.width)
	geom.IncrementRndCounter()

	degree := geom.rndAsync.Float64ByFloat64(point.Y()*float64(geom.rndCounter)) * maxPointMovingDegree
	geom.IncrementRndCounter()

	radians := degree * math.Pi / 180.0

	x := math.Max(math.Min(point.X()+distance*math.Cos(radians), geom.width), 0)
	y := math.Max(math.Min(point.Y()+distance*math.Sin(radians), geom.height), 0)

	return NewPoint(x, y)
}

// NewVectorGeom creates new instance of the vector geometry
func NewVectorGeom(
	height float64,
	width float64,
	barriers []Barrier,
	maxMoveDistance float64,
	distance Distance,
	rndAsync rnd.RndAsync,
) Geom {
	geom := new(vectorGeom)

	geom.height = height
	geom.width = width
	geom.barriers = barriers
	geom.distance = distance
	geom.maxMoveDistance = maxMoveDistance
	geom.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	geom.rndAsync = rndAsync

	return geom
}
