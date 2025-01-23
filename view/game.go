package view

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PavlushaSource/Radar/resources"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	"log"
	"math"
)

var (
	DogImageFight *ebiten.Image
	DogImageHiss  *ebiten.Image
	DogImageRun   *ebiten.Image
	borderImage   *ebiten.Image
)

func NewImgFromResource(res []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(res))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func init() {
	DogImageFight = NewImgFromResource(resources.DogFight_png)
	DogImageHiss = NewImgFromResource(resources.DogHiss_png)
	DogImageRun = NewImgFromResource(resources.DogRun_png)
	borderImage = NewImgFromResource(resources.Border_png)
}

type Application struct {
	Dogs     []*Dog
	NeedNext chan struct{}

	AppConfig     *config.ApplicationConfig
	RadarSettings *config.RadarSettings

	cursor  utils.Cursor
	bg      *ebiten.Image
	Menu    *ebitenui.UI
	borders []utils.Line

	Pause       bool
	BordersDraw bool

	InMainMenu bool

	cancelFunc context.CancelFunc
}

func NewApplication() *Application {
	app := &Application{}
	app.Dogs = make([]*Dog, 0)

	app.AppConfig = config.NewApplicationConfig()
	app.RadarSettings = config.NewRadarSettings()

	app.cursor = utils.NewCursor()
	app.bg = NewBackground(app.AppConfig)
	app.Menu = NewMenu(app)
	app.borders = make([]utils.Line, 0)

	app.InMainMenu = true

	return app
}

func (app *Application) ResetDragAndZoom() {
	app.AppConfig.CamScaleTo = 1
	app.AppConfig.CamScale = 1
	app.AppConfig.CamX = 0
	app.AppConfig.CamY = 0
}

func (app *Application) UpdateDragAndZoom() {
	var scrollY float64
	if ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyPageDown) {
		scrollY = -0.25
	} else if ebiten.IsKeyPressed(ebiten.KeyE) || ebiten.IsKeyPressed(ebiten.KeyPageUp) {
		scrollY = 0.25
	} else {
		_, scrollY = ebiten.Wheel()
		if scrollY < -1 {
			scrollY = -1
		} else if scrollY > 1 {
			scrollY = 1
		}
	}
	app.AppConfig.CamScaleTo += scrollY * (app.AppConfig.CamScaleTo / 7)

	// Clamp target zoom level.
	if app.AppConfig.CamScaleTo < 0.01 {
		app.AppConfig.CamScaleTo = 0.01
	} else if app.AppConfig.CamScaleTo > 100 {
		app.AppConfig.CamScaleTo = 100
	}

	// Smooth zoom transition.
	div := 10.0
	if app.AppConfig.CamScaleTo > app.AppConfig.CamScale {
		app.AppConfig.CamScale += (app.AppConfig.CamScaleTo - app.AppConfig.CamScale) / div
	} else if app.AppConfig.CamScaleTo < app.AppConfig.CamScale {
		app.AppConfig.CamScale -= (app.AppConfig.CamScale - app.AppConfig.CamScaleTo) / div
	}

	// Pan camera via keyboard.
	pan := 7.0 / app.AppConfig.CamScale
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		app.AppConfig.CamX -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		app.AppConfig.CamX += pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		app.AppConfig.CamY -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		app.AppConfig.CamY += pan
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		app.ResetDragAndZoom()
	}
}

func (app *Application) UpdateCursor() {
	var mx, my int = ebiten.CursorPosition()
	mx -= int(app.AppConfig.CamX * app.AppConfig.CamScale)
	my += int(app.AppConfig.CamY * app.AppConfig.CamScale)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !app.cursor.Pressed {
			app.cursor.Pressed = true
			app.cursor.StartX, app.cursor.StartY = mx, my
		}
		app.cursor.EndX, app.cursor.EndY = mx, my
	} else {
		app.cursor.Pressed = false
	}
}

func (app *Application) Update() error {
	if app.InMainMenu {
		err := app.UpdateMenu()
		return err
	}

	app.UpdateDragAndZoom()
	app.UpdateCursor()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		app.Pause = !app.Pause
	}

	if app.Pause {
		return nil
	}

	flagGetNext := true
	for _, d := range app.Dogs {
		flagGetNext = flagGetNext && d.Update()
	}
	if flagGetNext {
		app.NeedNext <- struct{}{}
	}
	return nil
}

func (app *Application) Draw(screen *ebiten.Image) {
	if app.InMainMenu {
		app.DrawMenu(screen)
		return
	}

	//screen.Fill(color.RGBA{0xff, 0, 0, 0})
	//cx, cy := ScreenWidth/2, ScreenHeight/2
	target := screen
	scale := app.AppConfig.CamScale
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(app.AppConfig.CamX, -app.AppConfig.CamY)
	op.GeoM.Scale(scale, scale)
	//op.GeoM.Translate(float64(cx), float64(cy))
	target.DrawImage(app.bg, op)

	if app.cursor.Pressed {
		op.GeoM.Reset()
		op.GeoM.Translate(-160, -160)
		op.GeoM.Scale(0.1, 0.1)

		//fmt.Println(scale)
		//op.GeoM.Translate(-app.camX/scale, app.camY/scale)
		//op.GeoM.Scale(scale, scale)

		//fmt.Println(app.cursor.startY, app.cursor.pressed)
		vector.StrokeLine(target, float32(app.cursor.StartX), float32(app.cursor.StartY), float32(app.cursor.EndX), float32(app.cursor.EndY), 3, color.RGBA{255, 0, 0, 64}, false)

		//app.bg.DrawImage(borderImage, op)
	} else {
		app.borders = append(app.borders, utils.NewLine(app.cursor.StartX, app.cursor.StartY, app.cursor.EndX, app.cursor.EndY))

	}

	app.DrawBorder(target)

	for _, d := range app.Dogs {

		op.GeoM.Reset()

		op.GeoM.Scale(DogImgScale, DogImgScale)
		op.GeoM.Translate(float64(d.X), float64(d.Y))
		op.GeoM.Translate(app.AppConfig.CamX, -app.AppConfig.CamY)

		op.GeoM.Scale(scale, scale)

		//op.GeoM.Translate(float64(cx), float64(cy))
		//screen.DrawImage(DogImageFight, op)
		d.Draw(target, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("ACTUAL FPS %f\n", ebiten.ActualFPS()))

	//screen.DrawImage(target, op)

}

func (app *Application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return app.AppConfig.WindowX, app.AppConfig.WindowY
}

func (app *Application) DrawBorder(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scale := app.AppConfig.CamScale

	for _, line := range app.borders {

		dx := float64(line.EndX - line.StartX)
		dy := float64(line.EndY - line.StartY)

		length := math.Hypot(dx, dy)

		dirX, dirY := dx/length, dy/length

		r := 16.0
		step := r * 1.5

		for i := 0.0; i < length; i += step {
			op.GeoM.Reset()

			op.GeoM.Translate(-160, -160)
			op.GeoM.Scale(borderImgScale, borderImgScale)

			centerX := float64(line.StartX) + dirX*i
			centerY := float64(line.StartY) + dirY*i

			op.GeoM.Translate(centerX, centerY)

			op.GeoM.Translate(app.AppConfig.CamX, -app.AppConfig.CamY)
			op.GeoM.Scale(scale, scale)

			screen.DrawImage(borderImage, op)
			//app.bg.DrawImage(borderImage, op)
			//vector.DrawFilledCircle(screen, float32(centerX), float32(centerY), float32(r), color.RGBA{0, 255, 0, 255}, false)
		}
		//currPosX := line.StartX
		//currPosY := line.StartY
		//op.GeoM.Reset()
		//op.GeoM.Translate(-160, -160)
		//op.GeoM.Scale(0.1, 0.1)
		//
		//op.GeoM.Translate(float64(currPosX)/scale, float64(currPosY)/scale)
		//
		//op.GeoM.Translate(-app.camX/scale, app.camY/scale)
		//op.GeoM.Scale(scale, scale)
	}

}
