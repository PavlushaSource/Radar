package engine

import "math"

const radius1 float64 = 1 << 2

const radius1NumCatsBound int = 1 << 18

const radius2 float64 = 1 << 3

const radius2NumCatsBound int = 1 << 16

const radius3 float64 = 1 << 4

const radius3NumCatsBound int = 1 << 13

const radius4 float64 = 1 << 7

const squareBase float64 = 1 << 20

func calculateOptimalRadius(numCats int, square float64) float64 {
	optimalRadius := radius4
	if numCats >= radius1NumCatsBound {
		optimalRadius = radius1
	} else if numCats >= radius2NumCatsBound {
		optimalRadius = radius2
	} else if numCats >= radius3NumCatsBound {
		optimalRadius = radius3
	}

	radiusCf := math.Sqrt(square / squareBase)

	if radiusCf == 0 {
		radiusCf = 1
	}

	return optimalRadius * radiusCf
}
