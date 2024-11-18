package viewModel

import (
	"context"
	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/model/runner"
	"github.com/PavlushaSource/Radar/view/api"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
	"sync"
	"time"
)

type Producer interface {
	StartAppAction(context.Context) []fyne.CanvasObject
}

type producer struct {
	chosenRadarSettings api.RadarSettings
	appConfig           config.ApplicationConfig
	runner              runner.Runner
}

func NewProducer(chosenRadarSettings api.RadarSettings, appConfig config.ApplicationConfig) Producer {
	return &producer{
		chosenRadarSettings: chosenRadarSettings,
		appConfig:           appConfig,
		runner:              newEngineRunner(chosenRadarSettings, appConfig),
	}
}

func newEngineRunner(chosenRadarSettings api.RadarSettings, appConfig config.ApplicationConfig) runner.Runner {
	geomCreateFunc := ConvertGeometryTypeToGeometry(chosenRadarSettings.GeometryType)

	geomImpl := geomCreateFunc(
		float64(appConfig.WindowSize.Height),
		float64(appConfig.WindowSize.Width),
		make([]geom.Barrier, 0),
		ConvertDistanceTypeToDistance(chosenRadarSettings.DistanceType),
	)

	engineImpl := engine.NewEngine(
		chosenRadarSettings.FightingRadius,
		chosenRadarSettings.HissingRadius,
		int64(chosenRadarSettings.CountCats),
		geomImpl,
	)

	return runner.NewRunner(engineImpl, chosenRadarSettings.BufferSize)
}

// StartAppAction TODO: in viewModel we must work with view Api, not directly with CanvasObjects
func (p *producer) StartAppAction(ctx context.Context) []fyne.CanvasObject {
	engineStateCh := p.runner.Run(ctx)

	// Initial cats positions
	cats := utils.ConvertVMCatToCanvasCat(ConvertStateToVMCat(<-engineStateCh), p.appConfig.CatSize)

	catsUpdater := func() {
		uiCats := utils.ConvertVMCatToCanvasCat(ConvertStateToVMCat(<-engineStateCh), p.appConfig.CatSize)
		wg := sync.WaitGroup{}
		for i, c := range uiCats {
			wg.Add(1)
			go func() {
				defer wg.Done()

				utils.AnimateCat(
					cats[i].Position(),
					c.Position(),
					cats[i],
					int32(p.chosenRadarSettings.UpdateTime.Milliseconds()),
				)

			}()
		}
		wg.Wait()
	}

	ticker := time.Tick(p.chosenRadarSettings.UpdateTime)

	withTicker(ctx, ticker, catsUpdater)

	return cats
}

// TODO: move to external core package
func withTicker(ctx context.Context, ticker <-chan time.Time, action func()) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker:
				action()
			}
		}
	}()
}
