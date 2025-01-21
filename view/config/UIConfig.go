package config

import "fyne.io/fyne/v2"

const (
	WindowH = 480
	WindowW = 480
	DogSize = 10
)

type ApplicationConfig struct {
	WindowSize     fyne.Size
	DogSize        fyne.Size
	LightThemeFlag bool
	FullScreenMode bool
	InMainMenu     bool

	ScaleEngineCoord   fyne.Size
	PaddingEngineCoord fyne.Position
}

func NewApplicationConfig() ApplicationConfig {
	return ApplicationConfig{
		WindowSize:         fyne.NewSize(WindowH, WindowW),
		DogSize:            fyne.NewSize(DogSize, DogSize),
		LightThemeFlag:     true,
		FullScreenMode:     false,
		InMainMenu:         true,
		ScaleEngineCoord:   fyne.NewSize(1, 1),
		PaddingEngineCoord: fyne.NewPos(0, 0),
	}
}
