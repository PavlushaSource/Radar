package main

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/custom_theme"
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
	backEnd := []CatBackend{{X: 0, Y: 0, color: Red}, {X: 10, Y: 10, color: Green}}
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

//func main() {
//	myApp := app.New()
//
//	show(myApp)
//
//	myApp.Run()
//}

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

//func main() {
//	a := app.New()
//	w := a.NewWindow("Hello World")
//
//	w.SetContent(widget.NewLabel("Hello World!"))
//	w.Show()
//	go func() {
//		time.Sleep(1 * time.Second)
//		w.Hide()
//	}()
//	w2 := a.NewWindow("Larger")
//	w2.SetContent(widget.NewLabel("More content"))
//	w2.Resize(fyne.NewSize(100, 100))
//	w2.Show()
//
//	a.Run()
//}

//func main() {
//	myApp := app.New()
//	myWindow := myApp.NewWindow("Center Layout")
//
//	img := canvas.NewImageFromResource(theme.FyneLogo())
//	img.FillMode = canvas.ImageFillOriginal
//	text := canvas.NewText("Overlay", color.Black)
//	content := container.New(layout.NewCenterLayout(), img, text)
//
//	myWindow.SetContent(content)
//	myWindow.ShowAndRun()
//}

func main() {
	App := app.New()
	App.Settings().SetTheme(custom_theme.NewLightTheme())

	//BuildUI(App)

	myWindow := App.NewWindow("Form Layout")

	secondWindow := App.NewWindow("HIHI")
	secondWindow.SetContent(container.NewHBox(widget.NewLabel("HELLO MAN")))
	//secondWindow.SetFullScreen(true)
	//secondWindow.Resize(fyne.NewSize(480, 480))

	myWindow.Resize(fyne.NewSize(480, 480))
	myWindow.SetContent(BuildMainMenu(App, myWindow, secondWindow))

	myWindow.Show()
	App.Run()
}

func BuildMainMenu(app fyne.App, firstWindow, secondWindow fyne.Window) fyne.CanvasObject {

	calcType := widget.NewSelect([]string{"Option 1", "Option 2", "Option 3"}, func(s string) {
		switch s {
		case "Option 1":
			firstWindow.SetFullScreen(true)
		}

		fmt.Println("selected", s)
	})
	calcType.PlaceHolder = "Выберите одно из предложенных"

	top := widget.NewLabel("BLYADINA PRESENT")
	top.Alignment = fyne.TextAlignTrailing

	LightThemeFlag := false
	FullScreenMode := false
	topToolbar := widget.Toolbar{
		Items: []widget.ToolbarItem{widget.NewToolbarAction(theme.HomeIcon(), func() {
		}), widget.NewToolbarAction(theme.ColorPaletteIcon(), func() {
			if !LightThemeFlag {
				app.Settings().SetTheme(custom_theme.NewDarkTheme())
			} else {
				app.Settings().SetTheme(custom_theme.NewLightTheme())
			}
			LightThemeFlag = !LightThemeFlag
		}), widget.NewToolbarAction(theme.ViewFullScreenIcon(), func() {
			if !FullScreenMode {
				firstWindow.SetFullScreen(true)
			} else {
				firstWindow.SetFullScreen(false)
			}
			FullScreenMode = !FullScreenMode
		})},
	}

	cfgUI := &widget.Form{}
	cfgUI.Append("Select calc", calcType)
	bg := canvas.NewRectangle(&color.NRGBA{R: 0, G: 0, B: 0x4d, A: 0xff})
	_ = bg

	startButton := widget.NewButtonWithIcon("Run", resourceCatIconWhiteSvg, func() {
		firstWindow.Hide()
		secondWindow.Show()

		secondWindow.SetMaster()
	})

	//img := canvas.NewImageFromResource(resourceIconWhite())
	//img.Resize(fyne.NewSize(200, 200))

	SelectTheme := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			fmt.Println("Dark")
			app.Settings().SetTheme(custom_theme.NewDarkTheme())
		}),
		widget.NewButton("Light", func() {
			fmt.Println("Light")
			app.Settings().SetTheme(custom_theme.NewLightTheme())
		}))

	return container.NewBorder(container.NewVBox(&topToolbar, top), SelectTheme, nil, nil, container.NewCenter(container.NewVBox(cfgUI, startButton)))
}
