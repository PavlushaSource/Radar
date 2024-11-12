package config

import "time"

type DistanceType uint8

const (
	SkipZero DistanceType = iota
	Euclidean
	Manhattan
	Curvilinear
)

type BackendConfig struct {
	DistanceType DistanceType
	UpdateTime   time.Duration
	CountCats    int
	AngryRadius  float64
}

func NewBackendConfig() BackendConfig {
	return BackendConfig{
		DistanceType: Euclidean,
		UpdateTime:   time.Second * 1,
		CountCats:    25,
		AngryRadius:  100,
	}
}
