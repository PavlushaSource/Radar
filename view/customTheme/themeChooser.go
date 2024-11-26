package customTheme

import (
	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/view/config"
)

func GetApplicationTheme(config *config.ApplicationConfig) fyne.Theme {
	switch config.LightThemeFlag {
	case true:
		return NewLightTheme()
	default:
		return NewDarkTheme()
	}
}
