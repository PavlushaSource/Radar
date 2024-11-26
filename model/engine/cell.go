package engine

import (
	"math"

	"github.com/PavlushaSource/Radar/model/geom"
)

func numColumns(width float64, cellSize float64) int {
	return int(math.Ceil(width / cellSize))
}

func numRows(height float64, cellSize float64) int {
	return int(math.Ceil(height / cellSize))
}

func (processor *processor) cellByRowColumn(row int, column int) int {
	return column*processor.numRows + row
}

func (processor *processor) tryGetCell(point geom.Point) (bool, int) {
	column := int(math.Trunc(point.X() / processor.cellSize))
	row := int(math.Trunc(point.Y() / processor.cellSize))
	return 0 <= row &&
			row < processor.numRows &&
			0 <= column &&
			column < processor.numColumns,
		column*processor.numRows + row
}
