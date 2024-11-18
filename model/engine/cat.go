package engine

import (
	"github.com/PavlushaSource/Radar/model/geom"
)

type cat struct {
	status    Status
	hissing   bool
	hissings  []int
	fightings []int
	_point    geom.Point
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

func (cat *cat) NumHissings() int {
	return len(cat.hissings)
}

func (cat *cat) HissingsElementAt(idx int) int {
	return cat.hissings[idx]
}

func (cat *cat) NumFightings() int {
	return len(cat.fightings)
}

func (cat *cat) FightingsElementAt(idx int) int {
	return cat.fightings[idx]
}

func (cat *cat) Copy(dst Cat) {
	dst.setStatus(cat.status)
	dst.setHissing(cat.hissing)
	dst.copyHissingsFrom(cat.hissings)
	dst.copyFightingsFrom(cat.fightings)

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

func (cat *cat) copyHissingsFrom(src []int) {
	cat.hissings = cat.hissings[:len(src)]

	copy(cat.hissings, src)
}

func (cat *cat) copyFightingsFrom(src []int) {
	cat.fightings = cat.fightings[:len(src)]

	copy(cat.fightings, src)
}

func (cat *cat) clean() {
	cat.status = Calm
	cat.hissing = false
	cat.hissings = cat.hissings[:0]
	cat.fightings = cat.fightings[:0]
}

func newCat(point geom.Point, numCats int) *cat {
	cat := new(cat)

	cat.status = Calm
	cat.hissing = false

	cat.hissings = make([]int, 0, numCats)
	cat.fightings = make([]int, 0, numCats)

	cat._point = point

	return cat
}
