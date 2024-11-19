package viewModel

import (
	"context"
	"image/color"
	"sync"
	"time"

	"fyne.io/fyne/v2/canvas"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view/api"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
)

type Producer interface {
	StartAppAction(context.Context) []*canvas.Circle
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

	// TODO: Think about passing max distance to geom
	geomImpl := geomCreateFunc(
		float64(appConfig.WindowSize.Height),
		float64(appConfig.WindowSize.Width),
		make([]geom.Barrier, 0),
		// math.Max(float64(appConfig.WindowSize.Height), float64(appConfig.WindowSize.Width))/10,
		150,
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
func (p *producer) StartAppAction(ctx context.Context) []*canvas.Circle {
	getCh, putCh := p.engine.Run(ctx)

	state := <-getCh

	// Initial cats positions
	VMCat := ConvertStateToVMCat(state, p.appConfig.ScaleEngineCoord, p.appConfig.PaddingEngineCoord)
	cats := ConvertVMCatToCanvasCat(VMCat, p.appConfig.CatSize)
	putCh <- state
	colorCats := make([]api.Color, len(cats))
	for i, c := range VMCat {
		colorCats[i] = c.Color
	}

	catsUpdater := func() {
		state := <-getCh
		VMCatNext := ConvertStateToVMCat(state, p.appConfig.ScaleEngineCoord, p.appConfig.PaddingEngineCoord)
		uiCats := ConvertVMCatToCanvasCat(VMCatNext, p.appConfig.CatSize)

		for i, c := range VMCatNext {
			VMCatNext[i].Color = c.Color
		}

		putCh <- state
		wg := sync.WaitGroup{}
		//fmt.Println("ScaleEngine Coord", p.appConfig.ScaleEngineCoord)
		for i, c := range uiCats {
			wg.Add(2)
			go func() {
				defer wg.Done()
				//fmt.Printf("Next postition %d cat = %f / %f\n", i, c.Position().X, c.Position().Y)
				utils.AnimateCat(
					cats[i].Position(),
					c.Position(),
					cats[i],
					int32(p.chosenRadarSettings.UpdateTime.Milliseconds()/2),
				)
			}()
			go func() {
				defer wg.Done()
				//fmt.Println(utils.ColorToRGBA(api.Red), utils.ColorToRGBA(api.Blue))
				canvas.NewColorRGBAAnimation(ColorToRGBA(colorCats[i]), ColorToRGBA(VMCatNext[i].Color), p.chosenRadarSettings.UpdateTime/2, func(c color.Color) {
					cats[i].FillColor = c
					canvas.Refresh(cats[i])
				}).Start()
				colorCats[i] = VMCatNext[i].Color
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
