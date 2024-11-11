package main

import "image/color"

type Color uint8

const (
	SkipColor Color = iota
	Red
	Blue
	Green
)

func ConvertColor(c Color) color.Color {
	switch c {
	case Red:
		return color.RGBA{R: 255}
	case Blue:
		return color.RGBA{B: 255}
	case Green:
		return color.RGBA{G: 255}

	default:
		panic("unhandled default case")
	}
}
