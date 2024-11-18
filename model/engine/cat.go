package engine

import (
	"github.com/PavlushaSource/Radar/model/geom"
)

type Status int

const (
	Calm Status = iota
	Hissing
	Fighting
)

type Cat struct {
	status  Status
	hissing bool
	geom.Point
}

func (cat *Cat) Status() Status {
	return cat.status
}

func (cat *Cat) clean() {
	cat.status = Calm
	cat.hissing = false
}

func newCat(point geom.Point) *Cat {
	cat := new(Cat)

	cat.status = Calm
	cat.hissing = false

	cat.Point = point

	return cat
}
