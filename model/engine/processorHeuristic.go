package engine

import "math"

const cellSize1 float64 = 1 << 2

const cellSize1NumCatsBound int = 1 << 18

const cellSize2 float64 = 1 << 3

const cellSize2NumCatsBound int = 1 << 16

const cellSize3 float64 = 1 << 4

const cellSize3NumCatsBound int = 1 << 13

const cellSize4 float64 = 1 << 7

const squareBase float64 = 1 << 20

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
