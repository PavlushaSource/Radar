package UI

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/PavlushaSource/Radar/view/config"
)

const scaleRatio = 1.1

func CreateContentRunWindow(app RadarApplication, Cats []fyne.CanvasObject) fyne.CanvasObject {
	layout := CatsLayout{Scale: 1, prevSize: app.AppConfig().WindowSize}

	catsContainer := container.New(&layout, Cats...)
	background := CreateCatsBoard(Cats, &layout)
	app.RadarWindow().Canvas().SetOnTypedRune(RegisterScaleRune(app.RadarWindow(), catsContainer, &layout, app.AppConfig()))

	return container.NewStack(background, catsContainer)
}

type CatsLayout struct {
	Cats []fyne.CanvasObject

	ScaleCenter fyne.Position
	prevSize    fyne.Size
	Scale       float32
}

func (d *CatsLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(config.WindowW, config.WindowH)
}

func (d *CatsLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	scaleX := d.prevSize.Width / containerSize.Width
	scaleY := d.prevSize.Height / containerSize.Height

	for _, obj := range objects {

		scaleVectorX := (d.ScaleCenter.X - obj.Position().X) * (scaleX - 1)
		scaleVectorY := (d.ScaleCenter.Y - obj.Position().Y) * (scaleY - 1)
		moveCat := fyne.NewPos(scaleVectorX, scaleVectorY)
		//fmt.Println("moveCat", moveCat)
		//fmt.Println("SCALE CENTER", d.ScaleCenter)
		//currentSize := obj.Size()
		nextSize := fyne.NewSize(config.CatSize*d.Scale, config.CatSize*d.Scale)
		obj.Resize(nextSize)
		obj.Move(obj.Position().Add(moveCat))
		//fmt.Println("Cat position", obj.Position())
	}
	d.prevSize = containerSize
}

func RegisterScaleRune(win fyne.Window, object fyne.CanvasObject, layout *CatsLayout, UIConfig config.ApplicationConfig) func(r rune) {
	return func(r rune) {
		if r == '=' {
			layout.Scale *= scaleRatio
		} else if r == '-' {
			layout.Scale /= scaleRatio
		} else {
			return
		}
		//fmt.Println("Size before resize", object.Size())
		//currSize := object.Size()

		nextSize := fyne.Size{Width: UIConfig.WindowSize.Width * layout.Scale, Height: UIConfig.WindowSize.Height * layout.Scale}

		//fmt.Println("Next size", nextSize.Width, nextSize.Height, "Scale", layout.Scale)

		object.Resize(nextSize)
		//fmt.Println("Size after resize", object.Size())
		//object.Refresh()
		//win.Canvas().Refresh(object)
	}
}

func AnimateCat(A, B fyne.Position, cat fyne.CanvasObject, iterations float32) {
	// TODO: Add scale logic
	vector := fyne.NewPos((B.X-A.X)/iterations, (B.Y-A.Y)/iterations)

	for i := 0; i < int(iterations); i++ {
		time.Sleep(time.Millisecond)
		cat.Move(cat.Position().Add(vector))
	}
}
