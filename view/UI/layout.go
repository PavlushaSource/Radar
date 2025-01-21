package UI

import (
	"fmt"
	"sync"

	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/view/config"
)

type CatsLayout struct {
	Cats []fyne.CanvasObject

	Border fyne.CanvasObject

	ScaleCenter fyne.Position
	PrevSize    fyne.Size
	AppConfig   *config.ApplicationConfig
	Scale       float32
}

func (d *CatsLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(config.WindowW, config.WindowH)
}

func (d *CatsLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {

	scaleX := d.PrevSize.Width / containerSize.Width
	scaleY := d.PrevSize.Height / containerSize.Height
	fmt.Println("SCALE NOW", d.Scale)
	fmt.Println("Border position", d.Border.Position())

	wg := sync.WaitGroup{}
	for _, obj := range objects[1:] {
		wg.Add(1)
		go func() {
			defer wg.Done()
			scaleVectorX := (d.Border.Position().X - obj.Position().X) * (scaleX - 1)
			scaleVectorY := (d.Border.Position().Y - obj.Position().Y) * (scaleY - 1)
			moveCat := fyne.NewPos(scaleVectorX, scaleVectorY)
			//fmt.Println("moveCat", moveCat)
			//fmt.Println("SCALE CENTER", d.ScaleCenter)
			//currentSize := obj.Size()
			nextSize := fyne.NewSize(config.DogSize*d.Scale, config.DogSize*d.Scale)
			obj.Resize(nextSize)
			obj.Move(obj.Position().Add(moveCat))
			//fmt.Println("Cat position", obj.Position())
		}()
	}
	wg.Wait()
	// border scale
	d.Border.Resize(containerSize)
	//d.Border.Move(d.Border.Position().Add(moveBorder))

	d.AppConfig.ScaleEngineCoord = fyne.NewSize(d.AppConfig.ScaleEngineCoord.Width/scaleX, d.AppConfig.ScaleEngineCoord.Height/scaleY)
	fmt.Println("New scaleEngineCoord", d.AppConfig.ScaleEngineCoord, "Container new size", containerSize)
	d.PrevSize = containerSize

	//scaleVectorX := (d.ScaleCenter.X - d.Border.Position().X) * (scaleX - 1)
	//scaleVectorY := (d.ScaleCenter.Y - d.Border.Position().Y) * (scaleY - 1)
	//moveBorder := fyne.NewPos(scaleVectorX, scaleVectorY)
	//d.Border.Resize(fyne.NewSize(containerSize.Width*d.Scale, containerSize.Height*d.Scale))

	//d.Border.Resize(fyne.NewSize(
	//	(containerSize.Width/d.Border.Size().Width)*d.Scale,
	//	(containerSize.Height/d.Border.Size().Height)*d.Scale))

	//fmt.Println("SCALE X", scaleX, "PREV SIZE", d.PrevSize, "Curr size", containerSize)
	//newScaleEngineCoord := fyne.NewSize(d.AppConfig.ScaleEngineCoord.Width/scaleX*d.Scale, d.AppConfig.ScaleEngineCoord.Height/scaleY*d.Scale)
	//
	//fmt.Println("X coef border", newScaleEngineCoord.Width/d.AppConfig.ScaleEngineCoord.Width)
	//d.Border.Resize(fyne.NewSize(d.Border.Size().Width*(1/scaleX)*d.Scale, d.Border.Size().Height*(1/scaleY)*d.Scale))
	////d.Border.Resize(containerSize)
	//fmt.Println("SIZE BORDER", d.Border.Size())
	////d.Border.Move(d.Border.Position().Add(moveBorder))
	//
	//d.AppConfig.ScaleEngineCoord = newScaleEngineCoord
	//fmt.Println("New scaleEngineCoord", d.AppConfig.ScaleEngineCoord, "Container new size", containerSize)
	//d.PrevSize = d.Border.Size()
	fmt.Println("______________________________________________________________")
}
