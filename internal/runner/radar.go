package runner

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
	Cats       []fyne.CanvasObject
	UpdateTime time.Duration
	CatSize    fyne.Size

	window fyne.Window
	Bg     *canvas.Rectangle
}

func MoveToCenter(obj fyne.CanvasObject, sizeSquare float32) {
	vector := fyne.Size{Width: -sizeSquare * 1000 / 2, Height: -sizeSquare * 1000 / 2}

	fmt.Println("VECTOR", vector)
	obj.Move(obj.Position().Add(vector))
}

func (r *Radar) UpdateRadar() {
	w, h := r.Bg.Size().Width, r.Bg.Size().Height

	//fmt.Println(w, h)
	for _, cat := range r.Cats {
		_, _, _ = w, h, cat

		//fmt.Println("Size cat", cat.Size())
		//time.Sleep(time.Second * 3)
		//MoveToCenter(cat, defaultCatSize)
		//time.Sleep(time.Second * 3)
		//fyne.Scrollable()
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
		case <-time.After(r.UpdateTime):
			r.UpdateRadar()

			r.Bg.Refresh()
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
	return r.r.Bg.MinSize()
}

func (r RadarRender) Objects() []fyne.CanvasObject {
	return r.r.Cats
}

func (r RadarRender) Refresh() {
	r.r.Refresh()
}

func (r *Radar) CreateCats(source []CatBackend) []fyne.CanvasObject {
	res := make([]fyne.CanvasObject, 0)

	for _, s := range source {
		img := canvas.NewImageFromResource(resourceCat(s.Color))
		img.Move(fyne.Position{X: s.X, Y: s.Y})
		img.Resize(r.CatSize)

		res = append(res, img)
	}

	return res
}

func resourceCat(color Color) fyne.Resource {
	switch color {
	case Red:
		return ResourceCatRedSvg
	case Purple:
		return ResourceCatPurpleSvg
	case Blue:
		return ResourceCatBlueSvg
	default:
		panic("resourceCat: unknown color")
	}
}
