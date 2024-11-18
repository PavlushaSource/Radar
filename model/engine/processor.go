package engine

import "sync"

func (engine *engine) moveCats() {
	var wg sync.WaitGroup

	for _, cat := range engine.state.cats {
		wg.Add(1)
		go func() {
			defer wg.Done()
			engine.geom.MovePoint(cat._point)
		}()
	}

	wg.Wait()

	for i, cat := range engine.state.cats {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cell := engine.cell(cat._point)
			engine.cells[cell] = append(engine.cells[cell], i)
		}()
	}

	wg.Wait()
}

func (engine *engine) processCatsNeighbours() {
	var wg sync.WaitGroup

	for i := range engine.state.cats {
		wg.Add(1)
		go func() {
			defer wg.Done()
			engine.processCatNeighbours(i)
		}()
	}

	wg.Wait()

	for i, cat := range engine.state.cats {
		wg.Add(2)
		go func() {
			defer wg.Done()
			engine.postprocessCatHissings(i)
		}()

		go func() {
			defer wg.Done()
			engine.updateCatStatus(cat)
		}()
	}

	wg.Wait()
}

func (engine *engine) processCatNeighbours(catIdx int) {
	cat := engine.state.cats[catIdx]
	cell := engine.cell(cat._point)

	engine.processCell(catIdx, cell)
	engine.processNeighbourCell(catIdx, cell, engine.tryGetUpCell)
	engine.processNeighbourCell(catIdx, cell, engine.tryGetDownCell)
	engine.processNeighbourCell(catIdx, cell, engine.tryGetLeftCell)
	engine.processNeighbourCell(catIdx, cell, engine.tryGetRightCell)
	engine.processNeighbourCell(catIdx, cell, engine.tryGetUpLeftCell)
	engine.processNeighbourCell(catIdx, cell, engine.tryGetUpRightCell)
	engine.processNeighbourCell(catIdx, cell, engine.tryGetDownLeftCell)
	engine.processNeighbourCell(catIdx, cell, engine.tryGetDownRightCell)
}

func (engine *engine) processNeighbourCell(catIdx int, cell int, tryGetNeighbourCell neighbourCellExtractor) {
	if success, neighbourCell := tryGetNeighbourCell(cell); success {
		engine.processCell(catIdx, neighbourCell)
	}
}

func (engine *engine) processCell(catIdx int, cell int) {
	for _, neighbour := range engine.cells[cell] {
		engine.proccessPair(catIdx, neighbour)
	}
}

func (engine *engine) proccessPair(catIdx int, neighbourIdx int) {
	cat := engine.state.cats[catIdx]
	neighbour := engine.state.cats[neighbourIdx]
	dist := engine.geom.Distance(cat.point(), neighbour.point())

	if dist <= engine.state.radiusFight {
		cat.setStatus(Fighting)
		neighbour.setStatus(Fighting)

		cat.fightings = append(cat.fightings, neighbourIdx)
	} else if dist <= engine.state.radiusHiss {
		if engine.rndCores[catIdx].Float64() <= hissingProbability(dist) {
			cat.hissing = true
			neighbour.hissing = true

			engine.hissings[catIdx][neighbourIdx] = true
		}

		cat.hissings = append(cat.hissings, neighbourIdx)
	}
}

func (engine *engine) postprocessCatHissings(catIdx int) {
	cat := engine.state.cats[catIdx]
	if !cat.hissing {
		return
	}

	n := 0
	for _, neighbour := range cat.hissings {
		if engine.hissings[catIdx][neighbour] || engine.hissings[neighbour][catIdx] {
			cat.hissings[n] = neighbour
			n++
		}
	}
	cat.hissings = cat.hissings[:n]
}

func (engine *engine) updateCatStatus(cat *cat) {
	if cat.hissing && (cat.status == Calm) {
		cat.status = Hissing
	}
}
