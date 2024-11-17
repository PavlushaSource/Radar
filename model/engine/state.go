package engine

import "github.com/PavlushaSource/Radar/model/geom"

type State interface {
	Height() float64
	Width() float64
	RadiusHiss() float64
	RadiusFight() float64
	NumCats() int64
	Cats() []Cat
	Copy() State
	clean()
}

type state struct {
	height      float64
	width       float64
	radiusHiss  float64
	radiusFight float64
	numCats     int64
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

func (state *state) NumCats() int64 {
	return state.numCats
}

func (state *state) Cats() []Cat {
	return state.cats
}

func (original *state) Copy() State {
	copy := new(state)

	copy.height = original.height
	copy.width = original.width
	copy.radiusFight = original.radiusFight
	copy.radiusHiss = original.radiusHiss
	copy.numCats = original.numCats

	copy.cats = make([]Cat, 0, copy.numCats)
	for _, cat := range original.cats {
		copy.cats = append(copy.cats, cat.Copy())
	}

	return copy
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
	state.numCats = numCats

	state.cats = make([]Cat, 0, numCats)
	for i := int64(0); i < numCats; i++ {
		state.cats = append(state.cats, newCat(geom.NewPoint(), numCats))
	}

	return state
}
