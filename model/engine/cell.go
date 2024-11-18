package engine

import (
	"math"

	"github.com/PavlushaSource/Radar/model/geom"
)

const cellDefault int = -1

func numColumns(width float64, radius float64) int {
	return int(math.Ceil(width / radius))
}

func numRows(height float64, radius float64) int {
	return int(math.Ceil(height / radius))
}

func (engine *engine) cell(point geom.Point) int {
	column := int(math.Trunc(point.X() / engine.radius))
	row := int(math.Trunc(point.Y() / engine.radius))
	return column*engine.numRows + row
}

func (engine *engine) checkCellColumnRow(column int, row int) bool {
	if 0 > row || row >= engine.numRows {
		return false
	}
	if 0 > column || column >= engine.numColumns {
		return false
	}
	return true
}

type neighbourCellExtractor func(cell int) (bool, int)

func (engine *engine) tryGetUpCell(cell int) (bool, int) {
	column := cell / engine.numRows
	row := cell%engine.numRows + 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetDownCell(cell int) (bool, int) {
	column := cell / engine.numRows
	row := cell%engine.numRows - 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetLeftCell(cell int) (bool, int) {
	column := cell/engine.numRows - 1
	row := cell % engine.numRows

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetRightCell(cell int) (bool, int) {
	column := cell/engine.numRows + 1
	row := cell % engine.numRows

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetUpLeftCell(cell int) (bool, int) {
	column := cell/engine.numRows - 1
	row := cell%engine.numRows + 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetUpRightCell(cell int) (bool, int) {
	column := cell/engine.numRows + 1
	row := cell%engine.numRows + 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetDownLeftCell(cell int) (bool, int) {
	column := cell/engine.numRows - 1
	row := cell%engine.numRows - 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}

func (engine *engine) tryGetDownRightCell(cell int) (bool, int) {
	column := cell/engine.numRows + 1
	row := cell%engine.numRows - 1

	if !engine.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*engine.numRows + row
}
