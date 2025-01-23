package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Status int

type Dog struct {
	status Status

	x, y           int
	xEnd, yEnd     int
	speedX, speedY int
}

func NewDog(image *ebiten.Image) *Dog {
	d := &Dog{}

	i := rand.Intn(2)
	switch i {
	case 0:
		d.status = fight
	case 1:
		d.status = hiss
	}

	w, h := int(float64(dogImageFight.Bounds().Dx())*dogImgScale), int(float64(dogImageFight.Bounds().Dy())*dogImgScale)

	x, y := rand.Int()%1920, rand.Int()%1080
	x = min(max(x, w), 1920-w)
	y = min(max(y, h), 1080-h)

	d.x = x
	d.y = y

	d.speedX, d.speedY = 6, 6
	return d
}

func (d *Dog) Update() {
	w, h := int(float64(dogImageFight.Bounds().Dx())*dogImgScale), int(float64(dogImageFight.Bounds().Dy())*dogImgScale)

	if d.x+d.speedX > 1920-w || d.x+d.speedX < 0 {
		d.speedX = -d.speedX
		if d.status == fight {
			d.status = hiss
		} else {
			d.status = fight
		}
	}

	if d.y+d.speedY > 1080-h || d.y+d.speedY < 0 {
		d.speedY = -d.speedY
		if d.status == fight {
			d.status = hiss
		} else {
			d.status = fight
		}
	}

	d.x += d.speedX
	d.y += d.speedY
}

func (d *Dog) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	//debugMessage := fmt.Sprintf("Screen size %v |  %v.\n", screen.Bounds().Dx(), screen.Bounds().Dy())
	//ebitenutil.DebugPrint(screen, debugMessage)
	//op.GeoM.Scale(d.dogImgScale, d.dogImgScale)
	//op.GeoM.Translate(float64(d.x), float64(d.y))
	//switch d.status {
	//case fight:
	//	screen.DrawImage(dogImageFight, op)
	//case hiss:
	//	screen.DrawImage(dogImageHiss, op)
	//default:
	//	fmt.Println("Dog status not recognized")
	//}
	screen.DrawImage(dogImageFight, op)
}

func (d *Dog) StatusToImg() *ebiten.Image {
	switch d.status {
	case run:

	}
}
