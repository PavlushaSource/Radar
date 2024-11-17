package engine

import (
	"math"

	"github.com/PavlushaSource/Radar/model/geom"
)

func numColumns(width float64, radius float64) int64 {
	return int64(math.Ceil(width / radius))
}

func numRows(height float64, radius float64) int64 {
	return int64(math.Ceil(height / radius))
}

func (engine *engine) cell(point geom.Point) int64 {
	column := int64(math.Ceil(point.X() / engine.radius))
	row := int64(math.Ceil(point.Y() / engine.radius))
	return column*engine.numRows + row
}

// TODO: implemet cell navigation
func (engine *engine) tryGetUp(cell int64) int64 {
	// column := cell / engine.numRows
	// row := cell % engine.numRows

	// if (row + 1) <
	return int64(0)
}

func (engine *engine) tryGetDown(cell int64) int64 {
	return int64(0)
}

func (engine *engine) tryGetLeft(cell int64) int64 {
	return int64(0)
}

func (engine *engine) tryGetRight(cell int64) int64 {
	return int64(0)
}

func (engine *engine) tryGetUpLeft(cell int64) int64 {
	return int64(0)
}

func (engine *engine) tryGetUpRighr(cell int64) int64 {
	return int64(0)
}

func (engine *engine) tryGetDownLeft(cell int64) int64 {
	return int64(0)
}

func (engine *engine) tryGetDownRight(cell int64) int64 {
	return int64(0)
}
