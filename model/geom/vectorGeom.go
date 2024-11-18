package geom

import (
	"math"
	"math/rand"
	"time"

	"github.com/PavlushaSource/Radar/model/core/rnd"
)

const fullCircleDegree = 360

type vectorGeom struct {
	baseGeom
}

func (geom *vectorGeom) MovePoint(point Point) {
	distance := geom.rndAsync.Float64ByFloat64(point.X())

	rndForDegree := geom.rndAsync.Float64ByFloat64(point.Y())
	degree := rndForDegree - (rndForDegree/fullCircleDegree)*fullCircleDegree

	radians := degree * math.Pi / 180.0

	x := math.Max(math.Min(point.X()+distance*math.Cos(radians), geom.width), 0)
	y := math.Max(math.Min(point.Y()+distance*math.Sin(radians), geom.height), 0)

	movedPoint := geom.LimitPointMovement(point, NewPoint(x, y))
	movedPoint = geom.CorrectMovingAchievable(point, movedPoint, geom.newVectorRandomPoint)

	point.set(movedPoint.X(), movedPoint.Y())
}

func (geom *vectorGeom) newVectorRandomPoint(point Point) Point {
	distance := geom.rnd.Float64() * math.Max(geom.width, geom.width)
	degree := geom.rnd.Float64() * (fullCircleDegree - 1)

	radians := degree * math.Pi / 180.0

	x := math.Max(math.Min(point.X()+distance*math.Cos(radians), geom.width), 0)
	y := math.Max(math.Min(point.Y()+distance*math.Sin(radians), geom.height), 0)

	return NewPoint(x, y)
}

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
