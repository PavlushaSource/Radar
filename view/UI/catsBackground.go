package UI

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type CatsBoard struct {
	widget.BaseWidget

	obj fyne.Layout

	prevEventPosition fyne.Position
	cats              []fyne.CanvasObject
}

// TODO: make opportunity to call scale not only by keyboard
func (c *CatsBoard) Scrolled(event *fyne.ScrollEvent) {
	//fmt.Println(event.Scrolled)
	//fmt.Println(event.Position)
}

func (c *CatsBoard) Dragged(event *fyne.DragEvent) {
	//fmt.Println(p)
	if c.prevEventPosition.X == 0 && c.prevEventPosition.Y == 0 {
		c.prevEventPosition = event.Position
		return
	}

	vector := fyne.Position{X: event.Position.X - c.prevEventPosition.X, Y: event.Position.Y - c.prevEventPosition.Y}
	//fmt.Println(vector)
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

func CreateCatsBoard(cats []fyne.CanvasObject, obj fyne.Layout) *CatsBoard {
	board := &CatsBoard{cats: cats, obj: obj}

	board.ExtendBaseWidget(board)

	return board
}
