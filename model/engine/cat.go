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
	Point() geom.Point
	Hissings() []int64
	Fightings() []int64
	Copy() Cat

	setStatus(status Status)

	isHissing() bool
	setHissing(hissing bool)

	setHissings(hissings []int64)
	setFightings(fightings []int64)

	clean()
}

type cat struct {
	status    Status
	hissing   bool
	hissings  []int64
	fightings []int64
	point     geom.Point
}

func (cat *cat) Status() Status {
	return cat.status
}

func (cat *cat) X() float64 {
	return cat.point.X()
}

func (cat *cat) Y() float64 {
	return cat.point.Y()
}

func (cat *cat) Point() geom.Point {
	return cat.point
}

func (cat *cat) Hissings() []int64 {
	return cat.hissings
}

func (cat *cat) Fightings() []int64 {
	return cat.fightings
}

func (original *cat) Copy() Cat {
	copyCat := new(cat)

	copyCat.status = original.status
	copyCat.hissing = original.hissing
	copyCat.point = original.point.Copy()

	copyCat.hissings = make([]int64, 0, len(original.hissings))
	copy(copyCat.hissings, original.hissings)

	copyCat.fightings = make([]int64, 0, len(original.fightings))
	copy(copyCat.fightings, original.fightings)

	return copyCat
}

func (cat *cat) setStatus(status Status) {
	cat.status = status
}

func (cat *cat) isHissing() bool {
	return cat.hissing
}

func (cat *cat) setHissing(hissing bool) {
	cat.hissing = hissing
}

func (cat *cat) setHissings(hissings []int64) {
	cat.hissings = hissings
}

func (cat *cat) setFightings(fightings []int64) {
	cat.fightings = fightings
}

func (cat *cat) clean() {
	cat.status = Calm
	cat.hissing = false

	capacity := cap(cat.hissings)
	cat.hissings = make([]int64, 0, capacity)
	capacity = cap(cat.fightings)
	cat.fightings = make([]int64, 0, capacity)
}

func newCat(point geom.Point, numCats int64) Cat {
	cat := new(cat)

	cat.status = Calm
	cat.hissing = false

	cat.hissings = make([]int64, 0, numCats)
	cat.fightings = make([]int64, 0, numCats)

	cat.point = point

	return cat
}
