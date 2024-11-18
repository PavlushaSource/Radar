package viewModel

import (
	"context"
	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/model/runner"
	"github.com/PavlushaSource/Radar/view"
	"github.com/PavlushaSource/Radar/view/config"
	"sync"
	"time"
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
	StartAppAction(context.Context) chan []fyne.CanvasObject
}

type producer struct {
	chosenRadarSettings view.RadarSettings
	appConfig           config.ApplicationConfig
	runner              runner.Runner
}

func NewProducer(chosenRadarSettings view.RadarSettings, appConfig config.ApplicationConfig) Producer {
	return &producer{
		chosenRadarSettings: chosenRadarSettings,
		appConfig:           appConfig,
		runner:              newEngineRunner(chosenRadarSettings, appConfig),
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

func (p *producer) StartAppAction(ctx context.Context) chan []fyne.CanvasObject {
	UICatChannel := make(chan []fyne.CanvasObject, p.chosenRadarSettings.BufferSize)

	engineStateCh := p.runner.Run(ctx)

	// TODO нужно одна init состояние для UICats, потом можно просто этот массив мувать и не переприсваивать
	// Первое чтение из канала как раз вычитает Init
	ticker := time.Tick(p.chosenRadarSettings.UpdateTime)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			// TODO put in ticker one value for init without timeout
			case <-ticker:
				uiCats := ConvertMVCatToCanvasCat(ConvertStateToVMCat(<-engineStateCh), p.appConfig.CatSize)
				wg := sync.WaitGroup{}
				for i, c := range uiCats {
					wg.Add(1)
					go func() {
						defer wg.Done()
						AnimateCat(UICats[i].Position(), c.Position(), UICats[i], 100)

						// TODO это не нужно делать, только поменять svg в случае смены цвета
						UICats[i] = c
					}()
				}
				wg.Wait()

			}
		}
	}()

	return UICatChannel
}
