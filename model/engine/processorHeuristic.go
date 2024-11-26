package engine

import "math"

// Optimal cell size for cellSize1NumCatsBound or more cats.
const cellSize1 float64 = 1 << 2

// First number of cats bound.
const cellSize1NumCatsBound int = 1 << 18

// Optimal cell size for the number of cats is from cellSize2NumCatsBound to cellSize1NumCatsBound.
const cellSize2 float64 = 1 << 3

// Second number of cats bound.
const cellSize2NumCatsBound int = 1 << 16

// Optimal cell size for the number of cats is from cellSize1NumCatsBound to cellSize3NumCatsBound.
const cellSize3 float64 = 1 << 4

// Third number of cats bound.
const cellSize3NumCatsBound int = 1 << 13

// Optimal cell size for cellSize3NumCatsBound or less cats.
const cellSize4 float64 = 1 << 7

// Base square value.
const squareBase float64 = 1 << 20

// calculateOptimalCellSize calculates the optimal cell size for given number of cats and square value.
func calculateOptimalCellSize(numCats int, square float64) float64 {
	optimalRadius := cellSize4
	if numCats >= cellSize1NumCatsBound {
		optimalRadius = cellSize1
	} else if numCats >= cellSize2NumCatsBound {
		optimalRadius = cellSize2
	} else if numCats >= cellSize3NumCatsBound {
		optimalRadius = cellSize3
	}

	cellSizeCf := math.Sqrt(square / squareBase)

	if cellSizeCf == 0 {
		cellSizeCf = 1
	}

	return optimalRadius * cellSizeCf
}
