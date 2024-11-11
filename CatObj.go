package main

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"time"
)

type Radar struct {
	widget.BaseWidget
	cats       []fyne.CanvasObject
	updateTime time.Duration
	catSize    fyne.Size

	bg *canvas.Rectangle
}

func MoveToCenter(obj fyne.CanvasObject, sizeSquare float32) {
	vector := fyne.Size{Width: -sizeSquare / 2, Height: -sizeSquare / 2}

	fmt.Println("VECTOR", vector)
	obj.Move(obj.Position().Add(vector))
}

func (r *Radar) UpdateRadar() {
	w, h := r.bg.Size().Width, r.bg.Size().Height

	//fmt.Println(w, h)
	for _, cat := range r.cats {
		_, _, _ = w, h, cat
		fmt.Println("Size cat", cat.Size())
		time.Sleep(time.Second * 3)
		MoveToCenter(cat, defaultCatSize)
		time.Sleep(time.Second * 3)
		//X := rand.Float32() * w
		//Y := rand.Float32() * h
		//fmt.Println(X, Y)
		//fmt.Println(cat.Position())
		//cat.Move(fyne.Position{X: X, Y: Y})
	}

}

func (r *Radar) CreateRenderer() fyne.WidgetRenderer {
	return &RadarRender{r}
}

func (r *Radar) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(r.updateTime):
			r.UpdateRadar()
			r.bg.Refresh()
		}
	}
}

func NewRadar() *Radar {
	r := &Radar{}
	r.ExtendBaseWidget(r)

	return r
}

type RadarRender struct {
	r *Radar
}

func (r RadarRender) Destroy() {
}

func (r RadarRender) Layout(size fyne.Size) {
}

func (r RadarRender) MinSize() fyne.Size {
	return r.r.catSize
}

func (r RadarRender) Objects() []fyne.CanvasObject {
	return r.r.cats
}

func (r RadarRender) Refresh() {
}

func (r *Radar) CreateCats(source []CatBackend) []fyne.CanvasObject {
	res := make([]fyne.CanvasObject, 0)

	for _, s := range source {
		//var resource fyne.Resource
		//switch s.color {
		//case Red:
		//	resource = resourceCatSvg
		//case Green:
		//	resource = resourceCatSvg
		//case Blue:
		//	resource = resourceCatSvg
		//default:
		//	panic("Unsupported color")
		//}

		img := canvas.NewImageFromResource(resourceCat(Blue))
		img.Move(fyne.Position{X: s.X, Y: s.Y})
		img.Resize(r.catSize)

		res = append(res, img)
	}

	return res
}

func resourceCat(color Color) fyne.Resource {
	return resourceCatSvg
}
