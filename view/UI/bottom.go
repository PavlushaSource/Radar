package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func CreateBottom() fyne.CanvasObject {
	texts := container.NewHBox(widget.NewLabel("tepa46, Arsene-Baitenov, PavlushaSoure"),
		layout.NewSpacer(),
		widget.NewButton("Source code", func() {
			u, _ := url.Parse("https://github.com/PavlushaSource/Radar")
			_ = fyne.CurrentApp().OpenURL(u)
		}))
	return texts
}
