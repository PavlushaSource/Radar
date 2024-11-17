package domain

type Cat struct {
	X     float32
	Y     float32
	Color Color
}

type Color uint8

const (
	Red Color = iota
	Blue
	Purple
)
