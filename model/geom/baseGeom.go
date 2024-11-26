package geom

import (
	"math/rand"

	"github.com/PavlushaSource/Radar/model/core/rnd"
)

// maxRndCounter is a number used to work with [RndAsync]
const maxRndCounter = 10000000

// baseGeom is a base struct for all geometry implementations
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

// Height returns a height of the field
func (geom *baseGeom) Height() float64 {
	return geom.height
}

// Width returns a width of the field
func (geom *baseGeom) Width() float64 {
	return geom.width
}

// Barriers returns a list of all barriers on the field
func (geom *baseGeom) Barriers() []Barrier {
	return geom.barriers
}

// Distance returns a distance between two points, taking into account the barriers on the field
func (geom *baseGeom) Distance(first Point, second Point) float64 {
	return geom.distance(first, second, geom.barriers)
}

// IncrementRndCounter increments the counter used to generate random numbers in geom package
func (geom *baseGeom) IncrementRndCounter() {
	geom.rndCounter += 1

	if geom.rndCounter > maxRndCounter {
		geom.rndCounter = geom.rndCounter % maxRndCounter
	}
}

// LimitPointMovement limits point movement.
// If the motion vector of the point is too large, the method tries to reduce this vector
// Returns a new point to move
func (geom *baseGeom) LimitPointMovement(point, movedPoint Point) Point {
	// Dogs move along a vector on the UI, so we consider the Euclidean distance
	distance := EuclideanDistance(point, movedPoint, geom.barriers)

	if distance-geom.maxMoveDistance > Eps {
		limitedPoint := movedPoint
		if distance != InfDistance {
			limitedPoint = geom.ReducePointMovement(point, movedPoint, distance/geom.maxMoveDistance)
		} else {
			for j := 0; j < 13; j++ {
				if arePointsAchievable(point, limitedPoint, geom.barriers, euclideanAchievability) {
					break
				} else {
					limitedPoint = geom.ReducePointMovement(point, limitedPoint, 2)
				}
			}
		}

		movedPoint.set(limitedPoint.X(), limitedPoint.Y())
	}

	return movedPoint
}

// ReducePointMovement reduces the motion vector of the point by k times
// Returns a new point to move
func (geom *baseGeom) ReducePointMovement(point, movedPoint Point, k float64) Point {
	newX := (movedPoint.X()-point.X())/k + point.X()
	newY := (movedPoint.Y()-point.Y())/k + point.Y()

	return NewPoint(newX, newY)
}

// CorrectMovingAchievable tries to correct a motion vector if needed.
// For example, the vector crosses the barrier, then the method tries to efficiently generate
// a new vector of movement of the point
// Returns a new point to move
// Returns the old point if the correct motion vector is not found
func (geom *baseGeom) CorrectMovingAchievable(oldPoint, movedPoint Point, newPointGenerator func(point Point) Point) Point {
	if !arePointsAchievable(oldPoint, movedPoint, geom.barriers, euclideanAchievability) {
		newPoint := geom.withSafetySearchMovedPoint(oldPoint, newPointGenerator)

		if newPoint == nil {
			return oldPoint
		} else {
			return newPoint
		}
	}

	return movedPoint
}

// withSafetySearchMovedPoint is a method for safety searching of the motion vector
// Returns a new point to move
func (geom *baseGeom) withSafetySearchMovedPoint(oldPoint Point, newPointGenerate func(point Point) Point) Point {
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

//TODO: What about a point generated in barrier?

// NewRandomPoint generates random point
func (geom *baseGeom) NewRandomPoint() Point {
	x := geom.rndAsync.Float64ByFloat64(float64(geom.rndCounter)) * geom.width
	geom.IncrementRndCounter()

	y := geom.rndAsync.Float64ByFloat64(float64(geom.rndCounter)) * geom.height
	geom.IncrementRndCounter()

	return NewPoint(x, y)
}
