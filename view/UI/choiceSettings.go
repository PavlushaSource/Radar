package UI

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/view"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
)

var convertStringToDistanceType = map[string]view.DistanceType{
	"Euclidean":   view.Euclidean,
	"Manhattan":   view.Manhattan,
	"Curvilinear": view.Curvilinear,
}

var convertDistanceTypeToString = map[view.DistanceType]string{
	view.Euclidean:   "Euclidean",
	view.Manhattan:   "Manhattan",
	view.Curvilinear: "Curvilinear",
}

var convertStringToGeometryType = map[string]view.GeometryType{
	"Simple": view.Simple,
	"Vector": view.Vector,
}

var convertGeometryTypeToString = map[view.GeometryType]string{
	view.Simple: "Simple",
	view.Vector: "Vector",
}

const (
	runText   = "Run"
	resetText = "Reset"
)

func CreateSettingsChoiceFunction(
	radarSettings view.RadarSettings,
	appConfig config.ApplicationConfig,
	onConfigChoice func(chosenRadarSettings view.RadarSettings, appConfig config.ApplicationConfig),
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
			radarSettings.DistanceType = convertStringToDistanceType[s]
		})
		selectDistanceType.PlaceHolder = ""
		selectDistanceTypeDefaultValue := convertDistanceTypeToString[radarSettings.DistanceType]
		selectDistanceType.SetSelected(selectDistanceTypeDefaultValue)

		selectGeometryType := widget.NewSelect([]string{"Simple", "Vector"}, func(s string) {
			radarSettings.GeometryType = convertStringToGeometryType[s]
		})
		selectGeometryType.PlaceHolder = ""
		selectGeometryTypeDefaultValue := convertGeometryTypeToString[radarSettings.GeometryType]
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
					onConfigChoice(radarSettings, appConfig)
				}
			},
			SubmitText: runText,
			CancelText: resetText,
		}

		return container.NewCenter(form)
	}
}
