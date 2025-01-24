package config

const (
	WindowW = 1920
	WindowH = 1080
)

type ApplicationConfig struct {
	WindowX, WindowY int
	InMainMenu       bool
	CamX, CamY       float64
	CamScale         float64
	CamScaleTo       float64
}

func NewApplicationConfig() *ApplicationConfig {
	return &ApplicationConfig{
		WindowX:    WindowW,
		WindowY:    WindowH,
		InMainMenu: true,
		CamScale:   1,
		CamScaleTo: 1,
		CamX:       0,
		CamY:       0,
	}
}
