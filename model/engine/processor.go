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
	radiusFight float64
	radiusHiss  float64

	radius float64

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

	optimalRadius := calculateOptimalRadius(numCats, geometry.Height()*geometry.Width())
	if radiusHiss < optimalRadius {
		processor.radius = optimalRadius
	} else {
		processor.radius = radiusHiss
	}

	processor.rndAsync = rndAsync
	processor.cf = 0

	processor.numColumns = numColumns(geometry.Width(), processor.radius)
	processor.numRows = numRows(geometry.Height(), processor.radius)
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

	for col := 0; col < processor.numColumns-1; col++ {
		for row := 0; row < processor.numRows-1; row++ {
			wg.Add(4)
			go func() {
				defer wg.Done()
				processor.processCellWithSelf(col, row)
			}()
			go func() {
				defer wg.Done()
				processor.processCellWithOther(col, row, col+1, row)
			}()
			go func() {
				defer wg.Done()
				processor.processCellWithOther(col, row, col, row+1)
			}()
			go func() {
				defer wg.Done()
				processor.processCellWithOther(col, row, col+1, row+1)
			}()
		}
	}

	for row := 0; row < processor.numRows-1; row++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			processor.processCellWithSelf(processor.numColumns-1, row)
		}()
		go func() {
			defer wg.Done()
			processor.processCellWithOther(processor.numColumns-1, row, processor.numColumns-1, row+1)
		}()
	}

	for col := 0; col < processor.numColumns-1; col++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			processor.processCellWithSelf(col, processor.numRows-1)
		}()
		go func() {
			defer wg.Done()
			processor.processCellWithOther(col, processor.numRows-1, col+1, processor.numRows-1)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		processor.processCellWithSelf(processor.numColumns-1, processor.numRows-1)
	}()

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

func (processor *processor) processCellWithSelf(col int, row int) {
	for _, ci := range processor.cells[processor.cellByColumnRow(col, row)] {
		for _, ni := range processor.cells[processor.cellByColumnRow(col, row)] {
			if ci == ni {
				continue
			}
			processor.proccessPair(ci, ni)
		}
	}
}

func (processor *processor) processCellWithOther(col int, row int, otherCol int, otherRow int) {
	for _, ci := range processor.cells[processor.cellByColumnRow(col, row)] {
		for _, ni := range processor.cells[processor.cellByColumnRow(otherCol, otherRow)] {
			processor.proccessPair(ci, ni)
		}
	}
}

func (processor *processor) proccessPair(catIdx int, neighbourIdx int) {
	dist := processor.geom.Distance(processor.state.cats[catIdx], processor.state.cats[neighbourIdx])

	if dist <= processor.radiusFight {
		processor.state.cats[catIdx].status = Fighting
		processor.state.cats[neighbourIdx].status = Fighting
	} else if dist <= processor.radiusHiss {
		if processor.rndAsync.Float64ByInt(catIdx*neighbourIdx*processor.cf) <= (1 / (dist * dist)) {
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
