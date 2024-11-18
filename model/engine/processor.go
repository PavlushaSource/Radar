package engine

import (
	"sync"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/geom"
)

type processor struct {
	state *State

	numCats     int
	geom        geom.Geom
	radius      float64
	radiusFight float64
	radiusHiss  float64

	rndAsync rnd.RndAsync
	cf       int

	numColumns int
	numRows    int
	cells      [][]int

	points []geom.Point
}

func newProcessor(radiusFight float64, radiusHiss float64, numCats int, geometry geom.Geom, rndAsync rnd.RndAsync) *processor {
	var wg sync.WaitGroup

	processor := new(processor)

	processor.numCats = numCats
	processor.geom = geometry

	processor.points = make([]geom.Point, numCats)

	for i := range processor.points {
		processor.points[i] = geometry.NewPoint()
	}

	processor.radiusFight = radiusFight
	processor.radiusHiss = radiusHiss
	processor.radius = radiusHiss

	processor.rndAsync = rndAsync
	processor.cf = 0

	processor.numColumns = numColumns(geometry.Width(), radiusHiss)
	processor.numRows = numRows(geometry.Height(), radiusHiss)
	numCells := processor.numColumns * processor.numRows
	processor.cells = make([][]int, numCells)
	for i := range processor.cells {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.cells[i] = make([]int, 0, numCats)
		}()
	}

	wg.Wait()
	return processor
}

func (processor *processor) process(state *State) *State {
	processor.state = state
	processor.setUp()

	processor.moveCats()
	processor.cellSplitting()
	processor.processCatsNeighbours()

	processor.tearDown()

	return processor.state
}

func (processor *processor) setUp() {
	var wg sync.WaitGroup

	processor.cf += 1
	processor.state.clean()

	for i := range processor.cells {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.cells[i] = processor.cells[i][:0]
		}()
	}

	for i := range processor.points {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.points[i].Copy(processor.state.cats[i])
		}()
	}
	wg.Wait()
}

func (processor *processor) tearDown() {
	var wg sync.WaitGroup

	for i := range processor.points {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.state.cats[i].Copy(processor.points[i])
		}()
	}

	wg.Wait()
}

func (processor *processor) moveCats() {
	var wg sync.WaitGroup

	for i := range processor.state.cats {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.geom.MovePoint(processor.state.cats[i])
		}()
	}

	wg.Wait()
}

func (processor *processor) cellSplitting() {
	for i := range processor.state.cats {
		if success, cell := processor.tryGetCell(processor.state.cats[i]); success {
			processor.cells[cell] = append(processor.cells[cell], i)
		}
	}
}

func (processor *processor) processCatsNeighbours() {
	var wg sync.WaitGroup

	for i := range processor.state.cats {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.processCatNeighbours(i)
		}()
	}

	wg.Wait()

	for i := range processor.state.cats {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.updateCatStatus(i)
		}()
	}

	wg.Wait()
}

func (processor *processor) processCatNeighbours(catIdx int) {
	success, cell := processor.tryGetCell(processor.state.cats[catIdx])
	if !success {
		return
	}

	processor.processCell(catIdx, cell)
	processor.processNeighbourCell(catIdx, cell, processor.tryGetUpCell)
	processor.processNeighbourCell(catIdx, cell, processor.tryGetDownCell)
	processor.processNeighbourCell(catIdx, cell, processor.tryGetLeftCell)
	processor.processNeighbourCell(catIdx, cell, processor.tryGetRightCell)
	processor.processNeighbourCell(catIdx, cell, processor.tryGetUpLeftCell)
	processor.processNeighbourCell(catIdx, cell, processor.tryGetUpRightCell)
	processor.processNeighbourCell(catIdx, cell, processor.tryGetDownLeftCell)
	processor.processNeighbourCell(catIdx, cell, processor.tryGetDownRightCell)
}

func (processor *processor) processNeighbourCell(catIdx int, cell int, tryGetNeighbourCell neighbourCellExtractor) {
	if success, neighbourCell := tryGetNeighbourCell(cell); success {
		processor.processCell(catIdx, neighbourCell)
	}
}

func (processor *processor) processCell(catIdx int, cell int) {
	for _, neighbour := range processor.cells[cell] {
		if catIdx == neighbour {
			continue
		}
		processor.proccessPair(catIdx, neighbour)
	}
}

func (processor *processor) proccessPair(catIdx int, neighbourIdx int) {
	dist := processor.geom.Distance(processor.state.cats[catIdx], processor.state.cats[neighbourIdx])

	if dist <= processor.radiusFight {
		// if false {
		processor.state.cats[catIdx].status = Fighting
		processor.state.cats[neighbourIdx].status = Fighting
	} else if dist <= processor.radiusHiss {
		if processor.rndAsync.Float64ByInt(catIdx*neighbourIdx*processor.cf) <= hissingProbability(dist) {
			processor.state.cats[catIdx].hissing = true
			processor.state.cats[neighbourIdx].hissing = true
		}
	}
}

func (processor *processor) updateCatStatus(idx int) {
	if processor.state.cats[idx].hissing && (processor.state.cats[idx].status == Calm) {
		processor.state.cats[idx].status = Hissing
	}
}
