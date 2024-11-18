package UI

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
)

type CatsScrollableContainer struct {
	widget.BaseWidget

	containerSize fyne.Size
	cats          []fyne.CanvasObject
	catsLayout    *CatsLayout
	catsContainer *fyne.Container
	AppConfig     *config.ApplicationConfig

	prevEventPosition fyne.Position
}

func (s *CatsScrollableContainer) Scrolled(event *fyne.ScrollEvent) {
	// we only process OY scrolls :)
	scrollDelta := event.Scrolled.DY

	if scrollDelta > 0 {
		fmt.Println("Increase Scale", s.catsLayout.Scale)
		s.catsLayout.Scale = utils.IncreaseScale(
			s.catsContainer,
			s.catsContainer.Size().Width,
			s.catsContainer.Size().Height,
			s.catsLayout.Scale)
	} else {
		fmt.Println("Decrease Scale", s.catsLayout.Scale)
		s.catsLayout.Scale = utils.DecreaseScale(
			s.catsContainer,
			s.catsContainer.Size().Width,
			s.catsContainer.Size().Height,
			s.catsLayout.Scale)
	}
}

func (s *CatsScrollableContainer) Dragged(event *fyne.DragEvent) {
	//fmt.Println(p)
	if s.prevEventPosition.X == 0 && s.prevEventPosition.Y == 0 {
		s.prevEventPosition = event.Position
		return
	}

	vector := fyne.Position{X: event.Position.X - s.prevEventPosition.X, Y: event.Position.Y - s.prevEventPosition.Y}

	s.AppConfig.PaddingEngineCoord = s.AppConfig.PaddingEngineCoord.Add(vector)
	//fmt.Println(vector)
	for _, cat := range s.cats {
		go func() {
			newPos := cat.Position().Add(vector)
			cat.Move(newPos)

			//cat.Refresh()
		}()
	}
	s.catsLayout.Border.Move(s.catsLayout.Border.Position().Add(vector))
	s.prevEventPosition = event.Position
}

func (s *CatsScrollableContainer) DragEnd() {
	s.prevEventPosition = fyne.Position{}
}

func (s *CatsScrollableContainer) CreateRenderer() fyne.WidgetRenderer {
	rect := canvas.NewRectangle(color.Transparent)
	//rect.StrokeColor = color.Black
	//rect.StrokeWidth = 3

	return widget.NewSimpleRenderer(rect)
}

func CreateCatsScrollableContainer(cats []fyne.CanvasObject, catsContainer *fyne.Container, catsLayout *CatsLayout, containerSize fyne.Size, appConfig *config.ApplicationConfig) *CatsScrollableContainer {
	board := &CatsScrollableContainer{
		cats:          cats,
		catsContainer: catsContainer,
		catsLayout:    catsLayout,
		containerSize: containerSize,
		AppConfig:     appConfig,
	}

	board.ExtendBaseWidget(board)

	return board
}
