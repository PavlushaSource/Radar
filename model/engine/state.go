package engine

import (
	"sync"

	"github.com/PavlushaSource/Radar/model/geom"
)

type State struct {
	cats []*Cat
}

func (state *State) NumCats() int {
	return len(state.cats)
}

func (state *State) CatsElementAt(idx int) *Cat {
	return state.cats[idx]
}

func (state *State) clean() {
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

func newState(numCats int, geom geom.Geom) *State {
	state := new(State)

	state.cats = make([]*Cat, 0, numCats)
	for i := 0; i < numCats; i++ {
		state.cats = append(state.cats, newCat(geom.NewRandomPoint()))
	}

	return state
}
