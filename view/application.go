package view

import (
	"context"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/PavlushaSource/Radar/view/UI"
	"github.com/PavlushaSource/Radar/view/api"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/customTheme"
	"github.com/PavlushaSource/Radar/view/utils"
	"github.com/PavlushaSource/Radar/viewModel"
)

type RadarApplication interface {
	Run(ctx context.Context)

	App() fyne.App
	AppConfig() config.ApplicationConfig
	SettingsMenuWindow() fyne.Window
	RadarWindow() fyne.Window
}

type radarApplication struct {
	app                  fyne.App
	appConfig            config.ApplicationConfig
	settingsMenuWindow   fyne.Window
	radarWindow          fyne.Window
	radarEngineCtxCancel context.CancelFunc
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

	app.createMainWindowContent(ctx)
	app.showWindow(app.settingsMenuWindow)

	app.app.Run()
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
		app.radarEngineCtxCancel()
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

func (app *radarApplication) applyTheme() {
	app.app.Settings().SetTheme(customTheme.GetApplicationTheme(app.appConfig))
}

func (app *radarApplication) registerScaleKeyboard(catsContainer *fyne.Container, layout *UI.CatsLayout) func(rune) {
	width := app.appConfig.WindowSize.Width
	height := app.appConfig.WindowSize.Height

	return func(r rune) {
		if r == '=' {
			layout.Scale = utils.IncreaseScale(catsContainer, width, height, layout.Scale)
		} else if r == '-' {
			layout.Scale = utils.DecreaseScale(catsContainer, width, height, layout.Scale)
		} else {
			return
		}
	}
}

func (app *radarApplication) createRadarWindowContent(cats []*canvas.Circle) fyne.CanvasObject {
	toolbarCreate := UI.CreateToolbarFunction(app.homeAction, app.themeAction, app.fullscreenAction)

	// TODO add CatsLayout constructor
	// create border rectangle and move this
	rect := canvas.NewRectangle(color.Transparent)
	rect.Resize(fyne.NewSize(app.appConfig.WindowSize.Width, app.appConfig.WindowSize.Height))
	rect.StrokeColor = color.Black
	rect.StrokeWidth = 3

	fmt.Println("SIZE RECT", rect.Size())
	layout := UI.CatsLayout{Scale: 1, PrevSize: app.appConfig.WindowSize, AppConfig: &app.appConfig, Border: rect}

	catsUI := make([]fyne.CanvasObject, 0)
	for i := range cats {
		catsUI = append(catsUI, cats[i])
	}
	objects := make([]fyne.CanvasObject, 0, len(cats)+1)
	objects = append(objects, rect)
	objects = append(objects, catsUI...)

	catsContainer := container.New(&layout, objects...)
	//TODO: change windowSize -> container size
	catsScrollableContainer := UI.CreateCatsScrollableContainer(catsUI, catsContainer, &layout, app.appConfig.WindowSize, &app.appConfig)
	content := container.NewStack(catsScrollableContainer, catsContainer)

	app.RadarWindow().Canvas().SetOnTypedRune(app.registerScaleKeyboard(catsContainer, &layout))

	return container.NewBorder(toolbarCreate(), nil, nil, nil, content)
}

func (app *radarApplication) createMainWindowContent(ctx context.Context) {
	loadWindowSize := fyne.NewSize(app.appConfig.WindowSize.Width/4, app.appConfig.WindowSize.Height/4)
	loadWindow := UI.NewLoader("Please wait...", loadWindowSize, app.settingsMenuWindow)

	toolbarCreateFunction := UI.CreateToolbarFunction(app.homeAction, app.themeAction, app.fullscreenAction)
	radarSettings := api.NewRadarSettings()

	onConfigChoice := func(chosenRadarSettings api.RadarSettings) {
		loadWindow.Start()

		producer := viewModel.NewProducer(chosenRadarSettings, &app.appConfig)
		// TODO: Don't store enfine ctx cancel func in radarApplication (check homeAction)
		ctx, cancel := context.WithCancel(ctx)
		app.radarEngineCtxCancel = cancel
		canvasCats := producer.StartAppAction(ctx)

		loadWindow.Stop()

		app.radarWindow.SetContent(app.createRadarWindowContent(canvasCats))

		app.hideWindow(app.settingsMenuWindow)
		app.showWindow(app.radarWindow)

		app.appConfig.InMainMenu = false
	}

	onConfigChoiceError := func(err error) {
		dialog.ShowError(err, app.settingsMenuWindow)
	}

	configChoice := UI.CreateSettingsChoiceFunction(radarSettings, onConfigChoice, onConfigChoiceError)

	app.settingsMenuWindow.SetContent(
		container.NewBorder(toolbarCreateFunction(), UI.CreateBottom(), nil, nil, configChoice()),
	)
}

func NewApplication() RadarApplication {
	app := new(radarApplication)

	app.app = fyneApp.New()
	app.appConfig = config.NewApplicationConfig()
	app.settingsMenuWindow = app.app.NewWindow("Main menu")
	app.radarWindow = app.app.NewWindow("Radar")

	return app
}
