package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/tools/playground"
)

type CatBackend struct {
	X     float32
	Y     float32
	color Color
}

func CaptureCat(cat CatBackend) *canvas.Image {
	c := playground.NewSoftwareCanvas()
	c.SetPadded(false)

	obj := canvas.NewCircle(ConvertColor(cat.color))
	obj.Move(fyne.NewPos(cat.X, cat.Y))
	c.SetContent(obj)

	img := c.Capture()
	return canvas.NewImageFromImage(img)
}
