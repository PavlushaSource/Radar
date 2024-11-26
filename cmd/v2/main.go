package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/PavlushaSource/Radar/view/UI"
	"github.com/PavlushaSource/Radar/view/api"
	config "github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
	"image/color"
	"time"
)

func main() {

	a := app.New()

	w := a.NewWindow("V2")
	config := config.NewApplicationConfig()

	rect := canvas.NewRectangle(color.White)
	rect.Resize(config.WindowSize)

	bg := canvas.NewImageFromResource(utils.ResourceBgJpg)
	bg.FillMode = canvas.ImageFillStretch

	//w.SetContent(bg)

	c := UI.NewCatsContainer(config, bg)
	w.SetContent(container.NewStack(bg, c))
	//dog := UI.NewDogUI(UI.GetResourceCatSvg(api.Color(1)))
	//w.SetContent(dog)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("SET COLOR NEW")
		for i, d := range c.Dogs {
			if i%2 == 0 {
				d.SetColor(UI.GetResourceCatSvg(api.Color(2)))
			}
		}
	}()
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	c.ChangeDogsColor(api.Red)
	//}()

	w.Resize(fyne.NewSize(612, 612))
	//w.SetFullScreen(true)
	//w.SetFixedSize(true)
	w.Show()
	a.Run()
}
