package engine

import (
	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/geom"
)

type engine struct {
	processor *processor
	state     *state
}

func (engine *engine) Run() State {

	engine.processor.process(engine.state)

	return engine.state
}

func NewEngine(radiusFight float64, radiusHiss float64, numCats int, geom geom.Geom, rndAsync rnd.RndAsync) Engine {
	engine := new(engine)

	engine.state = newState(geom.Height(), geom.Width(), radiusFight, radiusHiss, numCats, geom)

	engine.processor = newProcessor(radiusHiss, numCats, geom, rndAsync)

	return engine
}
