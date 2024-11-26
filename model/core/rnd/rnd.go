package rnd

import (
	"math"
	"math/rand"
	"time"
)

// RndAsync interface represents asyn random.
type RndAsync interface {
	// Float64ByInt is a RndAsync method
	// that generate float64 random value by given int base.
	Float64ByInt(int) float64

	// Float64ByFloat64 is a RndAsync method
	// that generate float64 random value by given float64 base.
	Float64ByFloat64(float64) float64
}

// The maximum number of pre-calculated random values.
const num int = 5e6

// Prime integer used in calcuation the index of the random value.
const cf float64 = 1e9 + 7

// Implementation of RndAsync.
type rndAsync struct {
	// Pool of pre-calculated random values.
	precalc []float64
}

// Float64ByInt is a rndAsync method
// that generate float64 random value by given int base.
//
// Implementation of RndAsync.Float64ByInt.
func (rndAsync *rndAsync) Float64ByInt(x int) float64 {
	idx := int(math.Abs(float64(x)*cf)) % num
	return rndAsync.precalc[idx]
}

// Float64ByFloat64 is a RndAsync method
// that generate float64 random value by given float64 base.
//
// Implementation of RndAsync.Float64ByFloat64.
func (rndAsync *rndAsync) Float64ByFloat64(x float64) float64 {
	idx := int(math.Abs(x*cf)) % num
	return rndAsync.precalc[idx]
}

// NewRndCore creates a new object, that inherits the RndAsync interface.
//
// NewRndCore pre-calculates pool of random values.
func NewRndCore() RndAsync {
	rndAsync := new(rndAsync)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndAsync.precalc = make([]float64, 0, num)

	for i := 0; i < num; i++ {
		rndAsync.precalc = append(rndAsync.precalc, rnd.Float64())
	}

	return rndAsync
}
