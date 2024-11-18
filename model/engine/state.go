package engine

import (
	"sync"

	"github.com/PavlushaSource/Radar/model/geom"
)

type state struct {
	height      float64
	width       float64
	radiusHiss  float64
	radiusFight float64
	cats        []*cat
}

func (state *state) Height() float64 {
	return state.height
}

func (state *state) Width() float64 {
	return state.width
}

func (state *state) RadiusHiss() float64 {
	return state.radiusHiss
}

func (state *state) RadiusFight() float64 {
	return state.radiusFight
}

func (state *state) NumCats() int {
	return len(state.cats)
}

func (state *state) CatsElementAt(idx int) Cat {
	return state.cats[idx]
}

func (state *state) Copy(dst State) {
	dst.setHeight(state.height)
	dst.setWidth(state.width)
	dst.setRadiusHiss(state.radiusHiss)
	dst.setRadiusFight(state.radiusFight)

	dst.copyCatsFrom(state.cats)
}

func (state *state) setHeight(height float64) {
	state.height = height
}

func (state *state) setWidth(width float64) {
	state.width = width
}

func (state *state) setRadiusHiss(radiusHiss float64) {
	state.radiusHiss = radiusHiss
}

func (state *state) setRadiusFight(radiusFight float64) {
	state.radiusFight = radiusFight
}

func (state *state) copyCatsFrom(src []*cat) {
	var wg sync.WaitGroup

	state.cats = state.cats[:len(src)]

	for i := range src {
		wg.Add(1)
		go func() {
			defer wg.Done()
			src[i].Copy(state.cats[i])
		}()
	}

	wg.Wait()
}

func (state *state) clean() {
	var wg sync.WaitGroup

	for _, cat := range state.cats {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cat.clean()
		}()
	}

	wg.Wait()
}

func newState(height float64, width float64, radiusFight float64, radiusHiss float64, numCats int, geom geom.Geom) *state {
	state := new(state)

	state.height = height
	state.width = width
	state.radiusFight = radiusFight
	state.radiusHiss = radiusHiss

	state.cats = make([]*cat, 0, numCats)
	for i := 0; i < numCats; i++ {
		state.cats = append(state.cats, newCat(geom.NewPoint()))
	}

	return state
}
