package config

import "fyne.io/fyne/v2"

const (
	WindowH = 480
	WindowW = 480
	CatSize = 10
)

type UIConfig struct {
	WindowSize     fyne.Size
	CatSize        fyne.Size
	LightThemeFlag bool
	FullScreenMode bool
	InMainMenu     bool
}

func NewUIConfig() UIConfig {
	return UIConfig{
		WindowSize:     fyne.NewSize(WindowH, WindowW),
		CatSize:        fyne.NewSize(CatSize, CatSize),
		LightThemeFlag: true,
		FullScreenMode: false,
		InMainMenu:     true,
	}
}
