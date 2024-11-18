package utils

import "fyne.io/fyne/v2"

const (
	scaleRatio = 1.1
	maxScale   = 2
	minScale   = 0.01
)

func IncreaseScale(container *fyne.Container, width, height, scale float32) float32 {
	newScale := min(scale*scaleRatio, maxScale)

	updateContainerSize(container, width, height, newScale)

	return newScale
}

func DecreaseScale(container *fyne.Container, width, height, scale float32) float32 {
	newScale := max(scale/scaleRatio, minScale)

	updateContainerSize(container, width, height, newScale)

	return newScale
}

func updateContainerSize(container *fyne.Container, width, height, scale float32) {
	nextSize := fyne.Size{
		Width:  width * scale,
		Height: height * scale,
	}

	container.Resize(nextSize)
}
