package geom

import (
	"math"
	"math/rand"
	"time"
)

const MaxPointMoveDegree = 359

type vectorGeom struct {
	baseGeom
}

func (geom *vectorGeom) MovePoint(point Point) {
	distance := geom.rng.Float64() * math.Max(geom.width, geom.width)
	degree := geom.rng.Float64() * MaxPointMoveDegree

	radians := degree * math.Pi / 180.0

	x := math.Max(math.Min(point.X()+distance*math.Cos(radians), geom.width), 0)
	y := math.Max(math.Min(point.Y()+distance*math.Sin(radians), geom.height), 0)

	point.set(x, y)
}

func NewVectorGeom(height float64, width float64, barriers []Barrier, distance Distance) Geom {
	geom := new(vectorGeom)
	geom.height = height
	geom.width = width
	geom.barriers = barriers
	geom.distance = distance
	geom.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	return geom
}
