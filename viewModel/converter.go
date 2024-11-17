package viewModel

import (
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/view/domain"
)

type Converter interface {
	ConvertStateToCat(state engine.State) domain.CatUI
}

type converter struct {
}

func (c *converter) ConvertStateToCat(state engine.State) domain.CatUI {
	return domain.CatUI{}
}
