package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type DogUIRenderer struct {
	dogUI *DogUI

	raster *canvas.Image
}

func (d DogUIRenderer) Destroy() {
}

func (d DogUIRenderer) Layout(size fyne.Size) {
}

func (d DogUIRenderer) MinSize() fyne.Size {
	return d.raster.MinSize()
}

func (d DogUIRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{d.raster}
}

func (d DogUIRenderer) Refresh() {
	d.raster.Resource = d.dogUI.Resource

	d.raster.Refresh()
}

func (d *DogUI) CreateRenderer() fyne.WidgetRenderer {

	img := canvas.NewImageFromResource(d.Resource)
	img.FillMode = canvas.ImageFillContain
	img.Resize(d.MinSize())

	return DogUIRenderer{raster: img, dogUI: d}
}

type DogUI struct {
	widget.BaseWidget

	Resource fyne.Resource
}

func (d *DogUI) SetColor(color fyne.Resource) {
	d.Resource = color
	d.Refresh()
}

func (d *DogUI) MinSize() fyne.Size {
	return fyne.NewSize(50, 50)
}

func NewDogUI(color fyne.Resource) *DogUI {
	dog := &DogUI{Resource: color}
	dog.ExtendBaseWidget(dog)

	return dog
}
