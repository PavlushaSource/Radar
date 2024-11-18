package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func createLink(text, link string) fyne.CanvasObject {
	parsedUrl, _ := url.Parse(link)
	return widget.NewHyperlink(text, parsedUrl)
}

func createBottom() fyne.CanvasObject {
	// TODO: Change to rectangle with link
	texts := container.NewHBox(widget.NewLabel("tepa46, Arsene-Baitenov, PavlushaSoure"),
		layout.NewSpacer(),
		createLink("Source code", "https://github.com/PavlushaSource/Radar"),
	)
	return texts
}
