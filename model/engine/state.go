package engine

import (
	"sync"
)

// The state type, which represents the state of the model, contains a pool of cats.
type State struct {
	cats []*Cat
}

// NumCats is a State method that returns the number of cats.
func (state *State) NumCats() int {
	return len(state.cats)
}

// CatsElementAt is a State method that returns the cat by its index.
func (state *State) CatsElementAt(idx int) *Cat {
	return state.cats[idx]
}

// clean is a State method, which cleans each cat in this state.
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

// newState creates a new model state by number of cats.
func newState(numCats int) *State {
	state := new(State)

	state.cats = make([]*Cat, 0, numCats)
	for i := 0; i < numCats; i++ {
		state.cats = append(state.cats, newCat())
	}

	return state
}
