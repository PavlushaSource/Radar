package engine

import (
	"context"
	"sync"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/geom"
)

// Engine type represents radar engine.
//
// Creation: NewEngine(params)
//
// Starting: (*Engine).Run(params)
type Engine struct {
	// The engine processor is used to process the state of the model.
	//
	// One for each engine.
	processor *processor

	// bufferSize is the size of the state buffer that is used while the engine is running.
	bufferSize int
}

// Run is the Engine method that runs it.
//
// Run receives a context (context.Context) and
// returns a pair of buffered state chanels:
// (getting chan, putting chan).
//
// The getting channel is the channel for getting the state processed by engine.
//
// The putting channel is the channel for putting the state after processing outside the engine.
//
// The engine reuses states from the putting channel while running.
//
// The channel buffer size is equal Engine.bufferSize.
func (engine *Engine) Run(ctx context.Context) (chan *State, chan *State) {
	resCh := make(chan *State, engine.bufferSize)
	buffCh := make(chan *State, engine.bufferSize)

	var wg sync.WaitGroup
	for i := 0; i < engine.bufferSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
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

// NewEngine creates new Engine object by
// fight and hiss radiuses,
// number of cats,
// geometry,
// async random and
// buffer size.
func NewEngine(radiusFight float64, radiusHiss float64, numCats int, geom geom.Geom, rndAsync rnd.RndAsync, bufferSize int) *Engine {
	engine := new(Engine)

	engine.processor = newProcessor(radiusFight, radiusHiss, numCats, geom, rndAsync)

	engine.bufferSize = bufferSize

	return engine
}
