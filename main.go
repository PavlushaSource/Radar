package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")
	w.Resize(fyne.Size{Width: 400, Height: 300})
	w.SetContent(widget.NewLabel("TODOs will go here"))
	w.ShowAndRun()
}
