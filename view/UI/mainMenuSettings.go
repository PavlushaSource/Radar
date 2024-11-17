package UI

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
)

var convertStringToDistanceType = map[string]config.DistanceType{
	"Euclidean":   config.Euclidean,
	"Manhattan":   config.Manhattan,
	"Curvilinear": config.Curvilinear,
}

var convertDistanceTypeToString = map[config.DistanceType]string{
	config.Euclidean:   "Euclidean",
	config.Manhattan:   "Manhattan",
	config.Curvilinear: "Curvilinear",
}

var convertStringToGeometryType = map[string]config.GeometryType{
	"Simple": config.Simple,
	"Vector": config.Vector,
}

var convertGeometryTypeToString = map[config.GeometryType]string{
	config.Simple: "Simple",
	config.Vector: "Vector",
}

const (
	runText   = "Run"
	resetText = "Reset"
)

func CreateConfigChoiceFunction(w1, w2 fyne.Window, backendSettings *config.BackendSettings, configUI *config.UIConfig, ctx context.Context) func() fyne.CanvasObject {
	return func() fyne.CanvasObject {

		catCount := widget.NewEntry()
		catDefaultValue := fmt.Sprintf("%d", backendSettings.CountCats)
		catCount.SetText(catDefaultValue)

		updateTime := widget.NewEntry()
		updateTimeDefaultValue := backendSettings.UpdateTime.String()
		updateTimeDefaultValue = updateTimeDefaultValue[:len(updateTimeDefaultValue)-1]
		updateTime.SetText(updateTimeDefaultValue)

		fightingRadius := widget.NewEntry()
		fightingRadiusDefaultValue := fmt.Sprintf("%f", backendSettings.FightingRadius)
		fightingRadius.SetText(fightingRadiusDefaultValue)

		hissingRadius := widget.NewEntry()
		hissingRadiusDefaultValue := fmt.Sprintf("%f", backendSettings.HissingRadius)
		hissingRadius.SetText(hissingRadiusDefaultValue)

		selectDistanceType := widget.NewSelect([]string{"Euclidean", "Manhattan", "Curvilinear"}, func(s string) {
			backendSettings.DistanceType = convertStringToDistanceType[s]
		})
		selectDistanceType.PlaceHolder = ""
		selectDistanceTypeDefaultValue := convertDistanceTypeToString[backendSettings.DistanceType]
		selectDistanceType.SetSelected(selectDistanceTypeDefaultValue)

		selectGeometryType := widget.NewSelect([]string{"Simple", "Vector"}, func(s string) {
			backendSettings.GeometryType = convertStringToGeometryType[s]
		})
		selectGeometryType.PlaceHolder = ""
		selectGeometryTypeDefaultValue := convertGeometryTypeToString[backendSettings.GeometryType]
		selectGeometryType.SetSelected(selectGeometryTypeDefaultValue)

		// TODO: mb move every variant to own constant entity
		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Count cats", Widget: catCount, HintText: "enter a positive number up to 5 * 10^4"},
				{Text: "Update time", Widget: updateTime, HintText: "time in second"},
				{Text: "Angry radius", Widget: fightingRadius, HintText: "size in pixels"},
				{Text: "Hissing radius", Widget: hissingRadius, HintText: "size in pixels"},
				{Text: "Distance type", Widget: selectDistanceType},
				{Text: "Geometry type", Widget: selectGeometryType},
			},
			OnCancel: func() {
				catCount.SetText(catDefaultValue)
				updateTime.SetText(updateTimeDefaultValue)
				fightingRadius.SetText(fightingRadiusDefaultValue)
				hissingRadius.SetText(hissingRadiusDefaultValue)
				selectDistanceType.SetSelected(selectDistanceTypeDefaultValue)
				selectGeometryType.SetSelected(selectGeometryTypeDefaultValue)

			},
			OnSubmit: func() {
				var resErr error

				resErr = utils.SaveError(resErr, backendSettings.SetCountCats(catCount.Text))
				resErr = utils.SaveError(resErr, backendSettings.SetUpdateTime(updateTime.Text))
				resErr = utils.SaveError(resErr, backendSettings.SetFightingRadius(fightingRadius.Text))
				resErr = utils.SaveError(resErr, backendSettings.SetHissingRadius(hissingRadius.Text))

				if resErr != nil {
					dialog.ShowError(resErr, w1)
				} else {
					w1.SetFullScreen(false)
					configUI.FullScreenMode = false
					w1.Hide()

					w2.SetMaster()
					w2.CenterOnScreen()
					w2.Resize(configUI.WindowSize)

					w2.Show()

					configUI.InMainMenu = false
				}
			},
			SubmitText: runText,
			CancelText: resetText,
		}

		return container.NewCenter(form)
	}
}
