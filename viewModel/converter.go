package viewModel

import (
	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view/api"
)

func ConvertStateToVMCat(state *engine.State, scaleEngineCoord fyne.Size, paddingEngineCoord fyne.Position) []api.Cat {
	vmCats := make([]api.Cat, 0, state.NumCats())

	for i := 0; i < state.NumCats(); i++ {
		c := state.CatsElementAt(i)
		x := float32(c.X())*scaleEngineCoord.Width + paddingEngineCoord.X
		y := float32(c.Y())*scaleEngineCoord.Height + paddingEngineCoord.Y
		vmCats = append(vmCats, api.Cat{X: x, Y: y, Color: ConvertStatusToColor(c)})
	}

	return vmCats
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
