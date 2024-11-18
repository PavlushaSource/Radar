package engine

import (
	"github.com/PavlushaSource/Radar/model/geom"
)

type cat struct {
	status  Status
	hissing bool
	geom.Point
}

func (cat *cat) Status() Status {
	return cat.status
}

func (cat *cat) copy(dst *cat) {
	dst.status = cat.status
	dst.hissing = cat.hissing

	cat.Point.Copy(dst.Point)
}

func (cat *cat) setStatus(status Status) {
	cat.status = status
}

func (cat *cat) setHissing(hissing bool) {
	cat.hissing = hissing
}

func (cat *cat) clean() {
	cat.status = Calm
	cat.hissing = false
}

func newCat(point geom.Point) *cat {
	cat := new(cat)

	cat.status = Calm
	cat.hissing = false

	cat.Point = point

	return cat
}
