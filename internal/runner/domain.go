package runner

type CatBackend struct {
	X     float32
	Y     float32
	Color Color
}

type Color uint8

const (
	SkipColor Color = iota
	Red
	Blue
	Purple
)
