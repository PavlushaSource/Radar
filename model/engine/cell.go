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

func (processor *processor) cellByColumnRow(column int, row int) int {
	return column*processor.numRows + row
}

func (processor *processor) tryGetCell(point geom.Point) (bool, int) {
	column := int(math.Trunc(point.X() / processor.radius))
	row := int(math.Trunc(point.Y() / processor.radius))
	return 0 <= row &&
			row < processor.numRows &&
			0 <= column &&
			column < processor.numColumns,
		column*processor.numRows + row
}
