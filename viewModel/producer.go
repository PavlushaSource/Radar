package viewModel

import (
	"context"
	"github.com/PavlushaSource/Radar/view"
	"sync"
	"time"

	"github.com/PavlushaSource/Radar/view/utils"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view/config"
)

type Producer interface {
	StartAppAction(context.Context)
}

type producer struct {
	app    *view.Application
	engine *engine.Engine
}

func NewProducer(app *view.Application) Producer {
	return &producer{
		app:    app,
		engine: newEngine(app.RadarSettings, app.AppConfig),
	}
}

func newEngine(chosenRadarSettings *config.RadarSettings, appConfig *config.ApplicationConfig) *engine.Engine {
	geomCreateFunc := ConvertGeometryTypeToGeometry(chosenRadarSettings.GeometryType)

	rndAsync := rnd.NewRndCore()

	geomImpl := geomCreateFunc(
		float64(appConfig.WindowY),
		float64(appConfig.WindowX),
		make([]geom.Barrier, 0),
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

func (p *producer) StartAppAction(ctx context.Context) {
	getCh, putCh := p.engine.Run(ctx)

	state := <-getCh

	// Initial dogs positions
	dogs := ConvertStateToViewDog(state)
	p.app.Dogs = dogs
	putCh <- state

	dogsUpdater := func() {
		<-p.app.NeedNext

		state = <-getCh
		nextDogs := ConvertStateToViewDog(state)
		putCh <- state

		wg := sync.WaitGroup{}
		for i, d := range nextDogs {
			wg.Add(1)
			go func() {
				defer wg.Done()
				dogs[i].UpdateDogMove(d, p.app.RadarSettings.UpdateTime)
				dogs[i].Status = d.Status
			}()
		}
		wg.Wait()
	}

	ticker := time.Tick(p.app.RadarSettings.UpdateTime)

	utils.WithTicker(ctx, ticker, dogsUpdater)

	return
}
