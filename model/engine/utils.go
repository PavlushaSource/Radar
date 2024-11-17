package engine

import (
	"math"
	"math/rand"
	"time"
)

func newRndCore() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func hissingProbability(dist float64) float64 {
	if dist <= 1.0 {
		return 1.0
	}
	return 1.0 - math.Sqrt(1.0-(1.0/(dist*dist)))
}
