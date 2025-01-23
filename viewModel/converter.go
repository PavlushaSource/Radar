package viewModel

import (
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view/api"
	"github.com/PavlushaSource/Radar/view/config"
)

func ConvertStateToVMCat(state *engine.State, scaleEngineCoord fyne.Size, paddingEngineCoord fyne.Position, catSize fyne.Size, catScale fyne.Size) []api.Cat {
	vmCats := make([]api.Cat, 0, state.NumCats())

	moveX := (catSize.Width * catScale.Width) / 2
	moveY := (catSize.Width * catScale.Height) / 2

	for i := 0; i < state.NumCats(); i++ {
		c := state.CatsElementAt(i)
		fmt.Println(c.X(), c.Y())
		x := float32(c.X())*scaleEngineCoord.Width + paddingEngineCoord.X - moveX
		y := float32(c.Y())*scaleEngineCoord.Height + paddingEngineCoord.Y - moveY
		vmCats = append(vmCats, api.Cat{X: x, Y: y, Color: ConvertStatusToColor(c)})
	}

	return vmCats
}

var choiceDistanceCalcType = map[config.DistanceType]geom.Distance{
	config.Euclidean:   geom.EuclideanDistance,
	config.Manhattan:   geom.ManhattanDistance,
	config.Curvilinear: geom.CurvilinearDistance,
}

func ConvertDistanceTypeToDistance(distanceType config.DistanceType) geom.Distance {
	return choiceDistanceCalcType[distanceType]
}

type GeomCreateFunction func(height float64, width float64, barriers []geom.Barrier, maxMoveDistance float64, distance geom.Distance, rndAsync rnd.RndAsync) geom.Geom

var choiceGeometryCalcType = map[config.GeometryType]GeomCreateFunction{
	config.Simple: geom.NewSimpleGeom,
	config.Vector: geom.NewVectorGeom,
}

func ConvertGeometryTypeToGeometry(geometryType config.GeometryType) GeomCreateFunction {
	return choiceGeometryCalcType[geometryType]
}
