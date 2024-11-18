package UI

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/view/api"
	"github.com/PavlushaSource/Radar/view/utils"
)

const (
	runText   = "Run"
	resetText = "Reset"
)

func CreateSettingsChoiceFunction(
	radarSettings api.RadarSettings,
	onConfigChoice func(radarSettings api.RadarSettings),
	onConfigChoiceError func(err error),
) func() fyne.CanvasObject {
	return func() fyne.CanvasObject {

		catCount := widget.NewEntry()
		catDefaultValue := fmt.Sprintf("%d", radarSettings.CountCats)
		catCount.SetText(catDefaultValue)

		updateTime := widget.NewEntry()
		updateTimeDefaultValue := radarSettings.UpdateTime.String()
		updateTimeDefaultValue = updateTimeDefaultValue[:len(updateTimeDefaultValue)-1]
		updateTime.SetText(updateTimeDefaultValue)

		fightingRadius := widget.NewEntry()
		fightingRadiusDefaultValue := fmt.Sprintf("%f", radarSettings.FightingRadius)
		fightingRadius.SetText(fightingRadiusDefaultValue)

		hissingRadius := widget.NewEntry()
		hissingRadiusDefaultValue := fmt.Sprintf("%f", radarSettings.HissingRadius)
		hissingRadius.SetText(hissingRadiusDefaultValue)

		selectDistanceType := widget.NewSelect([]string{"Euclidean", "Manhattan", "Curvilinear"}, func(s string) {
			radarSettings.DistanceType = utils.ConvertStringToDistanceType[s]
		})
		selectDistanceType.PlaceHolder = ""
		selectDistanceTypeDefaultValue := utils.ConvertDistanceTypeToString[radarSettings.DistanceType]
		//selectDistanceTypeDefaultValue := ""
		selectDistanceType.SetSelected(selectDistanceTypeDefaultValue)

		selectGeometryType := widget.NewSelect([]string{"Simple", "Vector"}, func(s string) {
			radarSettings.GeometryType = utils.ConvertStringToGeometryType[s]
		})
		selectGeometryType.PlaceHolder = ""
		selectGeometryTypeDefaultValue := utils.ConvertGeometryTypeToString[radarSettings.GeometryType]
		//selectGeometryTypeDefaultValue := ""
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

				resErr = utils.SaveError(resErr, radarSettings.SetCountCats(catCount.Text))
				resErr = utils.SaveError(resErr, radarSettings.SetUpdateTime(updateTime.Text))
				resErr = utils.SaveError(resErr, radarSettings.SetFightingRadius(fightingRadius.Text))
				resErr = utils.SaveError(resErr, radarSettings.SetHissingRadius(hissingRadius.Text))

				if resErr != nil {
					onConfigChoiceError(resErr)
				} else {
					onConfigChoice(radarSettings)
				}
			},
			SubmitText: runText,
			CancelText: resetText,
		}

		return container.NewCenter(form)
	}
}
