package viewModel

import (
	"context"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/model/runner"
	"github.com/PavlushaSource/Radar/view"
	"github.com/PavlushaSource/Radar/view/config"
)

var choiceDistanceCalcType = map[view.DistanceType]geom.Distance{
	view.Euclidean:   geom.EuclideanDistance,
	view.Manhattan:   geom.ManhattanDistance,
	view.Curvilinear: geom.CurvilinearDistance,
}

type geomCreateFunction func(height float64, width float64, barriers []geom.Barrier, distance geom.Distance) geom.Geom

var choiceGeometryCalcType = map[view.GeometryType]geomCreateFunction{
	view.Simple: geom.NewSimpleGeom,
	view.Vector: geom.NewVectorGeom,
}

type Producer interface {
	StartAppAction(context.Context)
}

type producer struct {
	chosenRadarSettings view.RadarSettings
	appConfig           config.ApplicationConfig

	MVCatChannel chan Cat
	runner       runner.Runner
}

func NewProducer(chosenRadarSettings view.RadarSettings, appConfig config.ApplicationConfig) Producer {
	return &producer{
		chosenRadarSettings: chosenRadarSettings,
		appConfig:           appConfig,
		runner:              newEngineRunner(chosenRadarSettings, appConfig),
		MVCatChannel:        make(chan Cat, chosenRadarSettings.BufferSize),
	}
}

func newEngineRunner(chosenRadarSettings view.RadarSettings, appConfig config.ApplicationConfig) runner.Runner {
	geomCreateFunc := choiceGeometryCalcType[chosenRadarSettings.GeometryType]

	geomImpl := geomCreateFunc(
		float64(appConfig.WindowSize.Height),
		float64(appConfig.WindowSize.Width),
		make([]geom.Barrier, 0),
		choiceDistanceCalcType[chosenRadarSettings.DistanceType],
	)

	engineImpl := engine.NewEngine(
		chosenRadarSettings.FightingRadius,
		chosenRadarSettings.HissingRadius,
		int64(chosenRadarSettings.CountCats),
		geomImpl,
	)

	return runner.NewRunner(engineImpl, chosenRadarSettings.BufferSize)
}

func (p *producer) StartAppAction(ctx context.Context) {

}
