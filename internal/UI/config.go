package UI

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/internal/config"
	"github.com/PavlushaSource/Radar/internal/utils"
)

func CreateConfigChoiceFunction(w1, w2 fyne.Window, configBack *config.BackendConfig, configUI *config.UIConfig, ctx context.Context) func() fyne.CanvasObject {
	return func() fyne.CanvasObject {

		catCount := widget.NewEntry()
		catDefaultValue := fmt.Sprintf("%d", configBack.CountCats)
		catCount.SetText(catDefaultValue)

		updateTime := widget.NewEntry()
		updateTimeDefaultValue := "0.5"
		updateTime.SetText(updateTimeDefaultValue)

		angryRadius := widget.NewEntry()
		angryRadiusDefaultValue := "100"
		angryRadius.SetText(angryRadiusDefaultValue)

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
		selectDistanceType.PlaceHolder = ""
		selectDistanceTypeDefaultValue := "Euclidean"
		selectDistanceType.SetSelected(selectDistanceTypeDefaultValue)

		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Count cats", Widget: catCount, HintText: "enter a positive number up to 5 * 10^4"},
				{Text: "Update time", Widget: updateTime, HintText: "time in second"},
				{Text: "Angry radius", Widget: angryRadius, HintText: "radius in pixels"},
				{Text: "Distance type", Widget: selectDistanceType},
			},
			OnCancel: func() {
				catCount.SetText(catDefaultValue)
				updateTime.SetText(updateTimeDefaultValue)
				angryRadius.SetText(angryRadiusDefaultValue)
				selectDistanceType.SetSelected(selectDistanceTypeDefaultValue)

			},
			OnSubmit: func() {
				var resErr error

				resErr = utils.SaveError(resErr, configBack.SetCountCats(catCount.Text))
				resErr = utils.SaveError(resErr, configBack.SetUpdateTime(updateTime.Text))
				resErr = utils.SaveError(resErr, configBack.SetAngryRadius(angryRadius.Text))

				if resErr != nil {
					dialog.ShowError(resErr, w1)
					fmt.Println(resErr)
				} else {
					fmt.Println("No errors in config")
					//w1.SetFullScreen(false)
					//configUI.FullScreenMode = false
					//w1.Hide()
					//
					//w2.SetMaster()
					//w2.CenterOnScreen()
					//w2.Resize(configUI.WindowSize)
					//w2.Show()
					//
					//configUI.InMainMenu = false
				}
			},
			SubmitText: "Run",
			CancelText: "Reset",
		}

		//form.Append("Distance type", selectDistanceType)

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

		return container.NewCenter(form)
	}
}
