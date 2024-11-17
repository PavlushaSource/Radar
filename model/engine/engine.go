package engine

import (
	"math/rand"

	"github.com/PavlushaSource/Radar/model/geom"
)

type Engine interface {
	Run() State
}

type engine struct {
	state  State
	geom   geom.Geom
	radius float64

	rndCore *rand.Rand

	numColumns int64
	numRows    int64
	cells      [][]int64

	hissings map[int64]bool
}

// TODO: add process logic
func (engine *engine) Run() State {
	engine.clean()

	// engine.moveCats()
	// engine.processCatsNeighbours()

	return engine.state.Copy()
}

func NewEngine(radiusFight float64, radiusHiss float64, numCats int64, geom geom.Geom) Engine {
	engine := new(engine)

	engine.geom = geom
	engine.state = newState(geom.Height(), geom.Width(), radiusFight, radiusHiss, numCats, geom)
	engine.radius = radiusHiss

	engine.rndCore = newRndCore()

	engine.numColumns = numColumns(geom.Width(), radiusHiss)
	engine.numRows = numRows(geom.Height(), radiusHiss)
	numCells := engine.numColumns * engine.numRows
	engine.cells = make([][]int64, 0, numCells)

	for i := int64(0); i < numCells; i++ {
		engine.cells = append(engine.cells, make([]int64, 0, numCats))
	}

	engine.hissings = make(map[int64]bool, numCats*numCats)

	return engine
}

func (engine *engine) clean() {
	for i := range engine.cells {
		capacity := cap(engine.cells[i])
		engine.cells[i] = make([]int64, 0, capacity)
	}

	engine.hissings = make(map[int64]bool, engine.state.NumCats()*engine.state.NumCats())

	engine.state.clean()
}
