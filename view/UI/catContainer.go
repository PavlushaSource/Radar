package UI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/view/api"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
	"math/rand"
	"sync"
	"time"
)

const fps30 = time.Second / 30

func GetResourceCatSvg(color api.Color) fyne.Resource {
	switch color {
	case api.Red:
		return utils.ResourceDogFightingSvg
	case api.Purple:
		return utils.ResourceDogHissingSvg
	case api.Blue:
		return utils.ResourceDogRunningSvg
	default:
		panic("GetResourceCatSvg: unknown color")
	}
}

type CatsContainer struct {
	widget.BaseWidget

	appConfig *config.ApplicationConfig
	Dogs      []*DogUI
	frame     fyne.CanvasObject

	bg                *canvas.Image
	prevEventPosition fyne.Position
	dragVector        fyne.Position
	dragStartTime     time.Time
}

func NewCatsContainer(config *config.ApplicationConfig, bg *canvas.Image) *CatsContainer {
	s := &CatsContainer{appConfig: config, bg: bg}

	s.ExtendBaseWidget(s)

	go s.EmulateCatsCreate()

	return s
}

func (c *CatsContainer) ChangeDogsColor(color api.Color) {
	//for i := range c.Dogs {
	//	cat := canvas.NewImageFromResource(GetResourceCatSvg(color))
	//	cat.Move(c.Dogs[i].Position())
	//	c.Dogs = append(c.Dogs, cat)
	//}
	//canvas.Refresh(c.bg)
}

func (c *CatsContainer) Dragged(event *fyne.DragEvent) {
	if c.prevEventPosition.X == 0 && c.prevEventPosition.Y == 0 {
		c.prevEventPosition = event.Position
		c.dragStartTime = time.Now()
		return
	}
	vector := fyne.Position{X: event.Position.X - c.prevEventPosition.X, Y: event.Position.Y - c.prevEventPosition.Y}
	if fps30 > time.Now().Sub(c.dragStartTime) {
		c.dragVector = vector
		return
	}
	c.appConfig.PaddingEngineCoord = c.appConfig.PaddingEngineCoord.Add(vector)

	wg := sync.WaitGroup{}

	for _, cat := range c.Dogs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newPos := cat.Position().Add(c.dragVector)
			cat.Move(newPos)
		}()
	}
	wg.Wait()
	c.prevEventPosition = event.Position
	c.dragStartTime = time.Now()
	c.dragVector = fyne.Position{}
}

func (c *CatsContainer) DragEnd() {
	c.prevEventPosition = fyne.Position{}
}

func MoveCat(c fyne.CanvasObject) {
	for {
		c.Move(c.Position().Add(fyne.NewPos(2, 0)))
		time.Sleep(fps30)
	}
}

func (c *CatsContainer) EmulateCatsCreate() {

	for i := 0; i < 1000; i++ {
		dog := NewDogUI(GetResourceCatSvg(api.Color(1)))
		dog.Move(fyne.Position{X: float32(rand.Intn(1080)), Y: float32(rand.Intn(1080))})
		dog.Resize(c.appConfig.CatSize)

		//cat := canvas.NewImageFromResource(GetResourceCatSvg(api.Color(1)))
		//cat.Move(fyne.Position{X: float32(rand.Intn(1080)), Y: float32(rand.Intn(1080))})
		//cat.Resize(c.appConfig.CatSize)
		go MoveCat(dog)
		c.Dogs = append(c.Dogs, dog)
	}
	fmt.Println("CREATE DOGS")
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	for i := 0; i < len(c.Dogs); i++ {
	//		cat := canvas.NewImageFromResource(GetResourceCatSvg(api.Color(2)))
	//		cat.Move(fyne.Position{X: float32(rand.Intn(1080)), Y: float32(rand.Intn(1080))})
	//		c.Dogs[i] = cat
	//	}
	//
	//	canvas.Refresh(c.bg)
	//}()
	canvas.Refresh(c.bg)
	fmt.Println("REFRESH")
}

func (c *CatsContainer) CreateRenderer() fyne.WidgetRenderer {
	return &catsContainerRenderer{container: c}
}

type catsContainerRenderer struct {
	container *CatsContainer
}

func (c *catsContainerRenderer) Destroy() {
	// TODO I have no idea what is mean
	fmt.Println("DESTROY CATS CONTAINER")
}

func (c *catsContainerRenderer) Layout(size fyne.Size) {
	//fmt.Println("RESIZE FRAME")
	//fmt.Println("FRAME POS", c.container.frame.Position())
	//fmt.Println("FRAME SIZE", c.container.frame.Size())
	//c.container.frame.Resize(size)
}

func (c *catsContainerRenderer) MinSize() fyne.Size {
	return c.container.appConfig.CatSize
}

func (c *catsContainerRenderer) Objects() []fyne.CanvasObject {
	//objects := []fyne.CanvasObject{c.container.frame}
	//objects = append(objects, c.container.Dogs...)

	res := make([]fyne.CanvasObject, 0, len(c.container.Dogs))
	for i := 0; i < len(c.container.Dogs); i++ {
		res = append(res, c.container.Dogs[i])
	}

	return res
}

func (c *catsContainerRenderer) Refresh() {
	//c.container.frame.Refresh()
}
