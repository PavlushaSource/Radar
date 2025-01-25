package config

import (
	"fmt"
	"strconv"
	"time"
)

type RadarSettings struct {
	DistanceType                  DistanceType
	UpdateTime                    time.Duration
	CountDogs                     int
	FightingRadius, HissingRadius float64
	BufferSize                    int
	GeometryType                  GeometryType
	MaxRadiusMove                 float64
}

func (settings *RadarSettings) SetCountDogs(s string) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("invalid convert count %s to integer: %w", s, err)
	}

	if n <= 0 || n > 5*1e5 {
		return fmt.Errorf("boundaries dogmust be 1 <= count <= 5 * 10^5, but got %d", n)
	}

	settings.CountDogs = n
	return nil
}

func (settings *RadarSettings) SetUpdateTime(s string) error {
	updateTime, err := time.ParseDuration(s + "s")
	if err != nil {
		return fmt.Errorf("invalid update time %s: %w", s, err)
	}

	if updateTime.Milliseconds() < (250 * time.Millisecond).Milliseconds() {
		return fmt.Errorf("time must be greater than or eqaul to 250ms")
	}

	settings.UpdateTime = updateTime
	return nil
}

func (settings *RadarSettings) SetFightingRadius(s string) error {
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid convert radius %s to float: %w", s, err)
	}
	if r <= 0 {
		return fmt.Errorf("radius must be greater than zero")
	}

	settings.FightingRadius = r
	return nil
}

func (settings *RadarSettings) SetHissingRadius(s string) error {
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid convert radius %s to float: %w", s, err)
	}
	if r <= 0 {
		return fmt.Errorf("radius must be greater than zero")
	}

	settings.HissingRadius = r
	return nil
}

func NewRadarSettings() *RadarSettings {
	return &RadarSettings{
		DistanceType:   Euclidean,
		UpdateTime:     time.Millisecond * 2000,
		CountDogs:      25,
		FightingRadius: 100,
		HissingRadius:  200,
		BufferSize:     16,
		GeometryType:   Simple,
		MaxRadiusMove:  5,
	}
}
