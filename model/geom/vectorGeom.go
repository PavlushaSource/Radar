package geom

import (
	"math"
	"math/rand"
	"time"

	"github.com/PavlushaSource/Radar/model/core/rnd"
)

const maxPointMovingDegree = 359

type vectorGeom struct {
	baseGeom
}

func (geom *vectorGeom) MovePoint(point Point) {
	movedPoint := geom.newVectorRandomPoint(point)
	movedPoint = geom.LimitPointMovement(point, movedPoint)
	movedPoint = geom.CorrectMovingAchievable(point, movedPoint, geom.newVectorRandomPoint)

	point.set(movedPoint.X(), movedPoint.Y())
}

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
