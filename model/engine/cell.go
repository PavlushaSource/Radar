package engine

import (
	"math"

	"github.com/PavlushaSource/Radar/model/geom"
)

// numColumns receives the field width and cell size and
// returns the number of columns.
func numColumns(width float64, cellSize float64) int {
	return int(math.Ceil(width / cellSize))
}

// numRows receives the field height and cell size and
// returns the number of rows.
func numRows(height float64, cellSize float64) int {
	return int(math.Ceil(height / cellSize))
}

// cellByRowColumn receives row and column indices and
// returns the cell index.
func (processor *processor) cellByRowColumn(row int, column int) int {
	return column*processor.numRows + row
}

// tryGetCell receives a point in the plane and
// tries to calculate the cell containing that point.
//
// If the cell containing the point is not outside the bounds of the field,
// it returns (true, cell index) and
// (false, _) otherwise.
func (processor *processor) tryGetCell(point geom.Point) (bool, int) {
	column := int(math.Trunc(point.X() / processor.cellSize))
	row := int(math.Trunc(point.Y() / processor.cellSize))
	return 0 <= row &&
			row < processor.numRows &&
			0 <= column &&
			column < processor.numColumns,
		column*processor.numRows + row
}
