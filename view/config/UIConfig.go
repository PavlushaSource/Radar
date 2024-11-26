package config

import "fyne.io/fyne/v2"

const (
	WindowW      = 1920
	WindowH      = 1080
	CatSize      = 50
	DefaultScale = 1
)

type ApplicationConfig struct {
	WindowSize     fyne.Size
	CatSize        fyne.Size
	LightThemeFlag bool
	FullScreenMode bool
	InMainMenu     bool

	Scale            float32
	PaddingEnginePos fyne.Position
}

func NewApplicationConfig() *ApplicationConfig {

	padding := fyne.NewPos(WindowW*(1-DefaultScale)/2, WindowH*(1-DefaultScale)/2)

	return &ApplicationConfig{
		WindowSize:       fyne.NewSize(WindowW, WindowH),
		CatSize:          fyne.NewSize(CatSize, CatSize),
		LightThemeFlag:   true,
		FullScreenMode:   false,
		InMainMenu:       true,
		Scale:            DefaultScale,
		PaddingEnginePos: padding,
	}
}
