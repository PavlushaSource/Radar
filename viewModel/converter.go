package viewModel

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view/api"
)

func ConvertStateToVMCat(state *engine.State, scaleEngineCoord fyne.Size, paddingEngineCoord fyne.Position, catSize fyne.Size, catScale fyne.Size) []api.Cat {
	vmCats := make([]api.Cat, 0, state.NumCats())

	moveX := (catSize.Width * catScale.Width) / 2
	moveY := (catSize.Width * catScale.Height) / 2

	for i := 0; i < state.NumCats(); i++ {
		c := state.CatsElementAt(i)
		x := float32(c.X())*scaleEngineCoord.Width + paddingEngineCoord.X - moveX
		y := float32(c.Y())*scaleEngineCoord.Height + paddingEngineCoord.Y - moveY
		vmCats = append(vmCats, api.Cat{X: x, Y: y, Color: ConvertStatusToColor(c)})
	}

	return vmCats
}

func ColorToRGBA(c api.Color) color.NRGBA {
	switch c {
	case api.Red:
		return color.NRGBA{R: 0xff, A: 0xff}
	case api.Blue:
		return color.NRGBA{B: 0xff, A: 0xff}
	case api.Purple:
		return color.NRGBA{G: 0xff, A: 0xff}
	}
	fmt.Println("HIHIHIHI")
	return color.NRGBA{}
}

func ConvertStatusToColor(cat *engine.Cat) api.Color {
	switch cat.Status() {
	case engine.Calm:
		return api.Blue
	case engine.Hissing:
		return api.Purple
	case engine.Fighting:
		return api.Red
	default:
		panic("Undefined Color")
	}
}

var choiceDistanceCalcType = map[api.DistanceType]geom.Distance{
	api.Euclidean:   geom.EuclideanDistance,
	api.Manhattan:   geom.ManhattanDistance,
	api.Curvilinear: geom.CurvilinearDistance,
}

func ConvertDistanceTypeToDistance(distanceType api.DistanceType) geom.Distance {
	return choiceDistanceCalcType[distanceType]
}

type GeomCreateFunction func(height float64, width float64, barriers []geom.Barrier, maxMoveDistance float64, distance geom.Distance, rndAsync rnd.RndAsync) geom.Geom

var choiceGeometryCalcType = map[api.GeometryType]GeomCreateFunction{
	api.Simple: geom.NewSimpleGeom,
	api.Vector: geom.NewVectorGeom,
}

func ConvertGeometryTypeToGeometry(geometryType api.GeometryType) GeomCreateFunction {
	return choiceGeometryCalcType[geometryType]
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
