package utils

import (
	"fmt"

	"fyne.io/fyne/v2"
)

const (
	ScaleRatio = 1.1
	MaxScale   = 3
	MinScale   = 0.01
)

func IncreaseScale(container *fyne.Container, width, height, scale float32) float32 {
	newScale := min(scale*ScaleRatio, MaxScale)

	fmt.Println("RESIZE CONTAINER", width, height)
	updateContainerSize(container, width, height, newScale)

	return newScale
}

func DecreaseScale(container *fyne.Container, width, height, scale float32) float32 {
	newScale := max(scale/ScaleRatio, MinScale)

	fmt.Println("RESIZE CONTAINER", width, height)
	updateContainerSize(container, width, height, newScale)

	return newScale
}

func updateContainerSize(container *fyne.Container, width, height, scale float32) {
	nextSize := fyne.Size{
		Width:  width * scale,
		Height: height * scale,
	}
	fmt.Println("NEW SCALE", scale, "PREV CONTAINER SIZE", container.Size(), "NEXT CONTAINER RESIZE", width*scale, height*scale)
	container.Resize(nextSize)
}
