package engine

import (
	"sync"
)

// The state type, which represents the state of the model, contains a pool of Dogs.
type State struct {
	// Pool of Dogs.
	Dogs []*Dog
}

// NumDogs is a State method that returns the number of Dogs.
func (state *State) NumDogs() int {
	return len(state.Dogs)
}

// DogsElementAt is a State method that returns the dogby its index.
func (state *State) DogsElementAt(idx int) *Dog {
	return state.Dogs[idx]
}

// clean is a State method, which cleans each dogin this state.
func (state *State) clean() {
	var wg sync.WaitGroup

	for _, dog := range state.Dogs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dog.clean()
		}()
	}

	wg.Wait()
}

// newState creates a new model state by number of Dogs.
func newState(numDogs int) *State {
	state := new(State)

	state.Dogs = make([]*Dog, 0, numDogs)
	for i := 0; i < numDogs; i++ {
		state.Dogs = append(state.Dogs, newDog())
	}

	return state
}
