package config

import (
	"fmt"
	"strconv"
	"time"
)

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
		UpdateTime:   time.Millisecond * 500,
		CountCats:    25,
		AngryRadius:  100,
	}
}

func (c *BackendConfig) SetCountCats(s string) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("invalid convert count %s to integer: %w", s, err)
	}

	if n <= 0 || n > 5*1e5 {
		return fmt.Errorf("boundaries cat must be 1 <= count <= 5 * 10^5, but got %d", n)
	}

	c.CountCats = n
	return nil
}

func (c *BackendConfig) SetUpdateTime(s string) error {
	updateTime, err := time.ParseDuration(s + "s")
	if err != nil {
		return fmt.Errorf("invalid update time %s: %w", s, err)
	}

	if updateTime.Milliseconds() < (250 * time.Millisecond).Milliseconds() {
		return fmt.Errorf("time must be greater than or eqaul to 250ms")
	}

	c.UpdateTime = updateTime
	return nil
}

func (c *BackendConfig) SetAngryRadius(s string) error {
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid convert radius %s to float: %w", s, err)
	}
	if r <= 0 {
		return fmt.Errorf("radius must be greater than zero")
	}

	c.AngryRadius = r
	return nil
}
