package view

import (
	"context"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/PavlushaSource/Radar/view/UI"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/customTheme"
	"github.com/PavlushaSource/Radar/viewModel"
)

const (
	scaleRatio = 1.1
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

func (app *radarApplication) applyTheme() {
	app.app.Settings().SetTheme(customTheme.GetApplicationTheme(app.appConfig))
}

func (app *radarApplication) registerScaleKeyboard(catsContainer fyne.CanvasObject, layout *UI.CatsLayout) func(rune) {
	return func(r rune) {
		if r == '=' {
			layout.Scale *= scaleRatio
		} else if r == '-' {
			layout.Scale /= scaleRatio
		} else {
			return
		}
		nextSize := fyne.Size{Width: app.appConfig.WindowSize.Width * layout.Scale, Height: app.appConfig.WindowSize.Height * layout.Scale}

		catsContainer.Resize(nextSize)
	}
}

func (app *radarApplication) createRadarWindowContent(cats []fyne.CanvasObject) fyne.CanvasObject {
	toolbarCreate := UI.CreateToolbarFunction(app.homeAction, app.themeAction, app.fullscreenAction)

	layout := UI.CatsLayout{Scale: 1, PrevSize: app.AppConfig().WindowSize}
	catsContainer := container.New(&layout, cats...)
	background := UI.CreateCatsBoard(cats, &layout)
	app.RadarWindow().Canvas().SetOnTypedRune(app.registerScaleKeyboard(catsContainer, &layout))
	content := container.NewStack(background, catsContainer)

	return container.NewBorder(toolbarCreate(), nil, nil, nil, content)
}

func (app *radarApplication) createMainWindow(ctx context.Context) {
	loadWindow := UI.NewLoader("Please wait...", app.appConfig, app.settingsMenuWindow)
	toolbarCreateFunction := UI.CreateToolbarFunction(app.homeAction, app.themeAction, app.fullscreenAction)
	radarSettings := NewRadarSettings()

	onConfigChoice := func(chosenRadarSettings RadarSettings, appConfig config.ApplicationConfig) {
		loadWindow.Start()
		producer := viewModel.NewProducer(chosenRadarSettings, appConfig)
		loadWindow.Stop()

		// тут будет канал возвращаться и мы не рисуем, пока первые позицию не вернуться
		canvasCatsChannel := producer.StartAppAction(ctx)

		app.radarWindow.SetContent(app.createRadarWindowContent(<-canvasCatsChannel))

		app.hideWindow(app.settingsMenuWindow)
		app.showWindow(app.radarWindow)

		app.appConfig.InMainMenu = false
	}

	onConfigChoiceError := func(err error) {
		dialog.ShowError(err, app.settingsMenuWindow)
	}

	configChoice := UI.CreateSettingsChoiceFunction(radarSettings, app.appConfig, onConfigChoice, onConfigChoiceError)

	app.settingsMenuWindow.SetContent(
		container.NewBorder(toolbarCreateFunction(), UI.CreateBottom(), nil, nil, configChoice()),
	)
	app.settingsMenuWindow.Resize(app.appConfig.WindowSize)
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
