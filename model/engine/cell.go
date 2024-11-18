package engine

import (
	"math"

	"github.com/PavlushaSource/Radar/model/geom"
)

func numColumns(width float64, radius float64) int {
	return int(math.Ceil(width / radius))
}

func numRows(height float64, radius float64) int {
	return int(math.Ceil(height / radius))
}

func (processor *processor) tryGetCell(point geom.Point) (bool, int) {
	column := int(math.Trunc(point.X() / processor.radius))
	row := int(math.Trunc(point.Y() / processor.radius))
	return processor.checkCellColumnRow(column, row), column*processor.numRows + row
}

func (processor *processor) checkCellColumnRow(column int, row int) bool {
	return 0 <= row && row < processor.numRows && 0 <= column && column < processor.numColumns
}

type neighbourCellExtractor func(cell int) (bool, int)

func (processor *processor) tryGetUpCell(cell int) (bool, int) {
	column := cell / processor.numRows
	row := cell%processor.numRows + 1

	return processor.checkCellColumnRow(column, row), column*processor.numRows + row
}

func (processor *processor) tryGetDownCell(cell int) (bool, int) {
	column := cell / processor.numRows
	row := cell%processor.numRows - 1

	return processor.checkCellColumnRow(column, row), column*processor.numRows + row
}

func (processor *processor) tryGetLeftCell(cell int) (bool, int) {
	column := cell/processor.numRows - 1
	row := cell % processor.numRows

	return processor.checkCellColumnRow(column, row), column*processor.numRows + row
}

func (processor *processor) tryGetRightCell(cell int) (bool, int) {
	column := cell/processor.numRows + 1
	row := cell % processor.numRows

	return processor.checkCellColumnRow(column, row), column*processor.numRows + row
}

func (processor *processor) tryGetUpLeftCell(cell int) (bool, int) {
	column := cell/processor.numRows - 1
	row := cell%processor.numRows + 1

	return processor.checkCellColumnRow(column, row), column*processor.numRows + row
}

func (processor *processor) tryGetUpRightCell(cell int) (bool, int) {
	column := cell/processor.numRows + 1
	row := cell%processor.numRows + 1

	return processor.checkCellColumnRow(column, row), column*processor.numRows + row
}

func (processor *processor) tryGetDownLeftCell(cell int) (bool, int) {
	column := cell/processor.numRows - 1
	row := cell%processor.numRows - 1

	return processor.checkCellColumnRow(column, row), column*processor.numRows + row
}

func (processor *processor) tryGetDownRightCell(cell int) (bool, int) {
	column := cell/processor.numRows + 1
	row := cell%processor.numRows - 1

	return processor.checkCellColumnRow(column, row), column*processor.numRows + row
}
