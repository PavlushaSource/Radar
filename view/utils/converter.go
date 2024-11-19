package utils

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/PavlushaSource/Radar/view/api"
	"image/color"
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

func ConvertVMCatToCanvasCat(source []api.Cat, catSize fyne.Size) []*canvas.Circle {
	canvasCatSlice := make([]*canvas.Circle, 0)

	// TODO parallel this. ATTENTION: len(source), cap(source)
	for _, s := range source {
		circle := canvas.NewCircle(ColorToRGBA(s.Color))
		circle.Move(fyne.Position{X: s.X, Y: s.Y})
		circle.Resize(catSize)
		circle.FillColor.RGBA()
		//img := canvas.NewImageFromResource(getResourceCatSvg(s.Color))
		//img.Move(fyne.Position{X: s.X, Y: s.Y})
		//img.Resize(catSize)

		//canvasCatSlice = append(canvasCatSlice, img)
		canvasCatSlice = append(canvasCatSlice, circle)
	}

	return canvasCatSlice
}

func ColorToRGBA(c api.Color) color.NRGBA {
	switch c {
	case api.Red:
		return color.NRGBA{R: 0xff, A: 0xff}
	case api.Blue:
		return color.NRGBA{B: 0xff, A: 0xff}
	case api.Purple:
		return color.NRGBA{B: 128, A: 0xff, R: 128}
	}
	fmt.Println("HIHIHIHI")
	return color.NRGBA{}
}
