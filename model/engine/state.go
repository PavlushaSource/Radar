package engine

import (
	"sync"

	"github.com/PavlushaSource/Radar/model/geom"
)

type State interface {
	Height() float64
	Width() float64
	RadiusHiss() float64
	RadiusFight() float64

	NumCats() int
	Cats() []Cat

	Copy(dst State)

	setHeight(height float64)
	setWidth(width float64)
	setRadiusHiss(radiusHiss float64)
	setRadiusFight(radiusFight float64)

	copyCatsFrom(src []Cat)

	clean()
}

type state struct {
	height      float64
	width       float64
	radiusHiss  float64
	radiusFight float64
	cats        []Cat
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

func (state *state) Cats() []Cat {
	return state.cats
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

func (state *state) copyCatsFrom(src []Cat) {
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
	for _, cat := range state.cats {
		cat.clean()
	}
}

func newState(height float64, width float64, radiusFight float64, radiusHiss float64, numCats int64, geom geom.Geom) State {
	state := new(state)

	state.height = height
	state.width = width
	state.radiusFight = radiusFight
	state.radiusHiss = radiusHiss

	state.cats = make([]Cat, 0, numCats)
	for i := int64(0); i < numCats; i++ {
		state.cats = append(state.cats, newCat(geom.NewPoint(), numCats))
	}

	return state
}
