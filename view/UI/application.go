package UI

import (
	"context"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/PavlushaSource/Radar/model/runner"
	"github.com/PavlushaSource/Radar/view"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/customTheme"
	"github.com/PavlushaSource/Radar/view/utils"
	"github.com/PavlushaSource/Radar/viewModel"
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
	app.applyTheme()

	app.createMainWindow(ctx)
	app.settingsMenuWindow.ShowAndRun()
}

func (app *radarApplication) hideWindow(window fyne.Window) {
	window.SetFullScreen(false)
	app.appConfig.FullScreenMode = false
	window.Hide()
}

func (app *radarApplication) showWindow(window fyne.Window) {
	window.CenterOnScreen()
	window.Resize(app.appConfig.WindowSize)
	window.SetMaster()
	window.Show()
}

func (app *radarApplication) homeAction() {
	if !app.appConfig.InMainMenu {
		app.hideWindow(app.radarWindow)
		app.showWindow(app.settingsMenuWindow)

		app.appConfig.InMainMenu = true
	}
}

func (app *radarApplication) themeAction() {
	app.appConfig.LightThemeFlag = !app.appConfig.LightThemeFlag
	app.app.Settings().SetTheme(customTheme.GetApplicationTheme(app.appConfig))
}

func (app *radarApplication) fullscreenAction() {
	app.appConfig.FullScreenMode = !app.appConfig.FullScreenMode
	if app.appConfig.InMainMenu {
		app.settingsMenuWindow.SetFullScreen(app.appConfig.FullScreenMode)
	} else {
		app.radarWindow.SetFullScreen(app.appConfig.FullScreenMode)
	}
}

func (app *radarApplication) RunRadarWindow(
	ctx context.Context,
	engineRunner runner.Runner,
	UpdateTime time.Duration,
	toolbarCreate func() fyne.CanvasObject,
) {

	app.showWindow(app.radarWindow)
	app.appConfig.InMainMenu = false

	ch := engineRunner.Run(ctx)

	VMCats := viewModel.ConvertStateToVMCat(<-ch)
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

	//RadarContainer := CreateContentRunWindow(app, UICats)

	//app.radarWindow.SetContent(container.NewBorder(toolbarCreate(), nil, nil, nil, RadarContainer))
}

func (app *radarApplication) applyTheme() {
	app.app.Settings().SetTheme(customTheme.GetApplicationTheme(app.appConfig))
}

func (app *radarApplication) createRadarWindowContent(cats []fyne.CanvasObject) fyne.CanvasObject {
	toolbarCreate := createToolbarFunction(app.homeAction, app.themeAction, app.fullscreenAction)

	layout := CatsLayout{Scale: 1, prevSize: app.AppConfig().WindowSize}
	catsContainer := container.New(&layout, cats...)
	background := CreateCatsBoard(cats, &layout)
	app.RadarWindow().Canvas().SetOnTypedRune(RegisterScaleRune(app.RadarWindow(), catsContainer, &layout, app.AppConfig()))
	content := container.NewStack(background, catsContainer)

	//RadarContainer := CreateContentRunWindow(app, cats)

	return container.NewBorder(toolbarCreate(), nil, nil, nil, content)
}

func (app *radarApplication) createMainWindow(ctx context.Context) {
	loadWindow := NewLoader("Please wait...", app.appConfig, app.settingsMenuWindow)

	toolbarCreate := createToolbarFunction(app.homeAction, app.themeAction, app.fullscreenAction)

	radarSettings := view.NewRadarSettings()

	onConfigChoice := func(chosenRadarSettings view.RadarSettings, appConfig config.ApplicationConfig) {
		loadWindow.Start()
		producer := viewModel.NewProducer(chosenRadarSettings, appConfig)
		loadWindow.Stop()

		// тут будет канал возвращаться и мы не рисуем, пока первые позицию не вернуться
		producer.StartAppAction()

		app.radarWindow.SetContent(app.createRadarWindowContent())

		app.hideWindow(app.settingsMenuWindow)
		app.showWindow(app.radarWindow)

		app.appConfig.InMainMenu = false

		app.RunRadarWindow(ctx, engineRunner, chosenRadarSettings.UpdateTime, toolbarCreate)
	}

	onConfigChoiceError := func(err error) {
		dialog.ShowError(err, app.settingsMenuWindow)
	}

	configChoice := CreateSettingsChoiceFunction(radarSettings, onConfigChoice, onConfigChoiceError)

	app.settingsMenuWindow.SetContent(
		container.NewBorder(toolbarCreate(), createBottom(), nil, nil, configChoice()),
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
