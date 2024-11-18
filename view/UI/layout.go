package UI

import (
	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/view/config"
)

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
