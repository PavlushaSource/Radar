package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/PavlushaSource/Radar/view/api"
)

var ConvertStringToDistanceType = map[string]api.DistanceType{
	"Euclidean":   api.Euclidean,
	"Manhattan":   api.Manhattan,
	"Curvilinear": api.Curvilinear,
}

var ConvertDistanceTypeToString = map[api.DistanceType]string{
	api.Euclidean:   "Euclidean",
	api.Manhattan:   "Manhattan",
	api.Curvilinear: "Curvilinear",
}

var ConvertStringToGeometryType = map[string]api.GeometryType{
	"Simple": api.Simple,
	"Vector": api.Vector,
}

var ConvertGeometryTypeToString = map[api.GeometryType]string{
	api.Simple: "Simple",
	api.Vector: "Vector",
}

func ConvertVMCatToCanvasCat(source []api.Cat, catSize fyne.Size) []fyne.CanvasObject {
	canvasCatSlice := make([]fyne.CanvasObject, 0, len(source))

	// TODO parallel this. ATTENTION: len(source), cap(source)
	for _, s := range source {
		img := canvas.NewImageFromResource(getResourceCatSvg(s.Color))
		img.Move(fyne.Position{X: s.X, Y: s.Y})
		img.Resize(catSize)

		canvasCatSlice = append(canvasCatSlice, img)
	}

	return canvasCatSlice
}
