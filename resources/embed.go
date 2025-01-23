package resources

import (
	_ "embed"
)

var (
	//go:embed DogFighting.png
	DogFight_png []byte

	//go:embed DogHissing.png
	DogHiss_png []byte

	//go:embed DogRunning.png
	DogRun_png []byte

	//go:embed woodCut.png
	Border_png []byte
)
