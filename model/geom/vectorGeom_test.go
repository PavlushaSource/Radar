package geom

import (
	"github.com/PavlushaSource/Radar/model/core/rnd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVectorGeom(t *testing.T) {
	rndAsync := rnd.NewRndCore()

	g := NewVectorGeom(
		100.0, // height
		200.0, // width
		nil,   // barriers
		10.0,  // maxMoveDistance
		nil,   // distance (not used here)
		rndAsync,
	)

	assert.NotNil(t, g, "Expected non-nil vector geometry")
}

func TestNewVectorRandomPoint(t *testing.T) {
	rndAsync := rnd.NewRndCore()

	g := NewVectorGeom(
		100.0, // height
		200.0, // width
		nil,   // barriers
		10.0,  // maxMoveDistance
		nil,   // distance (not used here)
		rndAsync,
	).(*vectorGeom)

	point := NewPoint(50.0, 50.0)
	newPoint := g.newVectorRandomPoint(point)

	assert.GreaterOrEqual(t, newPoint.X(), 0.0)
	assert.LessOrEqual(t, newPoint.X(), 200.0)
	assert.GreaterOrEqual(t, newPoint.Y(), 0.0)
	assert.LessOrEqual(t, newPoint.Y(), 100.0)
}

func TestMovePoint_Vector(t *testing.T) {
	rndAsync := rnd.NewRndCore()

	g := NewVectorGeom(
		100.0, // height
		200.0, // width
		nil,   // barriers
		10.0,  // maxMoveDistance
		nil,   // distance (not used here)
		rndAsync,
	).(*vectorGeom)

	point := NewPoint(50.0, 50.0)
	g.MovePoint(point)

	assert.GreaterOrEqual(t, point.X(), 0.0)
	assert.LessOrEqual(t, point.X(), 200.0)
	assert.GreaterOrEqual(t, point.Y(), 0.0)
	assert.LessOrEqual(t, point.Y(), 100.0)
}

func TestVectorMovementBoundaryConditions(t *testing.T) {
	rndAsync := rnd.NewRndCore()

	g := NewVectorGeom(
		100.0, // height
		200.0, // width
		nil,   // barriers
		10.0,  // maxMoveDistance
		nil,   // distance (not used here)
		rndAsync,
	).(*vectorGeom)

	point := NewPoint(195.0, 95.0)
	newPoint := g.newVectorRandomPoint(point)

	assert.GreaterOrEqual(t, newPoint.X(), 0.0)
	assert.LessOrEqual(t, newPoint.X(), 200.0)
	assert.GreaterOrEqual(t, newPoint.Y(), 0.0)
	assert.LessOrEqual(t, newPoint.Y(), 100.0)
}
