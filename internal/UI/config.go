package UI

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/internal/config"
)

func CreateConfigChoiceFunction(w1, w2 fyne.Window, configBack *config.BackendConfig, configUI *config.UIConfig, ctx context.Context) func() fyne.CanvasObject {
	return func() fyne.CanvasObject {
		runButton := widget.NewButton("Run", func() {
			w1.SetFullScreen(false)
			configUI.FullScreenMode = false
			w1.Hide()

			w2.SetMaster()
			w2.CenterOnScreen()
			w2.Resize(configUI.WindowSize)
			w2.Show()

			configUI.InMainMenu = false
		})

		return container.NewCenter(runButton)
	}
}
