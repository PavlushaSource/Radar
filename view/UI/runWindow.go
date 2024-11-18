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

func Create