package main

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"math/rand/v2"
	"time"
)

const defaultCatSize = 100

func show(app fyne.App) {
	//brd := NewBoard(100, 100)

	//game := NewGame(brd)

	window := app.NewWindow("Hello")

	rect := canvas.NewRectangle(color.White)

	r := NewRadar()
	r.bg = rect
	r.catSize = fyne.Size{Width: defaultCatSize, Height: defaultCatSize}
	r.updateTime = time.Second

	//backEnd := GenerateBackendCats(50000)
	backEnd := []CatBackend{{X: 0, Y: 0, color: Red}}
	r.cats = r.CreateCats(backEnd)

	ctx := context.Background()
	go func() {
		time.Sleep(1 * time.Second)
		r.Run(ctx)
	}()

	cnt := container.NewStack(rect, r)
	window.SetContent(cnt)

	window.Resize(fyne.NewSize(480, 480))
	window.Show()
}

func main() {
	myApp := app.New()

	show(myApp)

	myApp.Run()
}

func GenerateBackendCats(count int) []CatBackend {
	res := make([]CatBackend, count)

	minX, maxX := 0, 1920
	minY, maxY := 0, 1080

	for i := 0; i < count; i++ {
		currX := rand.IntN(maxX-minX) + minX
		currY := rand.IntN(maxY-minY) + minY

		res[i] = CatBackend{X: float32(currX), Y: float32(currY)}
		//res = append(res, CatBackend{X: float32(currX), Y: float32(currY)})
	}
	return res
}
