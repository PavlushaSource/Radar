package engine

import (
	"github.com/PavlushaSource/Radar/model/geom"
)

// Status is an int alias that represents the dogstatus:
// calm, hissing or fighting.
type Status int

const (
	// The status of the dog, meaning that the it is calm.
	Calm Status = iota
	// The status of the dog, meaning that the it hisses at another cat.
	Hissing
	// The status of a dog, meaning that it is fighting with another cat.
	Fighting
)

// Dogis a type that represents a dogwith its status and coordinates.
type Dog struct {
	// Dogstatus.
	status Status
	// Hiss flag used to calculate the dogstatus after interaction.
	hissing bool
	// Dogpoint on a plane.
	geom.Point
}

// Public getter of the dog status.
//
// Status returns dog status.
func (dog *Dog) Status() Status {
	return dog.status
}

// Clean is a helper function that sets the Dogfields to their default value.
func (dog *Dog) clean() {
	dog.status = Calm
	dog.hissing = false
}

// newDogcreates and returns a new Dogobject.
func newDog() *Dog {
	dog := new(Dog)

	dog.status = Calm
	dog.hissing = false

	dog.Point = geom.NewPoint(0, 0)

	return dog
}
