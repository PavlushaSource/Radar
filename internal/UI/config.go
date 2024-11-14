package UI

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/internal/config"
)

func CreateConfigChoiceFunction(w1, w2 fyne.Window, configBack *config.BackendConfig, configUI *config.UIConfig, ctx context.Context) func() fyne.CanvasObject {
	return func() fyne.CanvasObject {

		catCount := widget.NewEntry()
		catCount.PlaceHolder = fmt.Sprintf("%d", configBack.CountCats)

		updateTime := widget.NewEntry()
		updateTime.PlaceHolder = "1.5"

		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Count cats", Widget: catCount, HintText: "enter a positive number up to 5 * 10^4"},
				{Text: "Update time", Widget: updateTime, HintText: "time in second"},
			},
			OnCancel: func() {

			},
			OnSubmit: func() {

			},
			SubmitText: "Run",
			CancelText: "Reset",
		}
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

		// TODO Delete this
		_ = runButton

		selectDistanceType := widget.NewSelect([]string{"Euclidean", "Manhattan", "Curvilinear"}, func(s string) {
			switch s {
			case "Euclidean":
				configBack.DistanceType = config.Euclidean
			case "Manhattan":
				configBack.DistanceType = config.Manhattan
			case "Curvilinear":
				configBack.DistanceType = config.Curvilinear
			}
		})
		selectDistanceType.PlaceHolder = "Distance type"

		return container.NewCenter(form)
	}
}
