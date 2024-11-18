package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Loader interface {
	Start()
	Stop()
}

type loader struct {
	size         fyne.Size
	activityText string

	dialog   dialog.Dialog
	activity *widget.Activity
	win      fyne.Window
}

func NewLoader(text string, size fyne.Size, win fyne.Window) Loader {
	return &loader{
		activityText: text,
		size:         size,
		win:          win,
	}
}

func (l *loader) Start() {
	bg := canvas.NewRectangle(color.Transparent)
	bg.SetMinSize(l.size)
	activity := widget.NewActivity()
	d := dialog.NewCustomWithoutButtons(l.activityText, container.NewStack(bg, activity), l.win)

	activity.Start()
	d.Show()

	l.activity = activity
	l.dialog = d

}

func (l *loader) Stop() {
	l.activity.Stop()
	l.dialog.Hide()
}
