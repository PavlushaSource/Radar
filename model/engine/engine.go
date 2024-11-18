package engine

import (
	"math/rand"
	"sync"

	"github.com/PavlushaSource/Radar/model/geom"
)

type engine struct {
	state *state

	geom   geom.Geom
	radius float64

	rndCores []*rand.Rand

	numColumns int
	numRows    int
	cells      [][]int

	hissings [][]bool
}

func (engine *engine) Run() State {
	engine.clean()

	engine.moveCats()
	engine.processCatsNeighbours()

	// return engine.state.Copy()
	return engine.state
}

func NewEngine(radiusFight float64, radiusHiss float64, numCats int, geom geom.Geom) Engine {
	var wg sync.WaitGroup

	engine := new(engine)

	engine.geom = geom
	engine.state = newState(geom.Height(), geom.Width(), radiusFight, radiusHiss, numCats, geom)
	engine.radius = radiusHiss

	engine.rndCores = make([]*rand.Rand, numCats)
	for i := range engine.rndCores {
		wg.Add(1)
		go func() {
			defer wg.Done()
			engine.rndCores[i] = newRndCore()
		}()
	}

	engine.numColumns = numColumns(geom.Width(), radiusHiss)
	engine.numRows = numRows(geom.Height(), radiusHiss)
	numCells := engine.numColumns * engine.numRows
	engine.cells = make([][]int, numCells)
	for i := range engine.cells {
		wg.Add(1)
		go func() {
			defer wg.Done()
			engine.cells[i] = make([]int, 0, numCats)
		}()
	}

	engine.hissings = make([][]bool, numCats)
	for i := range engine.hissings {
		wg.Add(1)
		go func() {
			defer wg.Done()
			engine.hissings[i] = make([]bool, numCats)
		}()
	}

	wg.Wait()
	return engine
}

func (engine *engine) clean() {
	var wg sync.WaitGroup

	for i := range engine.cells {
		wg.Add(1)
		go func() {
			defer wg.Done()
			engine.cells[i] = engine.cells[i][:0]
		}()
	}

	engine.hissings = make([][]bool, engine.state.NumCats())
	for i := range engine.hissings {
		wg.Add(1)
		go func() {
			defer wg.Done()
			engine.hissings[i] = make([]bool, engine.state.NumCats())
		}()
	}

	// runBlocking(
	// 	&engine.hissings,
	// 	func(i int, _ []bool) {
	// 		for j := 0; j < engine.state.NumCats(); j++ {
	// 			engine.hissings[i][j] = false
	// 		}
	// 	})

	engine.state.clean()

	wg.Wait()
}
