package viewModel

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2/canvas"
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
	StartAppAction(context.Context) []*canvas.Circle
}

type producer struct {
	chosenRadarSettings api.RadarSettings
	appConfig           *config.ApplicationConfig
	runner              runner.Runner
}

func NewProducer(chosenRadarSettings api.RadarSettings, appConfig *config.ApplicationConfig) Producer {
	return &producer{
		chosenRadarSettings: chosenRadarSettings,
		appConfig:           appConfig,
		runner:              newEngineRunner(chosenRadarSettings, *appConfig),
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
func (p *producer) StartAppAction(ctx context.Context) []*canvas.Circle {
	engineStateCh := p.runner.Run(ctx)

	// Initial cats positions
	VMCat := ConvertStateToVMCat(<-engineStateCh, p.appConfig.ScaleEngineCoord, p.appConfig.PaddingEngineCoord)
	cats := utils.ConvertVMCatToCanvasCat(VMCat, p.appConfig.CatSize)
	colorCats := make([]api.Color, len(cats))
	//fmt.Println(len(VMCat))
	for i, c := range VMCat {
		colorCats[i] = c.Color
	}

	catsUpdater := func(clr api.Color) {
		st := time.Now()
		VMCatNext := ConvertStateToVMCat(<-engineStateCh, p.appConfig.ScaleEngineCoord, p.appConfig.PaddingEngineCoord)

		for i, _ := range VMCatNext {
			VMCatNext[i].Color = clr
		}
		uiCats := utils.ConvertVMCatToCanvasCat(VMCatNext, p.appConfig.CatSize)
		wg := sync.WaitGroup{}
		//fmt.Println("ScaleEngine Coord", p.appConfig.ScaleEngineCoord)
		for i, c := range uiCats {

			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				cats[i].Move(c.Position())
				cats[i].Refresh()
				//fmt.Printf("Next postition %d cat = %f / %f\n", i, c.Position().X, c.Position().Y)
				//utils.AnimateCat(
				//	cats[i].Position(),
				//	c.Position(),
				//	cats[i],
				//	int32(p.chosenRadarSettings.UpdateTime.Milliseconds()/2),
				//)

				//cats[i] = c
				//canvas.Refresh(cats[i])
				//cats[i].Refresh()

			}(i)
			//go func(i int) {
			//	defer wg.Done()
			//	//fmt.Println(utils.ColorToRGBA(api.Red), utils.ColorToRGBA(api.Blue))
			//	canvas.NewColorRGBAAnimation(utils.ColorToRGBA(colorCats[i]), utils.ColorToRGBA(VMCatNext[i].Color), p.chosenRadarSettings.UpdateTime/2, func(c color.Color) {
			//		cats[i].FillColor = c
			//		canvas.Refresh(cats[i])
			//	}).Start()
			//	colorCats[i] = VMCatNext[i].Color
			//}(i)
		}
		wg.Wait()
		end := time.Now()
		fmt.Println(end.Sub(st).Milliseconds())
		//canvas.Refresh()
	}

	ticker := time.Tick(p.chosenRadarSettings.UpdateTime)

	withTicker(ctx, ticker, catsUpdater)

	return cats
}

// TODO: move to external core package
func withTicker(ctx context.Context, ticker <-chan time.Time, action func(color api.Color)) {
	go func() {
		i := 0
		for {
			i += 1
			select {
			case <-ctx.Done():
				return

			case <-ticker:
				if i%3 == 0 {
					action(api.Red)
				} else if i%3 == 1 {
					action(api.Purple)
				} else {
					action(api.Blue)
				}

			}
		}
	}()
}
