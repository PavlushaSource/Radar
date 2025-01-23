package view

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

	fps = 60

	maxScale  = 100
	minScale  = 0.01
	scaleDiff = 0.25
)

func NewBackground(appConfig *config.ApplicationConfig) *ebiten.Image {
	bg := ebiten.NewImage(appConfig.WindowX, appConfig.WindowY)
	bg.Fill(color.RGBA{R: 51, G: 232, B: 78, A: 0xff})
	return bg
}
