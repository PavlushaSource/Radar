package engine

import "github.com/PavlushaSource/Radar/model/geom"

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
	setHissing(hissing bool)
	copyHissingsFrom(src []int)
	copyFightingsFrom(src []int)
}

type State interface {
	Height() float64
	Width() float64
	RadiusHiss() float64
	RadiusFight() float64

	NumCats() int
	CatsElementAt(idx int) Cat

	Copy(dst State)

	setHeight(height float64)
	setWidth(width float64)
	setRadiusHiss(radiusHiss float64)
	setRadiusFight(radiusFight float64)

	copyCatsFrom(src []*cat)

	clean()
}

type Engine interface {
	Run() State
}
