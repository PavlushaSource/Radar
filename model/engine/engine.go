package engine

import (
	"math/rand"

	"github.com/PavlushaSource/Radar/model/geom"
)

type Engine interface {
	Run() State
}

type engine struct {
	state State

	geom   geom.Geom
	radius float64

	rndCores []*rand.Rand

	numColumns int64
	numRows    int64
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
	engine := new(engine)

	engine.geom = geom
	engine.state = newState(geom.Height(), geom.Width(), radiusFight, radiusHiss, numCats, geom)
	engine.radius = radiusHiss

	engine.rndCores = make([]*rand.Rand, 0, numCats)
	for i := 0; i < numCats; i++ {
		engine.rndCores = append(engine.rndCores, newRndCore())
	}

	engine.numColumns = numColumns(geom.Width(), radiusHiss)
	engine.numRows = numRows(geom.Height(), radiusHiss)
	numCells := engine.numColumns * engine.numRows
	engine.cells = make([][]int, 0, numCells)

	for i := int64(0); i < numCells; i++ {
		engine.cells = append(engine.cells, make([]int, 0, numCats))
	}

	engine.hissings = make([][]bool, 0, numCats)
	for i := 0; i < numCats; i++ {
		engine.hissings = append(engine.hissings, make([]bool, numCats))
	}

	return engine
}

func (engine *engine) clean() {
	for i := range engine.cells {
		capacity := cap(engine.cells[i])
		engine.cells[i] = make([]int, 0, capacity)
	}

	engine.hissings = make([][]bool, 0, engine.state.NumCats())
	for i := 0; i < engine.state.NumCats(); i++ {
		engine.hissings = append(engine.hissings, make([]bool, engine.state.NumCats()))
	}

	engine.state.clean()
}
