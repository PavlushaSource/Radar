package geom

import (
	"math/rand"

	"github.com/PavlushaSource/Radar/model/core/rnd"
)

type baseGeom struct {
	height          float64
	width           float64
	barriers        []Barrier
	maxMoveDistance float64
	distance        Distance
	rnd             *rand.Rand
	rndAsync        rnd.RndAsync
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

func (geom *baseGeom) LimitPointMovement(point, movedPoint Point) Point {
	// Cats move along a vector on the UI, so we consider the Euclidean distance
	distance := EuclideanDistance(point, movedPoint, geom.barriers)

	limitedPoint := movedPoint

	if distance-maxPointMovingDegree > Eps {
		k := maxPointMovingDegree / distance

		limitedX := (movedPoint.X() - point.X()) * k
		limitedY := (movedPoint.Y() - point.Y()) * k

		limitedPoint.set(limitedX, limitedY)
	}

	return limitedPoint
}

func (geom *baseGeom) NewRandomPoint() Point {
	x := geom.rnd.Float64() * geom.width
	y := geom.rnd.Float64() * geom.height

	return NewPoint(x, y)
}
