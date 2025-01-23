package view

import (
	"github.com/PavlushaSource/Radar/model/geom"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

type Dog struct {
	Status     *Status
	NextStatus *Status

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

	if d.Status == nil {
		d.Status = newDog.Status
	} else {
		d.Status = d.NextStatus
		d.NextStatus = newDog.Status
	}

	newSpeedX, newSpeedY := calculateDogSpeed(d.X, newDog.X, d.Y, newDog.Y, updateTime)

	d.SpeedX = newSpeedX
	d.SpeedY = newSpeedY
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
	screen.DrawImage(d.StatusToImg(), op)
}

func (d *Dog) StatusToImg() *ebiten.Image {
	switch *d.Status {
	case Run:
		return DogImageRun
	case Fight:
		return DogImageFight
	case Hiss:
		return DogImageHiss
	}
	return nil
}
