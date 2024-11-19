package UI

import (
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"sync"

	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/view/config"
)

type CatsLayout struct {
	Cats []fyne.CanvasObject

	Border fyne.CanvasObject
	Lines  []*canvas.Line

	ScaleCenter fyne.Position
	PrevSize    fyne.Size
	AppConfig   *config.ApplicationConfig
	Scale       float32

	prevLineWidth float32
	BorderCount   int
}

func (d *CatsLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(config.WindowW, config.WindowH)
}

func (d *CatsLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {

	scaleX := d.PrevSize.Width / containerSize.Width
	scaleY := d.PrevSize.Height / containerSize.Height
	fmt.Println("SCALE NOW", d.Scale)
	fmt.Println("Border position", d.Border.Position())

	//line := objects[2 : 2+d.BorderCount+1]

	wg := sync.WaitGroup{}
	for i, obj := range objects[1:] {
		wg.Add(1)
		go func() {
			defer wg.Done()
			scaleVectorX := (d.Border.Position().X - obj.Position().X) * (scaleX - 1)
			scaleVectorY := (d.Border.Position().Y - obj.Position().Y) * (scaleY - 1)
			moveCat := fyne.NewPos(scaleVectorX, scaleVectorY)
			//fmt.Println("moveCat", moveCat)
			//fmt.Println("SCALE CENTER", d.ScaleCenter)
			//currentSize := obj.Size()
			nextSize := fyne.NewSize(config.CatSize*d.Scale, config.CatSize*d.Scale)
			if i >= 0 && i < d.BorderCount {
				fmt.Println("SKIP LINE")
				//l := obj.(*canvas.Line)
				//l.StrokeWidth = 3 * d.Scale / l.StrokeWidth
				//obj.Resize(fyne.NewSize(100, 100))
				//fmt.Println("BORDER", obj.Position(), "SIZE", obj.Size())
			} else {
				obj.Resize(nextSize)
				obj.Move(obj.Position().Add(moveCat))
			}
			//fmt.Println("Cat position", obj.Position())
		}()
	}
	for _, l := range d.Lines {
		//scaleVectorX := (d.Border.Position().X - l.Position().X) * (scaleX - 1)

		fmt.Println("SCALE STROKEWIDTH", d.Scale/l.StrokeWidth, "DSCALE", d.Scale, "STROKEWITH PREV", l.StrokeWidth)

		fmt.Println("NEW CENTER X BORDER", containerSize.Width/2)
		l.Position1 = fyne.NewPos(containerSize.Width/2+d.AppConfig.PaddingEngineCoord.X, d.AppConfig.PaddingEngineCoord.Y)
		l.Position2 = fyne.NewPos(containerSize.Width/2+d.AppConfig.PaddingEngineCoord.X, containerSize.Height+d.AppConfig.PaddingEngineCoord.Y)

		fmt.Println("CURRENT POSITION BORDER AFTER RESIZE", l.Position1.X)
		//fmt.Println("NEW X POS LINE", scaleVectorX*containerSize.Width/2)
		if d.prevLineWidth == 0 {
			d.prevLineWidth = 3
		}
		l.StrokeWidth = (1 / scaleX) * d.prevLineWidth
		l.Refresh()
		d.prevLineWidth = l.StrokeWidth
		//l.Move(l.Position().Add(fyne.NewPos(-100, 0)))
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
