package UI

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/domain"
	"github.com/PavlushaSource/Radar/view/utils"
	"time"
)

const scaleRatio = 1.1

func CreateContentRunWindow(w1, w2 fyne.Window, backConfig *config.BackendSettings, UIConfig *config.UIConfig, ctx context.Context) fyne.CanvasObject {
	//backEnd := GenerateBackendCats(50000)
	backEnd := []domain.CatBackend{{X: 0, Y: 0, Color: domain.Red}, {X: 100, Y: 100, Color: domain.Purple}, {X: 1000, Y: 100, Color: domain.Blue}}
	Cats := utils.CreateCats(backEnd, UIConfig.CatSize)

	// TODO temp check correct work
	//go func() {
	//	for {
	//		time.Sleep(5 * time.Second)
	//		//fmt.Println("MOVE CATS")
	//		for _, c := range Cats {
	//			go func() {
	//				AnimateCat(c.Position(), c.Position().Add(fyne.NewPos(-100, -100)), c, 600)
	//				time.Sleep(100 * time.Millisecond)
	//			}()
	//		}
	//	}
	//}()

	layout := CatsLayout{BackendConfig: backConfig, UIConfig: UIConfig, Scale: 1, prevSize: UIConfig.WindowSize}

	catsContainer := container.New(&layout, Cats...)
	background := CreateCatsBoard(Cats, &layout)
	w2.Canvas().SetOnTypedRune(RegisterScaleRune(w2, catsContainer, &layout, *UIConfig))

	return container.NewStack(background, catsContainer)
}

type CatsLayout struct {
	BackendConfig *config.BackendSettings
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
	scaleX := containerSize.Width / d.prevSize.Width
	scaleY := containerSize.Height / d.prevSize.Height

	for _, obj := range objects {
		scaleVectorX := (obj.Position().X - d.ScaleCenter.X) * scaleX
		scaleVectorY := (obj.Position().Y - d.ScaleCenter.Y) * scaleY
		moveCat := fyne.NewPos(scaleVectorX, scaleVectorY)
		nextSize := fyne.NewSize(d.UIConfig.CatSize.Width*d.Scale, d.UIConfig.CatSize.Height*d.Scale)
		obj.Resize(nextSize)
		obj.Move(moveCat)
	}
	d.prevSize = containerSize
}

func RegisterScaleRune(w2 fyne.Window, object fyne.CanvasObject, layout *CatsLayout, UIConfig config.UIConfig) func(r rune) {
	return func(r rune) {
		if r == '=' {
			layout.Scale *= scaleRatio
		} else if r == '-' {
			layout.Scale /= scaleRatio
		} else {
			return
		}
		//fmt.Println("Size before resize", object.Size())
		//win.Canvas().Size().
		nextSize := fyne.Size{Width: w2.Canvas().Size().Width * layout.Scale, Height: w2.Canvas().Size().Height * layout.Scale}

		//fmt.Println("Next size", nextSize.Width, nextSize.Height, "Scale", layout.Scale)

		object.Resize(nextSize)
		//fmt.Println("Size after resize", object.Size())
		//object.Refresh()
		//win.Canvas().Refresh(object)
	}
}

func AnimateCat(A, B fyne.Position, cat fyne.CanvasObject, iterations float32) {
	// TODO: Add scale logic
	vector := fyne.NewPos((A.X-B.X)/iterations, (A.Y-B.Y)/iterations)

	for i := 0; i < int(iterations); i++ {
		time.Sleep(time.Millisecond)
		cat.Move(cat.Position().Add(vector))
	}

}
