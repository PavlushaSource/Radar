package view

import (
	"fmt"
	"strconv"
	"time"
)

type DistanceType uint8

const (
	Euclidean DistanceType = iota
	Manhattan
	Curvilinear
)

type GeometryType uint8

const (
	Simple GeometryType = iota
	Vector
)

type RadarSettings struct {
	DistanceType                  DistanceType
	UpdateTime                    time.Duration
	CountCats                     int
	FightingRadius, HissingRadius float64
	BufferSize                    int64
	GeometryType                  GeometryType
}

func NewRadarSettings() RadarSettings {
	return RadarSettings{
		DistanceType:   Euclidean,
		UpdateTime:     time.Millisecond * 2000,
		CountCats:      25,
		FightingRadius: 100,
		HissingRadius:  200,
		BufferSize:     16,
		GeometryType:   Simple,
	}
}

func (c *RadarSettings) SetCountCats(s string) error {
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

func (c *RadarSettings) SetUpdateTime(s string) error {
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

func (c *RadarSettings) SetFightingRadius(s string) error {
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid convert radius %s to float: %w", s, err)
	}
	if r <= 0 {
		return fmt.Errorf("radius must be greater than zero")
	}

	c.FightingRadius = r
	return nil
}

func (c *RadarSettings) SetHissingRadius(s string) error {
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid convert radius %s to float: %w", s, err)
	}
	if r <= 0 {
		return fmt.Errorf("radius must be greater than zero")
	}

	c.HissingRadius = r
	return nil
}
