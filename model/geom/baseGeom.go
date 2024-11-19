package geom

import (
	"math/rand"

	"github.com/PavlushaSource/Radar/model/core/rnd"
)

const maxRndCounter = 10000000

type baseGeom struct {
	height          float64
	width           float64
	barriers        []Barrier
	maxMoveDistance float64
	distance        Distance
	rnd             *rand.Rand
	rndAsync        rnd.RndAsync
	rndCounter      int32
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

func (geom *baseGeom) IncrementRndCounter() {
	geom.rndCounter += 1

	if geom.rndCounter > maxRndCounter {
		geom.rndCounter = geom.rndCounter % maxRndCounter
	}
}

func (geom *baseGeom) LimitPointMovement(point, movedPoint Point) Point {
	// Cats move along a vector on the UI, so we consider the Euclidean distance
	distance := EuclideanDistance(point, movedPoint, geom.barriers)

	if distance-geom.maxMoveDistance > Eps {
		limitedPoint := geom.ReducePointMovement(point, movedPoint, distance/geom.maxMoveDistance)

		movedPoint.set(limitedPoint.X(), limitedPoint.Y())
	}

	return movedPoint
}

func (geom *baseGeom) ReducePointMovement(point, movedPoint Point, k float64) Point {
	newX := (movedPoint.X()-point.X())/k + point.X()
	newY := (movedPoint.Y()-point.Y())/k + point.Y()

	return NewPoint(newX, newY)
}

// CorrectMovingAchievable tries to correct movedPoint if needed.
// Returns `oldPoint` if correct movedPoint is not found.
func (geom *baseGeom) CorrectMovingAchievable(oldPoint, movedPoint Point, newPointGenerator func(point Point) Point) Point {
	if !arePointsAchievable(oldPoint, movedPoint, geom.barriers, euclideanAchievability) {
		newPoint := geom.withSaftySearchMovedPoint(oldPoint, newPointGenerator)

		if newPoint == nil {
			return oldPoint
		} else {
			return newPoint
		}
	}

	return movedPoint
}

func (geom *baseGeom) withSaftySearchMovedPoint(oldPoint Point, newPointGenerate func(point Point) Point) Point {
	for i := 0; i < 5; i++ {
		point := geom.LimitPointMovement(oldPoint, newPointGenerate(oldPoint))

		for j := 0; j < 13; j++ {
			if arePointsAchievable(oldPoint, point, geom.barriers, euclideanAchievability) {
				return point
			} else {
				point = geom.ReducePointMovement(oldPoint, point, 2)
			}
		}
	}

	return nil
}

// TODO: What about barriers?
func (geom *baseGeom) NewRandomPoint() Point {
	x := geom.rndAsync.Float64ByFloat64(float64(geom.rndCounter)) * geom.width
	geom.IncrementRndCounter()

	y := geom.rndAsync.Float64ByFloat64(float64(geom.rndCounter)) * geom.height
	geom.IncrementRndCounter()

	return NewPoint(x, y)
}
