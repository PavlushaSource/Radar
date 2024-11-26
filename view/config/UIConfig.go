package config

import "fyne.io/fyne/v2"

const (
	WindowW = 1920
	WindowH = 1080
	CatSize = 50
)

type ApplicationConfig struct {
	WindowSize     fyne.Size
	CatSize        fyne.Size
	LightThemeFlag bool
	FullScreenMode bool
	InMainMenu     bool

	ScaleEngineCoord   fyne.Size
	PaddingEngineCoord fyne.Position
}

func NewApplicationConfig() *ApplicationConfig {
	scale := fyne.NewSize(1, 1)

	padding := fyne.NewPos(WindowW*(1-scale.Width)/2, WindowH*(1-scale.Width)/2)

	return &ApplicationConfig{
		WindowSize:         fyne.NewSize(WindowW, WindowH),
		CatSize:            fyne.NewSize(CatSize, CatSize),
		LightThemeFlag:     true,
		FullScreenMode:     false,
		InMainMenu:         true,
		ScaleEngineCoord:   scale,
		PaddingEngineCoord: padding,
	}
}
