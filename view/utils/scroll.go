package utils

import "fyne.io/fyne/v2"

const (
	scaleRatio = 1.1
)

func IncreaseScale(container *fyne.Container, width, height, scale float32) float32 {
	newScale := scale * scaleRatio

	updateContainerSize(container, width, height, newScale)

	return newScale
}

func DecreaseScale(container *fyne.Container, width, height, scale float32) float32 {
	newScale := scale / scaleRatio

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
