package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func CreateCats(source []producer.Cat, catSize fyne.Size) []fyne.CanvasObject {
	res := make([]fyne.CanvasObject, 0)

	for _, s := range source {
		img := canvas.NewImageFromResource(resourceCat(s.Color))
		img.Move(fyne.Position{X: s.X, Y: s.Y})
		img.Resize(catSize)

		res = append(res, img)
	}

	return res
}
