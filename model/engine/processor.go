package engine

import (
	"sync"
)

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

	if success, neighbourCell := engine.tryGetUpCell(cell); success {
		engine.processCell(idx, neighbourCell)
	}
	if success, neighbourCell := engine.tryGetDownCell(cell); success {
		engine.processCell(idx, neighbourCell)
	}
	if success, neighbourCell := engine.tryGetLeftCell(cell); success {
		engine.processCell(idx, neighbourCell)
	}
	if success, neighbourCell := engine.tryGetRightCell(cell); success {
		engine.processCell(idx, neighbourCell)
	}
	if success, neighbourCell := engine.tryGetUpLeftCell(cell); success {
		engine.processCell(idx, neighbourCell)
	}
	if success, neighbourCell := engine.tryGetUpRighrCell(cell); success {
		engine.processCell(idx, neighbourCell)
	}
	if success, neighbourCell := engine.tryGetDownLeftCell(cell); success {
		engine.processCell(idx, neighbourCell)
	}
	if success, neighbourCell := engine.tryGetDownRightCell(cell); success {
		engine.processCell(idx, neighbourCell)
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

		cats[cat].setFightings(append(cats[cat].Fightings(), neighbour))
	} else if dist <= engine.state.RadiusHiss() {
		if engine.rndCore.Float64() >= hissingProbability(dist) {
			cats[cat].setHissing(true)
			cats[neighbour].setHissing(true)

			engine.hissings[engine.catsPairIdx(cat, neighbour)] = true
		}

		cats[cat].setHissings(append(cats[cat].Hissings(), neighbour))
	}
}

func (engine *engine) postprocessCatHissings(idx int64) {
	cat := engine.state.Cats()[idx]
	if !cat.isHissing() {
		return
	}

	hissings := make([]int64, 0, len(cat.Hissings()))
	for _, neighbour := range cat.Hissings() {
		if exist, success := engine.hissings[engine.catsPairIdx(idx, neighbour)]; exist && success {
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
