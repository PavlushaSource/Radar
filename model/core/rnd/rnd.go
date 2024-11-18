package rnd

import (
	"math"
	"math/rand"
	"time"
)

type RndAsync interface {
	Float64ByInt(int) float64

	Float64ByFloat64(float64) float64
}

const num int = 5e6
const cf float64 = 1e9 + 7

type rndAsync struct {
	precalc []float64
}

func (rndAsync *rndAsync) Float64ByInt(x int) float64 {
	idx := int(math.Abs(float64(x)*cf)) % num
	return rndAsync.precalc[idx]
}

func (rndAsync *rndAsync) Float64ByFloat64(x float64) float64 {
	idx := int(math.Abs(x*cf)) % num
	return rndAsync.precalc[idx]
}

func NewRndCore() RndAsync {
	rndAsync := new(rndAsync)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndAsync.precalc = make([]float64, 0, num)

	for i := 0; i < num; i++ {
		rndAsync.precalc = append(rndAsync.precalc, rnd.Float64())
	}

	return rndAsync
}
