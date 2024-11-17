package engine

import "sync"

func (engine *engine) moveCats() {
	for i, cat := range engine.state.Cats() {
		engine.geom.MovePoint(cat.Point())

		cell := engine.cell(cat.Point())
		engine.cells[cell] = append(engine.cells[cell], int64(i))
	}
}

func (engine *engine) processCatsNeighbours() {
	var wg sync.WaitGroup

	for i := range engine.state.Cats() {
		wg.Add(1)

		go func() {
			defer wg.Done()
			engine.processCatNeighbours(int64(i))
		}()
	}

	wg.Wait()

	for i := range engine.state.Cats() {
		wg.Add(1)

		go func() {
			defer wg.Done()
			engine.postprocessCatHissings(int64(i))
		}()
	}

	wg.Wait()

	engine.updateCatsStatus()
}

func (engine *engine) processCatNeighbours(idx int64) {
	cell := engine.cell(engine.state.Cats()[idx].Point())

	engine.processCell(idx, cell)
	engine.processNeighbourCell(idx, cell, engine.tryGetUpCell)
	engine.processNeighbourCell(idx, cell, engine.tryGetDownCell)
	engine.processNeighbourCell(idx, cell, engine.tryGetLeftCell)
	engine.processNeighbourCell(idx, cell, engine.tryGetRightCell)
	engine.processNeighbourCell(idx, cell, engine.tryGetUpLeftCell)
	engine.processNeighbourCell(idx, cell, engine.tryGetUpRightCell)
	engine.processNeighbourCell(idx, cell, engine.tryGetDownLeftCell)
	engine.processNeighbourCell(idx, cell, engine.tryGetDownRightCell)
}

func (engine *engine) processNeighbourCell(cat int64, cell int64, tryGetNeighbourCell neighbourCellExtractor) {
	if success, neighbourCell := tryGetNeighbourCell(cell); success {
		engine.processCell(cat, neighbourCell)
	}
}

func (engine *engine) processCell(cat int64, cell int64) {
	for _, neighbour := range engine.cells[cell] {
		engine.proccessPair(cat, neighbour)
	}
}

func (engine *engine) proccessPair(cat int64, neighbour int64) {
	cats := engine.state.Cats()
	dist := engine.geom.Distance(cats[cat].Point(), cats[neighbour].Point())
	if dist <= engine.state.RadiusFight() {
		cats[cat].setStatus(Fighting)
		cats[neighbour].setStatus(Fighting)

		// cats[cat].setFightings(append(cats[cat].Fightings(), neighbour))
	} else if dist <= engine.state.RadiusHiss() {
		if engine.rndCores[cat].Float64() <= hissingProbability(dist) {
			cats[cat].setHissing(true)
			cats[neighbour].setHissing(true)

			// engine.hissings[cat][neighbour] = true
		}

		// cats[cat].setHissings(append(cats[cat].Hissings(), neighbour))
	}
}

func (engine *engine) postprocessCatHissings(idx int64) {
	return
	cat := engine.state.Cats()[idx]
	if !cat.isHissing() {
		return
	}

	hissings := make([]int64, 0, len(cat.Hissings()))
	for _, neighbour := range cat.Hissings() {
		if engine.hissings[idx][neighbour] || engine.hissings[neighbour][idx] {
			hissings = append(hissings, neighbour)
		}
	}

	engine.state.Cats()[idx].setHissings(hissings)
}

func (engine *engine) updateCatsStatus() {
	for _, cat := range engine.state.Cats() {
		if cat.isHissing() && (cat.Status() == Calm) {
			cat.setStatus(Hissing)
		}
	}
}
