package main

import (
	"context"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/PavlushaSource/Radar/custom_theme"
	"github.com/PavlushaSource/Radar/internal/UI"
	"github.com/PavlushaSource/Radar/internal/config"
)

func main() {
	ctx := context.Background()

	App := app.New()
	App.Settings().SetTheme(custom_theme.NewLightTheme())

	mainWindow := App.NewWindow("Main menu")
	runWindow := App.NewWindow("Run")

	UIConfig := config.NewUIConfig()
	BackendConfig := config.NewBackendConfig()

	toolbarCreate := UI.CreateToolbarFunction(App, mainWindow, runWindow, &UIConfig, ctx)
	configChoice := UI.CreateConfigChoiceFunction(mainWindow, runWindow, &BackendConfig, &UIConfig, ctx)

	mainWindow.SetContent(container.NewBorder(toolbarCreate(), UI.CreateBottom(), nil, nil, configChoice()))

	RadarContainer := UI.CreateContentRunWindow(mainWindow, runWindow, &BackendConfig, &UIConfig, ctx)

	//go func() {
	//
	//	for {
	//		select {
	//		case <-time.After(time.Second):
	//			fmt.Println("Scale now")
	//			fmt.Println(RadarContainer.Size())
	//			scale += 0.15
	//			currSize := RadarContainer.Size()
	//			nextSize := fyne.Size{Width: currSize.Width * scale, Height: currSize.Height * scale}
	//			RadarContainer.Resize(nextSize)
	//			runWindow.Canvas().Refresh(RadarContainer)
	//		}
	//	}
	//}()

	runWindow.SetContent(container.NewBorder(toolbarCreate(), nil, nil, nil, RadarContainer))

	mainWindow.Resize(UIConfig.WindowSize)
	mainWindow.CenterOnScreen()
	runWindow.Resize(UIConfig.WindowSize)

	mainWindow.Show()
	App.Run()
}

//
//func main() {
//	a := app.New()
//	w := a.NewWindow("Diagonal")
//
//	text1 := widget.NewLabel("topleft")
//	text2 := widget.NewLabel("Middle Label")
//	text3 := widget.NewLabel("bottomright")
//
//	w.SetContent(container.New(&diagonal{}, text1, text2, text3))
//	w.ShowAndRun()
//}
