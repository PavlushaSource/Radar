package runner

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type CatsBoard struct {
	widget.BaseWidget

	obj fyne.Layout

	scaleCenter *fyne.Position
	//background        fyne.CanvasObject
	prevEventPosition fyne.Position
	cats              []fyne.CanvasObject
}

func (c *CatsBoard) Scrolled(event *fyne.ScrollEvent) {

	c.scaleCenter = &event.Position

	//c.background.Resize()
	fmt.Println(event.Scrolled)
	fmt.Println(event.Position)
}

func (c *CatsBoard) Dragged(event *fyne.DragEvent) {

	var p fyne.Position
	fmt.Println(p)
	if c.prevEventPosition.X == 0 && c.prevEventPosition.Y == 0 {
		c.prevEventPosition = event.Position
		return
	}

	vector := fyne.Position{X: event.Position.X - c.prevEventPosition.X, Y: event.Position.Y - c.prevEventPosition.Y}
	fmt.Println(vector)
	for _, cat := range c.cats {
		go func() {
			cat.Move(cat.Position().Add(vector))
			cat.Refresh()
		}()
	}
	c.prevEventPosition = event.Position
}

func (c *CatsBoard) DragEnd() {
	c.prevEventPosition = fyne.Position{}
}

func (c *CatsBoard) CreateRenderer() fyne.WidgetRenderer {
	rect := canvas.NewRectangle(color.Transparent)

	return widget.NewSimpleRenderer(rect)
}

func CreateCatsBoard(cats []fyne.CanvasObject, obj fyne.Layout, scaleCenter *fyne.Position) *CatsBoard {
	board := &CatsBoard{cats: cats, obj: obj, scaleCenter: scaleCenter}

	board.ExtendBaseWidget(board)

	return board
}
