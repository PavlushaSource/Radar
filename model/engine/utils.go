package engine

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

type action[T any] func(int, T)

func runBlocking[T any](slice *[]T, action action[T]) {
	var wg sync.WaitGroup

	for idx, item := range *slice {
		wg.Add(1)

		go func() {
			defer wg.Done()

			action(idx, item)
		}()
	}

	wg.Wait()
}

func newRndCore() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func hissingProbability(dist float64) float64 {
	if dist <= 1.0 {
		return 1.0
	}
	return 1.0 - math.Sqrt(1.0-(1.0/(dist*dist)))
}
