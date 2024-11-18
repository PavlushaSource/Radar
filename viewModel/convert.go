package viewModel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/view/utils"
)

func ConvertStateToVMCat(state engine.State) []Cat {
	vmCats := make([]Cat, 0, state.NumCats())

	for _, c := range state.Cats() {
		vmCats = append(vmCats, Cat{X: float32(c.X()), Y: float32(c.Y()), Color: ConvertStatusToColor(c)})
	}

	return vmCats
}

func ConvertStatusToColor(cat engine.Cat) Color {
	switch cat.Status() {
	case engine.Calm:
		return Blue
	case engine.Hissing:
		return Purple
	case engine.Fighting:
		return Red
	default:
		panic("Undefined Color")
	}
}

func ConvertMVCatToCanvasCat(source []Cat, catSize fyne.Size) []fyne.CanvasObject {
	canvasCatSlice := make([]fyne.CanvasObject, 0, len(source))

	// TODO parallel this. ATTENTION: len(source), cap(source)
	for _, s := range source {
		img := canvas.NewImageFromResource(utils.ResourceCat(s.Color))
		img.Move(fyne.Position{X: s.X, Y: s.Y})
		img.Resize(catSize)

		canvasCatSlice = append(canvasCatSlice, img)
	}

	return canvasCatSlice
}
