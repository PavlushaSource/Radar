package engine

import (
	"context"
	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/geom"
	"testing"
	"time"
)

func TestPerformance(t *testing.T) {
	rndAsync := rnd.NewRndCore()
	g := geom.NewSimpleGeom(1080, 1920, make([]geom.Barrier, 0), 100, geom.EuclideanDistance, rndAsync)
	eng := NewEngine(15, 30, 50000, g, rndAsync, 5)

	getCh, putCh := eng.Run(context.Background())

	start := time.Now()
	iterationCount := 0
	for state := range getCh {
		end := time.Now()

		putCh <- state
		spend := end.Sub(start)
		if spend > 250*time.Millisecond {
			t.Errorf("spend too long (%v)", spend)
		}

		if iterationCount >= 5 {
			return
		}
		iterationCount++

		start = time.Now()
	}
}
