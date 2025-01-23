package viewModel

import (
	"context"
	"github.com/PavlushaSource/Radar/view"
	"github.com/hajimehoshi/ebiten/v2"
	"sync"
	"time"

	"github.com/PavlushaSource/Radar/view/utils"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/view/config"
)

type Producer interface {
	StartAppAction(context.Context)
}

type producer struct {
	next          chan struct{}
	radarSettings *config.RadarSettings
	appConfig     *config.ApplicationConfig
	app           *view.Application
	engine        *engine.Engine
}

func NewProducer(app *view.Application) Producer {
	return &producer{
		app:    app,
		engine: newEngine(app.RadarSettings, app.AppConfig, app.Borders),
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

func StartApp(app *view.Application) {
	app.InMainMenu = false

	ebiten.SetWindowTitle("Собака съела товар, теперь она наркоман")
	ebiten.SetWindowSize(app.AppConfig.WindowX, app.AppConfig.WindowY)

	ctx, cancel := context.WithCancel(context.Background())

	app.CancelFunc = cancel
	prod := NewProducer(app)
	prod.StartAppAction(ctx)
}
