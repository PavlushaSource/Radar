package engine

import "math"

// Optimal cell size for cellSize1NumDogsBound or more Dogs.
const cellSize1 float64 = 1 << 2

// First number of Dogs bound.
const cellSize1NumDogsBound int = 1 << 18

// Optimal cell size for the number of Dogs is from cellSize2NumDogsBound to cellSize1NumDogsBound.
const cellSize2 float64 = 1 << 3

// Second number of Dogs bound.
const cellSize2NumDogsBound int = 1 << 16

// Optimal cell size for the number of Dogs is from cellSize1NumDogsBound to cellSize3NumDogsBound.
const cellSize3 float64 = 1 << 4

// Third number of Dogs bound.
const cellSize3NumDogsBound int = 1 << 13

// Optimal cell size for cellSize3NumDogsBound or less Dogs.
const cellSize4 float64 = 1 << 7

// Base square value.
const squareBase float64 = 1 << 20

// calculateOptimalCellSize calculates the optimal cell size for given number of Dogs and square value.
func calculateOptimalCellSize(numDogs int, square float64) float64 {
	optimalRadius := cellSize4
	if numDogs >= cellSize1NumDogsBound {
		optimalRadius = cellSize1
	} else if numDogs >= cellSize2NumDogsBound {
		optimalRadius = cellSize2
	} else if numDogs >= cellSize3NumDogsBound {
		optimalRadius = cellSize3
	}

	cellSizeCf := math.Sqrt(square / squareBase)

	if cellSizeCf == 0 {
		cellSizeCf = 1
	}

	return optimalRadius * cellSizeCf
}
