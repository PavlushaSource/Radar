package viewModel

import (
	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view/api"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
)

func convertEngineToViewStatus(status engine.Status) *utils.Status {
	var st utils.Status
	switch status {
	case engine.Hissing:
		st = utils.Hiss
	case engine.Fighting:
		st = utils.Fight
	case engine.Calm:
		st = utils.Run
	}

	return &st
}

func ConvertStateToViewDog(state *engine.State) []*api.Dog {
	dogs := make([]*api.Dog, 0, state.NumDogs())

	for i := 0; i < state.NumDogs(); i++ {
		c := state.DogsElementAt(i)
		dogs = append(dogs, &api.Dog{
			Status: convertEngineToViewStatus(c.Status()),
			X:      c.X(), Y: c.Y(),
			XNext: -1, YNext: -1,
			SpeedX: 0, SpeedY: 0,
		})
	}

	return dogs
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

func ConvertBorderViewToBarriers(borders []utils.Line) []geom.Barrier {
	barriers := make([]geom.Barrier, 0)

	for _, b := range borders {
		startPoint := geom.NewPoint(b.StartX, b.StartY)
		endPoint := geom.NewPoint(b.EndX, b.EndY)

		barriers = append(barriers, geom.NewBarrier(startPoint, endPoint))
	}

	return barriers
}
