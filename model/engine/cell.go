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

func (processor *processor) cell(point geom.Point) int {
	column := int(math.Trunc(point.X() / processor.radius))
	row := int(math.Trunc(point.Y() / processor.radius))
	return column*processor.numRows + row
}

func (processor *processor) checkCellColumnRow(column int, row int) bool {
	if 0 > row || row >= processor.numRows {
		return false
	}
	if 0 > column || column >= processor.numColumns {
		return false
	}
	return true
}

type neighbourCellExtractor func(cell int) (bool, int)

func (processor *processor) tryGetUpCell(cell int) (bool, int) {
	column := cell / processor.numRows
	row := cell%processor.numRows + 1

	if !processor.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*processor.numRows + row
}

func (processor *processor) tryGetDownCell(cell int) (bool, int) {
	column := cell / processor.numRows
	row := cell%processor.numRows - 1

	if !processor.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*processor.numRows + row
}

func (processor *processor) tryGetLeftCell(cell int) (bool, int) {
	column := cell/processor.numRows - 1
	row := cell % processor.numRows

	if !processor.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*processor.numRows + row
}

func (processor *processor) tryGetRightCell(cell int) (bool, int) {
	column := cell/processor.numRows + 1
	row := cell % processor.numRows

	if !processor.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*processor.numRows + row
}

func (processor *processor) tryGetUpLeftCell(cell int) (bool, int) {
	column := cell/processor.numRows - 1
	row := cell%processor.numRows + 1

	if !processor.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*processor.numRows + row
}

func (processor *processor) tryGetUpRightCell(cell int) (bool, int) {
	column := cell/processor.numRows + 1
	row := cell%processor.numRows + 1

	if !processor.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*processor.numRows + row
}

func (processor *processor) tryGetDownLeftCell(cell int) (bool, int) {
	column := cell/processor.numRows - 1
	row := cell%processor.numRows - 1

	if !processor.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*processor.numRows + row
}

func (processor *processor) tryGetDownRightCell(cell int) (bool, int) {
	column := cell/processor.numRows + 1
	row := cell%processor.numRows - 1

	if !processor.checkCellColumnRow(column, row) {
		return false, cellDefault
	}
	return true, column*processor.numRows + row
}
