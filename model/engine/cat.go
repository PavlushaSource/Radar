package engine

import (
	"github.com/PavlushaSource/Radar/model/geom"
)

type cat struct {
	status  Status
	hissing bool
	_point  geom.Point
}

func (cat *cat) Status() Status {
	return cat.status
}

func (cat *cat) X() float64 {
	return cat._point.X()
}

func (cat *cat) Y() float64 {
	return cat._point.Y()
}

func (cat *cat) Copy(dst Cat) {
	dst.setStatus(cat.status)
	dst.setHissing(cat.hissing)

	cat._point.Copy(dst.point())
}

func (cat *cat) point() geom.Point {
	return cat._point
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

	cat._point = point

	return cat
}
