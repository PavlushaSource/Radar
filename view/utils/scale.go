package utils

import (
	"fmt"

	"fyne.io/fyne/v2"
)

const (
	scaleRatio = 1.1
	maxScale   = 3
	minScale   = 0.01
)

func IncreaseScale(container *fyne.Container, width, height, scale float32) float32 {
	newScale := min(scale*scaleRatio, maxScale)

	fmt.Println("RESIZE CONTAINER", width, height)
	updateContainerSize(container, width, height, newScale)

	return newScale
}

func DecreaseScale(container *fyne.Container, width, height, scale float32) float32 {
	newScale := max(scale/scaleRatio, minScale)

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
