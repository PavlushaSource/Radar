package api

type DistanceType uint8

const (
	Euclidean DistanceType = iota
	Manhattan
	Curvilinear
)
