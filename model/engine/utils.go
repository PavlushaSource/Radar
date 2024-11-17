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

func (engine *engine) catsPairIdx(first int64, second int64) int64 {
	if first <= second {
		return first*engine.state.NumCats() + second
	} else {
		return second*engine.state.NumCats() + first
	}
}
