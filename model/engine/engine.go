package engine

import (
	"context"
	"fmt"
	"sync"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/geom"
)

type Engine struct {
	processor  *processor
	bufferSize int
}

func (engine *Engine) Run(ctx context.Context) (chan *State, chan *State) {
	resCh := make(chan *State, engine.bufferSize)
	buffCh := make(chan *State, engine.bufferSize)

	var wg sync.WaitGroup
	for i := 0; i < engine.bufferSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("state creation")
			buffCh <- newState(engine.processor.numCats, engine.processor.geom)
		}()
	}
	wg.Wait()

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(resCh)
				return
			default:
				resCh <- engine.processor.process(<-buffCh)
			}
		}
	}()

	return resCh, buffCh
}

func NewEngine(radiusFight float64, radiusHiss float64, numCats int, geom geom.Geom, rndAsync rnd.RndAsync, bufferSize int) *Engine {
	engine := new(Engine)

	engine.processor = newProcessor(radiusFight, radiusHiss, numCats, geom, rndAsync)

	engine.bufferSize = bufferSize

	return engine
}
