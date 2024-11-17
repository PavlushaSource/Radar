package main

import (
	"context"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/PavlushaSource/Radar/view/UI"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/customTheme"
)

func main() {
	ctx := context.Background()

	App := app.New()
	App.Settings().SetTheme(customTheme.NewLightTheme())

	mainWindow := App.NewWindow("Main menu")
	runWindow := App.NewWindow("Run")

	UIConfig := config.NewUIConfig()
	BackendConfig := config.NewBackendConfig()

	toolbarCreate := UI.CreateToolbarFunction(App, mainWindow, runWindow, &UIConfig, ctx)
	configChoice := UI.CreateConfigChoiceFunction(mainWindow, runWindow, &BackendConfig, &UIConfig, ctx)

	mainWindow.SetContent(container.NewBorder(toolbarCreate(), UI.CreateBottom(), nil, nil, configChoice()))

	RadarContainer := UI.CreateContentRunWindow(mainWindow, runWindow, &BackendConfig, &UIConfig, ctx)

	runWindow.SetContent(container.NewBorder(toolbarCreate(), nil, nil, nil, RadarContainer))

	mainWindow.Resize(UIConfig.WindowSize)
	mainWindow.CenterOnScreen()
	runWindow.Resize(UIConfig.WindowSize)

	mainWindow.Show()
	App.Run()
}
