package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/PavlushaSource/Radar/view/domain"
	"math/rand/v2"
)

func CreateCats(source []domain.CatBackend, catSize fyne.Size) []fyne.CanvasObject {
	res := make([]fyne.CanvasObject, 0)

	for _, s := range source {
		img := canvas.NewImageFromResource(resourceCat(s.Color))
		img.Move(fyne.Position{X: s.X, Y: s.Y})
		img.Resize(catSize)

		res = append(res, img)
	}

	return res
}

func GenerateBackendCats(count int) []domain.CatBackend {
	res := make([]domain.CatBackend, count)

	minX, maxX := 0, 1920
	minY, maxY := 0, 1080

	for i := 0; i < count; i++ {
		currX := rand.IntN(maxX-minX) + minX
		currY := rand.IntN(maxY-minY) + minY

		res[i] = domain.CatBackend{X: float32(currX), Y: float32(currY)}
		res = append(res, domain.CatBackend{X: float32(currX), Y: float32(currY)})
	}
	return res
}
