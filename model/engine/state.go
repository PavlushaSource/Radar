package engine

import (
	"github.com/PavlushaSource/Radar/model/geom"
)

type State interface {
	Height() float64
	Width() float64
	RadiusHiss() float64
	RadiusFight() float64

	NumCats() int
	CatsElementAt(idx int) Cat

	Copy(dst State)

	setHeight(height float64)
	setWidth(width float64)
	setRadiusHiss(radiusHiss float64)
	setRadiusFight(radiusFight float64)

	copyCatsFrom(src []Cat)

	catsRunBlocking(action[Cat])

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

func (state *state) CatsElementAt(idx int) Cat {
	return state.cats[idx]
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
	state.cats = state.cats[:len(src)]

	runBlocking(
		&state.cats,
		func(i int, _ Cat) {
			src[i].Copy(state.cats[i])
		})
}

func (state *state) catsRunBlocking(action action[Cat]) {
	runBlocking(
		&state.cats,
		func(i int, cat Cat) {
			action(i, cat)
		})
}

func (state *state) clean() {
	state.catsRunBlocking(
		func(_ int, cat Cat) {
			cat.clean()
		})
}

func newState(height float64, width float64, radiusFight float64, radiusHiss float64, numCats int, geom geom.Geom) State {
	state := new(state)

	state.height = height
	state.width = width
	state.radiusFight = radiusFight
	state.radiusHiss = radiusHiss

	state.cats = make([]Cat, 0, numCats)
	for i := 0; i < numCats; i++ {
		state.cats = append(state.cats, newCat(geom.NewPoint(), numCats))
	}

	return state
}
