package utils

import (
	"bytes"
	"github.com/PavlushaSource/Radar/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

var (
	DogImageFight *ebiten.Image
	DogImageHiss  *ebiten.Image
	DogImageRun   *ebiten.Image
	BorderImage   *ebiten.Image
)

func NewImgFromResource(res []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(res))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func init() {
	DogImageFight = NewImgFromResource(resources.DogFight_png)
	DogImageHiss = NewImgFromResource(resources.DogHiss_png)
	DogImageRun = NewImgFromResource(resources.DogRun_png)
	BorderImage = NewImgFromResource(resources.Border_png)
}
