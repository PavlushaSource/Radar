package UI

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/customTheme"
	"github.com/PavlushaSource/Radar/view/domain"
	"github.com/PavlushaSource/Radar/view/utils"
	"sync"
	"time"
)

type RadarApplication interface {
	Run(ctx context.Context)

	App() fyne.App
	AppConfig() config.ApplicationConfig
	SettingsMenuWindow() fyne.Window
	RadarWindow() fyne.Window
}

type radarApplication struct {
	app                fyne.App
	appConfig          config.ApplicationConfig
	settingsMenuWindow fyne.Window
	radarWindow        fyne.Window
}

func (app *radarApplication) App() fyne.App {
	return app.app
}

func (app *radarApplication) AppConfig() config.ApplicationConfig {
	return app.appConfig
}

func (app *radarApplication) SettingsMenuWindow() fyne.Window {
	return app.settingsMenuWindow
}

func (app *radarApplication) RadarWindow() fyne.Window {
	return app.radarWindow
}

func (app *radarApplication) Run(ctx context.Context) {
	app.setupApplication()

	app.setupSettingsMenuWindow(ctx)
	app.settingsMenuWindow.ShowAndRun()
}

func (app *radarApplication) RunRadarWindow(
	ctx context.Context,
	engineRunner engine.Runner,
	UpdateTime time.Duration,
	toolbarCreate func() fyne.CanvasObject,
) {
	app.radarWindow.SetMaster()
	app.radarWindow.CenterOnScreen()
	app.radarWindow.Resize(app.appConfig.WindowSize)

	app.radarWindow.Show()

	app.appConfig.InMainMenu = false

	ch := engineRunner.Run(ctx)

	VMCats := ConvertStateToVMCat(<-ch)
	UICats := utils.CreateCats(VMCats, app.appConfig.CatSize)

	//TODO: move to ViewModel
	ticker := time.Tick(UpdateTime)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				vmCats := ConvertStateToVMCat(<-ch)
				<-ticker
				uiCats := utils.CreateCats(vmCats, app.appConfig.CatSize)
				//time.Sleep(6 * time.Second)
				wg := sync.WaitGroup{}
				for i, c := range uiCats {
					wg.Add(1)
					go func() {
						defer wg.Done()
						//fmt.Println("prevCatPosition", UICats[i].Position())
						AnimateCat(UICats[i].Position(), c.Position(), UICats[i], 100)
						UICats[i] = c
						//fmt.Println("afterCatPosition", UICats[i].Position())
					}()
				}
				wg.Wait()

			}
		}
	}()

	RadarContainer := CreateContentRunWindow(app, UICats)

	app.radarWindow.SetContent(container.NewBorder(toolbarCreate(), nil, nil, nil, RadarContainer))
}

// TODO: move to ViewModel
func ConvertStateToVMCat(state engine.State) []domain.Cat {
	vmCats := make([]domain.Cat, 0, state.NumCats())

	for _, c := range state.Cats() {
		fmt.Println(c.Status())
		fmt.Println(ConvertStatusToColor(c))
		vmCats = append(vmCats, domain.Cat{X: float32(c.X()), Y: float32(c.Y()), Color: ConvertStatusToColor(c)})
	}

	return vmCats
}

// TODO: move to ViewModel
func ConvertStatusToColor(cat engine.Cat) domain.Color {
	switch cat.Status() {
	case engine.Calm:
		return domain.Blue
	case engine.Hissing:
		return domain.Purple
	case engine.Fighting:
		return domain.Red
	default:
		panic("Undefined Color")
	}
}

func (app *radarApplication) setupApplication() {
	app.app.Settings().SetTheme(customTheme.GetApplicationTheme(app.appConfig))
}

func (app *radarApplication) setupSettingsMenuWindow(ctx context.Context) {
	onHomeAction := func() {
		if !app.appConfig.InMainMenu {
			app.radarWindow.Hide()

			app.settingsMenuWindow.CenterOnScreen()
			app.settingsMenuWindow.Resize(app.appConfig.WindowSize)
			app.settingsMenuWindow.Show()
			app.settingsMenuWindow.SetMaster()

			app.appConfig.InMainMenu = true
		}
	}

	onThemeAction := func() {
		app.appConfig.LightThemeFlag = !app.appConfig.LightThemeFlag
		app.app.Settings().SetTheme(customTheme.GetApplicationTheme(app.appConfig))
	}

	onFullScreenAction := func() {
		app.appConfig.FullScreenMode = !app.appConfig.FullScreenMode
		if app.appConfig.InMainMenu {
			app.settingsMenuWindow.SetFullScreen(app.appConfig.FullScreenMode)
		} else {
			app.radarWindow.SetFullScreen(app.appConfig.FullScreenMode)
		}
	}

	toolbarCreate := CreateToolbarFunction(onHomeAction, onThemeAction, onFullScreenAction)

	radarSettings := view.NewRadarSettings()

	onConfigChoice := func(chosenRadarSettings view.RadarSettings) {
		app.settingsMenuWindow.Hide()

		//TODO: move to ViewModel
		geomImpl := geom.NewSimpleGeom(
			float64(app.appConfig.WindowSize.Height),
			float64(app.appConfig.WindowSize.Width),
			make([]geom.Barrier, 0),
			geom.EuclideanDistance,
		)

		engineImpl := engine.NewEngine(
			chosenRadarSettings.FightingRadius,
			chosenRadarSettings.HissingRadius,
			int64(chosenRadarSettings.CountCats),
			geomImpl,
		)

		engineRunner := engine.NewRunner(engineImpl, 30)

		app.RunRadarWindow(ctx, engineRunner, chosenRadarSettings.UpdateTime, toolbarCreate)
	}

	onConfigChoiceError := func(err error) {
		dialog.ShowError(err, app.settingsMenuWindow)
	}

	configChoice := CreateConfigChoiceFunction(radarSettings, onConfigChoice, onConfigChoiceError)

	app.settingsMenuWindow.SetContent(
		container.NewBorder(toolbarCreate(), CreateBottom(), nil, nil, configChoice()),
	)
	app.settingsMenuWindow.Resize(app.appConfig.WindowSize)
	app.settingsMenuWindow.CenterOnScreen()
}

func (app *radarApplication) setupRadarWindow(cats []fyne.CanvasObject, toolbarCreate func() fyne.CanvasObject) {
	radarContainer := CreateContentRunWindow(app, cats)

	app.radarWindow.SetContent(container.NewBorder(toolbarCreate(), nil, nil, nil, radarContainer))
	app.radarWindow.Resize(app.appConfig.WindowSize)
	app.settingsMenuWindow.CenterOnScreen()
}

func NewApplication() RadarApplication {
	app := new(radarApplication)

	app.app = fyneApp.New()
	app.appConfig = config.NewApplicationConfig()
	app.settingsMenuWindow = app.app.NewWindow("Main menu")
	app.radarWindow = app.app.NewWindow("Radar")

	return app
}
