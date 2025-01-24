package viewModel

import (
	"context"
	"github.com/PavlushaSource/Radar/view/api"
	"sync"
	"time"

	"github.com/PavlushaSource/Radar/view/utils"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/view/config"
)

type Producer interface {
	StartAppAction(context.Context) []*api.Dog
}

type producer struct {
	next          chan struct{}
	radarSettings *config.RadarSettings
	appConfig     *config.ApplicationConfig
	engine        *engine.Engine
}

func NewProducer(settings *config.RadarSettings, appConfig *config.ApplicationConfig, next chan struct{}, borders []utils.Line) Producer {
	return &producer{
		next:          next,
		radarSettings: settings,
		appConfig:     appConfig,
		engine:        newEngine(settings, appConfig, borders),
	}
}

func newEngine(chosenRadarSettings *config.RadarSettings, appConfig *config.ApplicationConfig, borders []utils.Line) *engine.Engine {
	geomCreateFunc := ConvertGeometryTypeToGeometry(chosenRadarSettings.GeometryType)

	rndAsync := rnd.NewRndCore()

	geomImpl := geomCreateFunc(
		float64(appConfig.WindowY),
		float64(appConfig.WindowX),
		ConvertBorderViewToBarriers(borders),
		chosenRadarSettings.MaxRadiusMove,
		ConvertDistanceTypeToDistance(chosenRadarSettings.DistanceType),
		rndAsync,
	)

	return engine.NewEngine(
		chosenRadarSettings.FightingRadius,
		chosenRadarSettings.HissingRadius,
		chosenRadarSettings.CountDogs,
		geomImpl,
		rndAsync,
		chosenRadarSettings.BufferSize,
	)
}

func (p *producer) StartAppAction(ctx context.Context) []*api.Dog {
	getCh, putCh := p.engine.Run(ctx)

	state := <-getCh

	// Initial dogs positions
	dogs := ConvertStateToViewDog(state)
	putCh <- state
	go func() {
		p.next <- struct{}{}
	}()

	dogsUpdater := func() {
		state = <-getCh
		nextDogs := ConvertStateToViewDog(state)
		putCh <- state

		wg := sync.WaitGroup{}
		for i, d := range nextDogs {
			wg.Add(1)
			go func() {
				defer wg.Done()
				dogs[i].UpdateDogMove(d, p.radarSettings.UpdateTime)
				dogs[i].Status = d.Status
			}()
		}
		wg.Wait()
	}

	ticker := time.Tick(p.radarSettings.UpdateTime)

	utils.WithTicker(ctx, ticker, dogsUpdater)
	return dogs
}
