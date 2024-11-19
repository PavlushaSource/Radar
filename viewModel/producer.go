package viewModel

import (
	"context"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view/api"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
)

type Producer interface {
	StartAppAction(context.Context) []fyne.CanvasObject
}

type producer struct {
	chosenRadarSettings api.RadarSettings
	appConfig           *config.ApplicationConfig
	engine              *engine.Engine
}

func NewProducer(chosenRadarSettings api.RadarSettings, appConfig *config.ApplicationConfig) Producer {
	return &producer{
		chosenRadarSettings: chosenRadarSettings,
		appConfig:           appConfig,
		engine:              newEngine(chosenRadarSettings, *appConfig),
	}
}

func newEngine(chosenRadarSettings api.RadarSettings, appConfig config.ApplicationConfig) *engine.Engine {
	geomCreateFunc := ConvertGeometryTypeToGeometry(chosenRadarSettings.GeometryType)

	rndAsync := rnd.NewRndCore()

	geomImpl := geomCreateFunc(
		float64(appConfig.WindowSize.Height),
		float64(appConfig.WindowSize.Width),
		make([]geom.Barrier, 0),
		ConvertDistanceTypeToDistance(chosenRadarSettings.DistanceType),
		rndAsync,
	)

	return engine.NewEngine(
		chosenRadarSettings.FightingRadius,
		chosenRadarSettings.HissingRadius,
		chosenRadarSettings.CountCats,
		geomImpl,
		rndAsync,
		chosenRadarSettings.BufferSize,
	)

	//return runner.NewRunner(engineImpl, chosenRadarSettings.BufferSize)
}

// StartAppAction TODO: in viewModel we must work with view Api, not directly with CanvasObjects
func (p *producer) StartAppAction(ctx context.Context) []fyne.CanvasObject {
	getCh, putCh := p.engine.Run(ctx)

	state := <-getCh

	// Initial cats positions
	cats := utils.ConvertVMCatToCanvasCat(ConvertStateToVMCat(state, p.appConfig.ScaleEngineCoord, p.appConfig.PaddingEngineCoord), p.appConfig.CatSize)
	putCh <- state

	catsUpdater := func() {
		state := <-getCh
		uiCats := utils.ConvertVMCatToCanvasCat(ConvertStateToVMCat(state, p.appConfig.ScaleEngineCoord, p.appConfig.PaddingEngineCoord), p.appConfig.CatSize)
		putCh <- state
		wg := sync.WaitGroup{}
		//fmt.Println("ScaleEngine Coord", p.appConfig.ScaleEngineCoord)
		for i, c := range uiCats {
			wg.Add(1)
			go func() {
				defer wg.Done()
				//fmt.Printf("Next postition %d cat = %f / %f\n", i, c.Position().X, c.Position().Y)
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
