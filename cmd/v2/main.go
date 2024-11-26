package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/PavlushaSource/Radar/view/UI"
	"github.com/PavlushaSource/Radar/view/api"
	config "github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
)

func main() {

	a := app.New()

	w := a.NewWindow("V2")
	config := config.NewApplicationConfig()

	bg := canvas.NewImageFromResource(utils.ResourceBgJpg)
	bg.FillMode = canvas.ImageFillStretch

	//w.SetContent(bg)

	dog := UI.NewDogUI(UI.GetResourceDogSvg(api.Color(1)))
	dog.Resize(fyne.NewSize(100, 100))
	dog.Move(fyne.NewPos(300, 300))

	c := UI.NewDogsContainer(config, bg, []*UI.DogUI{dog})
	c.Resize(config.WindowSize)
	w.SetContent(container.NewStack(bg, c))
	//dog := UI.NewDogUI(UI.GetResourceDogSvg(api.Color(1)))
	//w.SetContent(dog)

	//go func() {
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("SET COLOR NEW")
	//	for i, d := range c.Dogs {
	//		if i%2 == 0 {
	//			d.SetImage(UI.GetResourceDogSvg(api.Color(2)))
	//		}
	//	}
	//}()
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	c.ChangeDogsColor(api.Red)
	//}()

	w.Resize(config.WindowSize)
	//w.SetFullScreen(true)
	//w.SetFixedSize(true)
	w.Show()
	a.Run()
}
