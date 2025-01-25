package engine

import (
	"context"
	"fmt"
	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/geom"
	"testing"
	"time"
)

var inputPerformanceTest = []struct {
	winX, winY      float64
	maxMoveDistance float64
	dist            geom.Distance
	rFight, rHiss   float64
}{
	{1920, 1080, 300, geom.ManhattanDistance, 15, 30},
	{640, 640, 50, geom.EuclideanDistance, 8, 15},
	{1920, 1080, 75, geom.CurvilinearDistance, 16, 32},
}

func TestPerformance(t *testing.T) {

	for i, tt := range inputPerformanceTest {

		name := fmt.Sprintf("Test %d", i)
		t.Run(name, func(t *testing.T) {
			rndAsync := rnd.NewRndCore()
			g := geom.NewSimpleGeom(tt.winY, tt.winX, make([]geom.Barrier, 0), tt.maxMoveDistance, tt.dist, rndAsync)
			eng := NewEngine(tt.rFight, tt.rHiss, 50000, g, rndAsync, 5)

			getCh, putCh := eng.Run(context.Background())

			start := time.Now()
			iterationCount := 0
			for state := range getCh {
				end := time.Now()

				putCh <- state
				spend := end.Sub(start)
				if spend > 500*time.Millisecond {
					t.Errorf("spend too long (%v)", spend)
				}

				if iterationCount >= 5 {
					return
				}
				iterationCount++

				start = time.Now()
			}
		})
	}

}
