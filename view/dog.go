package view

import (
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"math/rand"
	"time"
)

type Dog struct {
	Status Status

	X, Y           float64
	XNext, YNext   float64
	SpeedX, SpeedY float64
}

func calculateDogSpeed(x_a, x_b, y_a, y_b float64, updateTime time.Duration) (float64, float64) {
	speedX := (x_b - x_a) / (updateTime.Seconds() * fps)
	speedY := (y_b - y_a) / (updateTime.Seconds() * fps)

	return speedX, speedY
}

func (d *Dog) UpdateDogMove(newDog *Dog, updateTime time.Duration) {
	d.XNext = newDog.X
	d.YNext = newDog.Y

	d.Status = newDog.Status

	newSpeedX, newSpeedY := calculateDogSpeed(d.X, newDog.X, d.Y, newDog.Y, updateTime)

	d.SpeedX = newSpeedX
	d.SpeedY = newSpeedY
}

func NewDog() *Dog {
	d := &Dog{}

	i := rand.Intn(2)
	switch i {
	case 0:
		d.Status = Fight
	case 1:
		d.Status = Hiss
	}

	w, h := int(float64(DogImageFight.Bounds().Dx())*DogImgScale), int(float64(DogImageFight.Bounds().Dy())*DogImgScale)

	x, y := rand.Int()%1920, rand.Int()%1080
	x = min(max(x, w), 1920-w)
	y = min(max(y, h), 1080-h)

	d.X = x
	d.Y = y

	d.SpeedX, d.SpeedY = 6, 6
	return d
}

func (d *Dog) Update() bool {

	if d.X+d.SpeedX <= config.WindowW && d.X+d.SpeedX >= 0 {
		d.X += d.SpeedX
	}

	if d.Y+d.SpeedY <= config.WindowH && d.Y+d.SpeedY >= 0 {
		d.Y += d.SpeedY
	}

	if math.Abs(d.X-d.XNext) < geom.Eps || math.Abs(d.Y-d.YNext) < geom.Eps {
		return true
	}
	return false
}

func (d *Dog) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	//debugMessage := fmt.Sprintf("Screen size %v |  %v.\n", screen.Bounds().Dx(), screen.Bounds().Dy())
	//ebitenutil.DebugPrint(screen, debugMessage)
	//op.GeoM.Scale(d.DogImgScale, d.DogImgScale)
	//op.GeoM.Translate(float64(d.X), float64(d.Y))
	//switch d.Status {
	//case Fight:
	//	screen.DrawImage(DogImageFight, op)
	//case Hiss:
	//	screen.DrawImage(DogImageHiss, op)
	//default:
	//	fmt.Println("Dog Status not recognized")
	//}
	screen.DrawImage(DogImageFight, op)
}

func (d *Dog) StatusToImg() *ebiten.Image {
	switch d.Status {
	case Run:
		return DogImageRun
	case Fight:
		return DogImageFight
	case Hiss:
		return DogImageHiss
	}
	return nil
}
