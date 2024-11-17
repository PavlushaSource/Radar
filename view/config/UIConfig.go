package config

import "fyne.io/fyne/v2"

const (
	WindowH = 480
	WindowW = 480
	CatSize = 100
)

type ApplicationConfig struct {
	WindowSize     fyne.Size
	CatSize        fyne.Size
	LightThemeFlag bool
	FullScreenMode bool
	InMainMenu     bool
}

func NewApplicationConfig() ApplicationConfig {
	return ApplicationConfig{
		WindowSize:     fyne.NewSize(WindowH, WindowW),
		CatSize:        fyne.NewSize(CatSize, CatSize),
		LightThemeFlag: true,
		FullScreenMode: false,
		InMainMenu:     true,
	}
}
