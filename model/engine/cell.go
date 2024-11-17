package engine

import (
	"math"

	"github.com/PavlushaSource/Radar/model/geom"
)

const cellDefault int64 = -1

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

func (engine *engine) checkCellColumnRow(column int64, row int64) bool {
	if 0 > row || row >= engine.numRows {
		return false
	}
	if 0 > column || column >= engine.numColumns {
		return false
	}
	return true
}

type neighbourCellExtractor func(cell int64) (bool, int64)

func (engine *engine) tryGetUpCell(cell int64) (bool, int64) {
	column := cell / engine.numRows
	row := cell%engine.numRows + 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetDownCell(cell int64) (bool, int64) {
	column := cell / engine.numRows
	row := cell%engine.numRows - 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetLeftCell(cell int64) (bool, int64) {
	column := cell/engine.numRows - 1
	row := cell % engine.numRows

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetRightCell(cell int64) (bool, int64) {
	column := cell/engine.numRows + 1
	row := cell % engine.numRows

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetUpLeftCell(cell int64) (bool, int64) {
	column := cell/engine.numRows - 1
	row := cell%engine.numRows + 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetUpRightCell(cell int64) (bool, int64) {
	column := cell/engine.numRows + 1
	row := cell%engine.numRows + 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetDownLeftCell(cell int64) (bool, int64) {
	column := cell/engine.numRows - 1
	row := cell%engine.numRows - 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetDownRightCell(cell int64) (bool, int64) {
	column := cell/engine.numRows + 1
	row := cell%engine.numRows - 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}
