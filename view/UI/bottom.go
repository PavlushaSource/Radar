package UI

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func createLink(text, link string) fyne.CanvasObject {
	parsedUrl, _ := url.Parse(link)
	return widget.NewHyperlink(text, parsedUrl)
}

func CreateBottom() fyne.CanvasObject {
	// TODO: Change to rectangle with link
	texts := container.NewHBox(widget.NewLabel("tepa46, Arsene-Baitenov, PavlushaSoure"),
		layout.NewSpacer(),
		createLink("Source code", "https://github.com/PavlushaSource/Radar"),
	)
	return texts
}
