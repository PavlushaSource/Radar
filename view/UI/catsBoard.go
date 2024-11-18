package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/view/utils"
	"image/color"
)

type CatsScrollableContainer struct {
	widget.BaseWidget

	containerSize fyne.Size

	cats          []fyne.CanvasObject
	catsLayout    *CatsLayout
	catsContainer *fyne.Container

	paddingEngineCoord *fyne.Position

	prevEventPosition fyne.Position
}

func (s *CatsScrollableContainer) Scrolled(event *fyne.ScrollEvent) {
	// we only process OY scrolls :)
	scrollDelta := event.Scrolled.DY

	if scrollDelta > 0 {
		s.catsLayout.Scale = utils.IncreaseScale(s.catsContainer, s.containerSize.Width, s.containerSize.Height, s.catsLayout.Scale)
	} else {
		s.catsLayout.Scale = utils.DecreaseScale(s.catsContainer, s.containerSize.Width, s.containerSize.Height, s.catsLayout.Scale)
	}
}

func (s *CatsScrollableContainer) Dragged(event *fyne.DragEvent) {
	//fmt.Println(p)
	if s.prevEventPosition.X == 0 && s.prevEventPosition.Y == 0 {
		s.prevEventPosition = event.Position
		return
	}

	vector := fyne.Position{X: event.Position.X - s.prevEventPosition.X, Y: event.Position.Y - s.prevEventPosition.Y}

	newPadding := s.paddingEngineCoord.Add(vector)
	s.paddingEngineCoord = &newPadding
	//fmt.Println(vector)
	for _, cat := range s.cats {
		go func() {
			cat.Move(cat.Position().Add(vector))
			cat.Refresh()
		}()
	}
	s.prevEventPosition = event.Position
}

func (s *CatsScrollableContainer) DragEnd() {
	s.prevEventPosition = fyne.Position{}
}

func (s *CatsScrollableContainer) CreateRenderer() fyne.WidgetRenderer {
	rect := canvas.NewRectangle(color.Transparent)

	return widget.NewSimpleRenderer(rect)
}

func CreateCatsScrollableContainer(cats []fyne.CanvasObject, catsContainer *fyne.Container, catsLayout *CatsLayout, containerSize fyne.Size) *CatsScrollableContainer {
	board := &CatsScrollableContainer{
		cats:          cats,
		catsContainer: catsContainer,
		catsLayout:    catsLayout,
		containerSize: containerSize,
	}

	board.ExtendBaseWidget(board)

	return board
}
