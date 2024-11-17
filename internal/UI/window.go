package UI

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/PavlushaSource/Radar/internal/config"
	"github.com/PavlushaSource/Radar/internal/runner"
	"image/color"
	"time"
)

func CreateContentRunWindow(w1, w2 fyne.Window, backConf *config.BackendConfig, frontConf *config.UIConfig, ctx context.Context) fyne.CanvasObject {
	rect := canvas.NewRectangle(color.Transparent)

	r := runner.NewRadar()
	r.Bg = rect
	r.CatSize = frontConf.CatSize
	r.UpdateTime = backConf.UpdateTime

	//backEnd := GenerateBackendCats(50000)
	backEnd := []runner.CatBackend{{X: 0, Y: 0, Color: runner.Red}, {X: 100, Y: 100, Color: runner.Purple}, {X: 1000, Y: 100, Color: runner.Blue}}
	r.Cats = r.CreateCats(backEnd)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			fmt.Println("MOVE CATS")
			for _, c := range r.Cats {
				go func() {
					AnimateCat(c.Position(), c.Position().Add(fyne.NewPos(-100, -100)), c, 600)
					time.Sleep(100 * time.Millisecond)
				}()

				//move := canvas.NewPositionAnimation(c.Position(), c.Position().Add(fyne.NewPos(200, 200)), time.Second*3, c.Move)
				//move.Start()
				//time.Sleep(500 * time.Millisecond)
				//c.Move(c.Position().Add(fyne.NewPos(100, 0)))
			}
		}
	}()

	layout := CatsLayout{BackendConfig: backConf, UIConfig: frontConf, Scale: 1, prevSize: frontConf.WindowSize}
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	r.Run(ctx)
	//}()

	//text1 := widget.NewLabel("topleft")
	//text2 := widget.NewLabel("Middle Label")
	//text3 := widget.NewLabel("bottomright")

	//cnt := container.NewStack(rect, r)

	catsContainer := container.New(&layout, r.Cats...)
	background := runner.CreateCatsBoard(r.Cats, &layout, &layout.ScaleCenter)
	w2.Canvas().SetOnTypedRune(RegisterScaleRune(w2, catsContainer, &layout))

	return container.NewStack(background, catsContainer)
	//return runner.CreateCatsBoard(&layout, r.Cats)
	//return runner.NewCustomContainer(&layout, r.Cats)
}

type CatsLayout struct {
	BackendConfig *config.BackendConfig
	UIConfig      *config.UIConfig
	Cats          []fyne.CanvasObject

	ScaleCenter fyne.Position
	prevSize    fyne.Size
	Scale       float32
}

func (d *CatsLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return d.UIConfig.WindowSize
}

func (d *CatsLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	scaleX := d.prevSize.Width / containerSize.Width
	scaleY := d.prevSize.Height / containerSize.Height

	for _, obj := range objects {

		scaleVectorX := (d.ScaleCenter.X - obj.Position().X) * (scaleX - 1)
		scaleVectorY := d.ScaleCenter.Y - obj.Position().Y*(scaleY-1)
		moveCat := fyne.NewPos(scaleVectorX, scaleVectorY)
		//fmt.Println("moveCat", moveCat)
		fmt.Println("SCALE CENTER", d.ScaleCenter)
		//currentSize := obj.Size()
		nextSize := fyne.NewSize(d.UIConfig.CatSize.Width*d.Scale, d.UIConfig.CatSize.Height*d.Scale)
		obj.Resize(nextSize)
		obj.Move(obj.Position().Add(moveCat))
		//fmt.Println("Cat position", obj.Position())
	}
	d.prevSize = containerSize
}

func RegisterScaleRune(win fyne.Window, object fyne.CanvasObject, layout *CatsLayout) func(r rune) {
	return func(r rune) {
		if r == '=' {
			layout.Scale *= 1.1
		} else if r == '-' {
			layout.Scale /= 1.1
		} else {
			return
		}
		fmt.Println("Size before resize", object.Size())
		//currSize := object.Size()
		nextSize := fyne.Size{Width: 1920 * layout.Scale, Height: 1080 * layout.Scale}

		fmt.Println("Next size", nextSize.Width, nextSize.Height, "Scale", layout.Scale)

		object.Resize(nextSize)
		fmt.Println("Size after resize", object.Size())
		//object.Refresh()
		//win.Canvas().Refresh(object)
	}
}

//type DragWrap struct {
//	widget.BaseWidget
//
//	container fyne.CanvasObject
//}
//
//func (w *DragWrap) Dragged(event *fyne.DragEvent) {
//	fmt.Println("Hello")
//}
//
//func (w *DragWrap) DragEnd() {
//
//}

func AnimateCat(A, B fyne.Position, cat fyne.CanvasObject, iterations float32) {
	vector := fyne.NewPos((A.X-B.X)/iterations, (A.Y-B.Y)/iterations)

	for i := 0; i < int(iterations); i++ {
		time.Sleep(time.Millisecond)
		cat.Move(cat.Position().Add(vector))
	}

}
