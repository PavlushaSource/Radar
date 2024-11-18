package engine

import (
	"math/rand"
	"sync"

	"github.com/PavlushaSource/Radar/model/geom"
)

type processor struct {
	state *state

	geom   geom.Geom
	radius float64

	rndCores []*rand.Rand

	numColumns int
	numRows    int
	cells      [][]int

	points []geom.Point

	hissings [][]bool
}

func newProcessor(radius float64, numCats int, geometry geom.Geom) *processor {
	var wg sync.WaitGroup

	processor := new(processor)

	processor.geom = geometry

	processor.points = make([]geom.Point, numCats)

	for i := range processor.points {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.points[i] = geometry.NewPoint()
		}()
	}

	processor.radius = radius

	processor.rndCores = make([]*rand.Rand, numCats)
	for i := range processor.rndCores {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.rndCores[i] = newRndCore()
		}()
	}

	processor.numColumns = numColumns(geometry.Width(), radius)
	processor.numRows = numRows(geometry.Height(), radius)
	numCells := processor.numColumns * processor.numRows
	processor.cells = make([][]int, numCells)
	for i := range processor.cells {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.cells[i] = make([]int, 0, numCats)
		}()
	}

	processor.hissings = make([][]bool, numCats)
	for i := range processor.hissings {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.hissings[i] = make([]bool, numCats)
		}()
	}

	wg.Wait()
	return processor
}

func (processor *processor) process(state *state) *state {
	processor.state = state
	processor.setUp()

	processor.moveCats()
	processor.processCatsNeighbours()

	processor.tearDown()

	return processor.state
}

func (processor *processor) setUp() {
	var wg sync.WaitGroup

	for i := range processor.cells {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.cells[i] = processor.cells[i][:0]
		}()
	}

	processor.hissings = make([][]bool, processor.state.NumCats())
	for i := range processor.hissings {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.hissings[i] = make([]bool, processor.state.NumCats())
		}()
	}

	// runBlocking(
	// 	&processor.hissings,
	// 	func(i int, _ []bool) {
	// 		for j := 0; j < processor.state.NumCats(); j++ {
	// 			processor.hissings[i][j] = false
	// 		}
	// 	})

	processor.state.clean()

	for i := range processor.points {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.points[i].Copy(processor.state.cats[i]._point)
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
			processor.state.cats[i]._point.Copy(processor.points[i])
		}()
	}

	wg.Wait()
}

func (processor *processor) moveCats() {
	var wg sync.WaitGroup

	for _, cat := range processor.state.cats {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.geom.MovePoint(cat._point)
		}()
	}

	wg.Wait()

	for i, cat := range processor.state.cats {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cell := processor.cell(cat._point)
			processor.cells[cell] = append(processor.cells[cell], i)
		}()
	}

	wg.Wait()
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

	for i, cat := range processor.state.cats {
		wg.Add(2)
		go func() {
			defer wg.Done()
			processor.postprocessCatHissings(i)
		}()

		go func() {
			defer wg.Done()
			processor.updateCatStatus(cat)
		}()
	}

	wg.Wait()
}

func (processor *processor) processCatNeighbours(catIdx int) {
	cat := processor.state.cats[catIdx]
	cell := processor.cell(cat._point)

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
		processor.proccessPair(catIdx, neighbour)
	}
}

func (processor *processor) proccessPair(catIdx int, neighbourIdx int) {
	cat := processor.state.cats[catIdx]
	neighbour := processor.state.cats[neighbourIdx]
	dist := processor.geom.Distance(cat.point(), neighbour.point())

	if dist <= processor.state.radiusFight {
		cat.setStatus(Fighting)
		neighbour.setStatus(Fighting)

		cat.fightings = append(cat.fightings, neighbourIdx)
	} else if dist <= processor.state.radiusHiss {
		if processor.rndCores[catIdx].Float64() <= hissingProbability(dist) {
			cat.hissing = true
			neighbour.hissing = true

			processor.hissings[catIdx][neighbourIdx] = true
		}

		cat.hissings = append(cat.hissings, neighbourIdx)
	}
}

func (processor *processor) postprocessCatHissings(catIdx int) {
	cat := processor.state.cats[catIdx]
	if !cat.hissing {
		return
	}

	n := 0
	for _, neighbour := range cat.hissings {
		if processor.hissings[catIdx][neighbour] || processor.hissings[neighbour][catIdx] {
			cat.hissings[n] = neighbour
			n++
		}
	}
	cat.hissings = cat.hissings[:n]
}

func (processor *processor) updateCatStatus(cat *cat) {
	if cat.hissing && (cat.status == Calm) {
		cat.status = Hissing
	}
}
