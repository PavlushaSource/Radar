package viewModel

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/view/UI"
	"github.com/PavlushaSource/Radar/view/utils"
	"sync"
	"time"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view/api"
	"github.com/PavlushaSource/Radar/view/config"
)

type Producer interface {
	StartAppAction(context.Context) []*UI.DogUI
}

const maxRadiusMove = 500

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
		maxRadiusMove,
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
}

// StartAppAction TODO: in viewModel we must work with view Api, not directly with CanvasObjects
func (p *producer) StartAppAction(ctx context.Context) []*UI.DogUI {
	getCh, putCh := p.engine.Run(ctx)

	state := <-getCh

	// Initial dogs positions
	VMDog := p.ConvertStateToVMDog(state)
	dogs := p.ConvertVMDogToUI(VMDog)
	putCh <- state
	prevTime := time.Now()
	catsUpdater := func() {
		state := <-getCh
		VMCatNext := p.ConvertStateToVMDog(state)
		UIDogs := p.ConvertVMDogToUI(VMCatNext)
		putCh <- state
		wg := sync.WaitGroup{}
		fmt.Println("СОБАКИ НАЧАЛО АНИМАЦИИ ", time.Since(prevTime))
		prevTime = time.Now()
		for i, d := range UIDogs {
			wg.Add(1)
			go func() {
				defer wg.Done()
				utils.AnimateDog(
					dogs[i].Position(),
					d.Position(),
					dogs[i],
					int32(p.chosenRadarSettings.UpdateTime.Milliseconds()),
				)
				dogs[i].SetImage(d.Resource)
			}()
		}
		wg.Wait()
		fmt.Println("ВСЕ ГОРУТИНЫ АНИМАЦИИ Отработали")
	}

	ticker := time.Tick(p.chosenRadarSettings.UpdateTime)

	withTicker(ctx, ticker, catsUpdater)

	return dogs
}

func (p *producer) ConvertStateToVMDog(state *engine.State) []api.Dog {

	vmDogs := make([]api.Dog, 0, state.NumCats())

	for i := 0; i < state.NumCats(); i++ {
		c := state.CatsElementAt(i)
		x := float32(c.X())*p.appConfig.Scale + p.appConfig.PaddingEnginePos.X*p.appConfig.Scale
		y := float32(c.Y())*p.appConfig.Scale + p.appConfig.PaddingEnginePos.Y*p.appConfig.Scale
		vmDogs = append(vmDogs, api.Dog{X: x, Y: y, Color: ConvertStatusToColor(c)})
	}

	return vmDogs

}

func (p *producer) ConvertVMDogToUI(source []api.Dog) []*UI.DogUI {
	DogUISlice := make([]*UI.DogUI, 0)

	// TODO parallel this. ATTENTION: len(source), cap(source)
	for _, s := range source {
		dog := UI.NewDogUI(UI.GetResourceDogSvg(s.Color))
		dog.Resize(fyne.NewSize(p.appConfig.CatSize.Width*p.appConfig.Scale, p.appConfig.CatSize.Height*p.appConfig.Scale))
		dog.Move(fyne.Position{X: s.X, Y: s.Y})
		dog.MoveCenterPosition()

		DogUISlice = append(DogUISlice, dog)
	}
	return DogUISlice

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
