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

type Cat interface {
	Status() Status
	X() float64
	Y() float64

	NumHissings() int
	HissingsElementAt(idx int) int
	NumFightings() int
	FightingsElementAt(idx int) int

	Copy(dst Cat)

	point() geom.Point

	setStatus(status Status)

	hissing() bool
	setHissing(hissing bool)

	addHissing(hissingCat int)
	addFighting(fightingCat int)

	copyHissingsFrom(src []int)
	copyFightingsFrom(src []int)

	clean()
}

type cat struct {
	_status    Status
	_hissing   bool
	_hissings  []int
	_fightings []int
	_point     geom.Point
}

func (cat *cat) Status() Status {
	return cat._status
}

func (cat *cat) X() float64 {
	return cat._point.X()
}

func (cat *cat) Y() float64 {
	return cat._point.Y()
}

func (cat *cat) NumHissings() int {
	return len(cat._hissings)
}

func (cat *cat) HissingsElementAt(idx int) int {
	return cat._hissings[idx]
}

func (cat *cat) NumFightings() int {
	return len(cat._fightings)
}

func (cat *cat) FightingsElementAt(idx int) int {
	return cat._fightings[idx]
}

func (cat *cat) Copy(dst Cat) {
	dst.setStatus(cat._status)
	dst.setHissing(cat._hissing)
	dst.copyHissingsFrom(cat._hissings)
	dst.copyFightingsFrom(cat._fightings)

	cat._point.Copy(dst.point())
}

func (cat *cat) point() geom.Point {
	return cat._point
}

func (cat *cat) setStatus(status Status) {
	cat._status = status
}

func (cat *cat) hissing() bool {
	return cat._hissing
}

func (cat *cat) setHissing(hissing bool) {
	cat._hissing = hissing
}

func (cat *cat) addHissing(hissingCat int) {
	cat._hissings = append(cat._hissings, hissingCat)
}

func (cat *cat) addFighting(fightingCat int) {
	cat._fightings = append(cat._fightings, fightingCat)
}

func (cat *cat) copyHissingsFrom(src []int) {
	cat._hissings = cat._hissings[:len(src)]

	copy(cat._hissings, src)
}

func (cat *cat) copyFightingsFrom(src []int) {
	cat._fightings = cat._fightings[:len(src)]

	copy(cat._fightings, src)
}

func (cat *cat) clean() {
	cat._status = Calm
	cat._hissing = false
	cat._hissings = cat._hissings[:0]
	cat._fightings = cat._fightings[:0]
}

func newCat(point geom.Point, numCats int64) Cat {
	cat := new(cat)

	cat._status = Calm
	cat._hissing = false

	cat._hissings = make([]int, 0, numCats)
	cat._fightings = make([]int, 0, numCats)

	cat._point = point

	return cat
}
