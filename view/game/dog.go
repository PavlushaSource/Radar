package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Status int

const (
	fight Status = iota
	hiss
	run
)

type Dog struct {
	status                   Status
	imageScaleX, imageScaleY float64

	screenWidth, screenHeight int

	x, y int

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

	d.imageScaleX = 0.1
	d.imageScaleY = 0.1

	w, h := int(float64(dogImageFight.Bounds().Dx())*d.imageScaleX), int(float64(dogImageFight.Bounds().Dy())*d.imageScaleY)

	x, y := rand.Int()%1920, rand.Int()%1080
	x = min(max(x, w), 1920-w)
	y = min(max(y, h), 1080-h)

	d.x = x
	d.y = y

	d.speedX, d.speedY = 6, 6
	return d
}

func (d *Dog) Update() {
	w, h := int(float64(dogImageFight.Bounds().Dx())*d.imageScaleX), int(float64(dogImageFight.Bounds().Dy())*d.imageScaleY)

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
	//op.GeoM.Scale(d.imageScaleX, d.imageScaleY)
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
