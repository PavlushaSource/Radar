package geom

import (
	"testing"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/stretchr/testify/assert"
)

func TestNewSimpleGeom(t *testing.T) {
	rndAsync := rnd.NewRndCore()

	g := NewSimpleGeom(
		100.0, // height
		200.0, // width
		nil,   // barriers
		10.0,  // maxMoveDistance
		nil,   // distance (not used here)
		rndAsync,
	)

	assert.NotNil(t, g, "Expected non-nil geometry")
}

func TestNewSimpleRandomPoint(t *testing.T) {
	rndAsync := rnd.NewRndCore()

	g := NewSimpleGeom(
		100.0, // height
		200.0, // width
		nil,   // barriers
		10.0,  // maxMoveDistance
		nil,   // distance (not used here)
		rndAsync,
	).(*simpleGeom)

	point := NewPoint(50.0, 50.0)
	newPoint := g.newSimpleRandomPoint(point)

	assert.GreaterOrEqual(t, newPoint.X(), 0.0)
	assert.LessOrEqual(t, newPoint.X(), 200.0)
	assert.GreaterOrEqual(t, newPoint.Y(), 0.0)
	assert.LessOrEqual(t, newPoint.Y(), 100.0)
}

func TestLimitPointMovement(t *testing.T) {
	rndAsync := rnd.NewRndCore()

	g := NewSimpleGeom(
		100.0, // height
		200.0, // width
		nil,   // barriers
		10.0,  // maxMoveDistance
		nil,   // distance (not used here)
		rndAsync,
	).(*simpleGeom)

	point := NewPoint(195.0, 95.0)
	newPoint := g.newSimpleRandomPoint(point)
	limitedPoint := g.LimitPointMovement(point, newPoint)

	assert.GreaterOrEqual(t, limitedPoint.X(), 0.0)
	assert.LessOrEqual(t, limitedPoint.X(), 200.0)
	assert.GreaterOrEqual(t, limitedPoint.Y(), 0.0)
	assert.LessOrEqual(t, limitedPoint.Y(), 100.0)
}
