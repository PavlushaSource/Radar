package engine

import (
	"sync"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/geom"
)

// The processor type represents the engine processor and is used for processing while the engine is running.
type processor struct {
	// Current state
	state *State

	// Number of Dogs.
	numDogs int
	// Using geometry.
	geom geom.Geom
	// Fight radius.
	radiusFight float64
	// Hiss radius.
	radiusHiss float64

	// Cell size.
	cellSize float64

	// Async random used for generating random values from any async job.
	rndAsync rnd.RndAsync
	// coefficient used to parameterize the asynchronous random call.
	cf int

	// Number of columns in the field.
	numColumns int
	// Number of rows in the field.
	numRows int
	// Array of cells.
	//
	// cell[cellIdx] stores the indeces of Dogs in the cell with index cellIdx.
	cells [][]int

	// Point pool used to store Dogs coordinates between model state processing.
	points []geom.Point
}

// newProcessor creates new processor object by
// fight and hiss radiuses,
// number of Dogs,
// geometry and
// async random.
func newProcessor(radiusFight float64, radiusHiss float64, numDogs int, geometry geom.Geom, rndAsync rnd.RndAsync) *processor {
	var wg sync.WaitGroup

	processor := new(processor)

	processor.numDogs = numDogs
	processor.geom = geometry

	processor.points = make([]geom.Point, numDogs)

	for i := range processor.points {
		processor.points[i] = geometry.NewRandomPoint()
	}

	processor.radiusFight = radiusFight
	processor.radiusHiss = radiusHiss

	optimalCellSize := calculateOptimalCellSize(numDogs, geometry.Height()*geometry.Width())
	if radiusHiss < optimalCellSize {
		processor.cellSize = optimalCellSize
	} else {
		processor.cellSize = radiusHiss
	}

	processor.rndAsync = rndAsync
	processor.cf = 0

	processor.numColumns = numColumns(geometry.Width(), processor.cellSize)
	processor.numRows = numRows(geometry.Height(), processor.cellSize)
	numCells := processor.numColumns * processor.numRows
	processor.cells = make([][]int, numCells)
	for i := range processor.cells {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.cells[i] = make([]int, 0, numDogs)
		}()
	}

	wg.Wait()
	return processor
}

// process is a processor method that the given process state.
func (processor *processor) process(state *State) *State {
	processor.state = state
	processor.setUp()

	processor.moveDogs()
	processor.cellSplitting()
	processor.processDogsNeighbours()

	processor.tearDown()

	return processor.state
}

// setUp is a processor method
// used to configure the processor before processing state.
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
			processor.points[i].Copy(processor.state.Dogs[i])
		}()
	}
	wg.Wait()
}

// tearDown is a processor method
// used to tune the processor after state has been processed.
func (processor *processor) tearDown() {
	var wg sync.WaitGroup

	for i := range processor.points {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.state.Dogs[i].Copy(processor.points[i])
		}()
	}

	wg.Wait()
}

// moveDogs is a processor method that changes the coordinates of Dogs in model state.
func (processor *processor) moveDogs() {
	var wg sync.WaitGroup

	for i := range processor.state.Dogs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.geom.MovePoint(processor.state.Dogs[i])
		}()
	}

	wg.Wait()
}

// moveDogs is a processor method that partitions Dogs into cells based on their coordinates.
func (processor *processor) cellSplitting() {
	for i := range processor.state.Dogs {
		if success, cell := processor.tryGetCell(processor.state.Dogs[i]); success {
			processor.cells[cell] = append(processor.cells[cell], i)
		}
	}
}

// processDogsNeighbors is a processor method
// that calculates the new statuses of Dogs based on their interactions with neighbors.
func (processor *processor) processDogsNeighbours() {
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

	for i := range processor.state.Dogs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processor.updateDogstatus(i)
		}()
	}

	wg.Wait()
}

// processCellWithSelf is a processor method
// that handles the interaction of Dogs in a cell
// based on the given column and row of the cell.
func (processor *processor) processCellWithSelf(col int, row int) {
	for _, ci := range processor.cells[processor.cellByRowColumn(row, col)] {
		for _, ni := range processor.cells[processor.cellByRowColumn(row, col)] {
			if ci == ni {
				continue
			}
			processor.proccessPair(ci, ni)
		}
	}
}

// processCellWithOther is a processor method
// that handles the interaction of Dogs from cell with Dogs from another cell
// based on the given columns and rows of the cells.
func (processor *processor) processCellWithOther(col int, row int, otherCol int, otherRow int) {
	for _, ci := range processor.cells[processor.cellByRowColumn(row, col)] {
		for _, ni := range processor.cells[processor.cellByRowColumn(otherRow, otherCol)] {
			processor.proccessPair(ci, ni)
		}
	}
}

// proccessPair is a processor method
// that handles the interation a dogwith each other.
func (processor *processor) proccessPair(dogIdx int, neighbourIdx int) {
	dist := processor.geom.Distance(processor.state.Dogs[dogIdx], processor.state.Dogs[neighbourIdx])

	if dist <= processor.radiusFight {
		processor.state.Dogs[dogIdx].status = Fighting
		processor.state.Dogs[neighbourIdx].status = Fighting
	} else if dist <= processor.radiusHiss {
		if processor.rndAsync.Float64ByInt(dogIdx*neighbourIdx*processor.cf) <= (1 / (dist * dist)) {
			processor.state.Dogs[dogIdx].hissing = true
			processor.state.Dogs[neighbourIdx].hissing = true
		}
	}
}

// updateDogstatus is a processor method
// that attempts to set the given dogstatus to Hissing.
func (processor *processor) updateDogstatus(idx int) {
	if processor.state.Dogs[idx].hissing && (processor.state.Dogs[idx].status == Calm) {
		processor.state.Dogs[idx].status = Hissing
	}
}
