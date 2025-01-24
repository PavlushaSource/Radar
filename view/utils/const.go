package utils

import (
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Status int

const (
	Fight Status = iota
	Hiss
	Run
)

const (
	DogImgScale     = 0.1
	BorderImgScale  = 0.1
	BorderLineWidth = 3

	Fps = 60

	MaxScale  = 100
	MinScale  = 0.01
	ScaleDiff = 0.25
)

func NewBackground(appConfig *config.ApplicationConfig) *ebiten.Image {
	bg := ebiten.NewImage(appConfig.WindowX, appConfig.WindowY)
	bg.Fill(color.RGBA{R: 102, G: 255, B: 102, A: 0xff})
	return bg
}
