package engine

import (
	"github.com/PavlushaSource/Radar/model/geom"
)

// Status is an int alias that represents the cat status:
// calm, hissing or fighting.
type Status int

const (
	// The status of the cat, meaning that the it is calm.
	Calm Status = iota
	// The status of the cat, meaning that the it hisses at another cat.
	Hissing
	// The status of a cat, meaning that it is fighting with another cat.
	Fighting
)

// Cat is a type that represents a cat with its status and coordinates.
type Cat struct {
	// Cat status
	status Status
	// Hiss flag used to calculate the cat status after interaction.
	hissing bool
	// Cat point on a plane.
	geom.Point
}

// Public getter of the cat status.
//
// Status returns cat status.
func (cat *Cat) Status() Status {
	return cat.status
}

// Clean is a helper function that sets the Cat fields to their default value.
func (cat *Cat) clean() {
	cat.status = Calm
	cat.hissing = false
}

// newCat creates and returns a new Cat object.
func newCat() *Cat {
	cat := new(Cat)

	cat.status = Calm
	cat.hissing = false

	cat.Point = geom.NewPoint(0, 0)

	return cat
}
