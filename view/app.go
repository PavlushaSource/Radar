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
	Borders []utils.Line

	Pause       bool
	BordersDraw bool

	InMainMenu bool

	CancelFunc context.CancelFunc
}

func NewApplication() *Application {
	app := &Application{}
	app.Dogs = make([]*Dog, 0)

	app.AppConfig = config.NewApplicationConfig()
	app.RadarSettings = config.NewRadarSettings()

	app.cursor = utils.NewCursor()
	app.bg = NewBackground(app.AppConfig)
	app.Menu = NewMenu(app)
	app.Borders = make([]utils.Line, 0)

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
		scrollY = -scaleDiff
	} else if ebiten.IsKeyPressed(ebiten.KeyE) || ebiten.IsKeyPressed(ebiten.KeyPageUp) {
		scrollY = scaleDiff
	} else {
		_, scrollY = ebiten.Wheel()
		if scrollY < -scaleDiff*4 {
			scrollY = -scaleDiff * 4
		} else if scrollY > scaleDiff*4 {
			scrollY = scaleDiff * 4
		}
	}
	app.AppConfig.CamScaleTo += scrollY * (app.AppConfig.CamScaleTo / 7)

	// Clamp target zoom level.
	if app.AppConfig.CamScaleTo < minScale {
		app.AppConfig.CamScaleTo = minScale
	} else if app.AppConfig.CamScaleTo > maxScale {
		app.AppConfig.CamScaleTo = maxScale
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
	} else if app.cursor.EndX != 0 && app.cursor.EndY != 0 && app.cursor.StartX != 0 && app.cursor.StartY != 0 {
		app.Borders = append(app.Borders, utils.NewLine(app.cursor.StartX, app.cursor.StartY, app.cursor.EndX, app.cursor.EndY))
		app.cursor.Reset()
	}
}

func (app *Application) Update() error {
	if app.InMainMenu {
		err := app.UpdateMenu()
		return err
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		app.Pause = !app.Pause
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		if app.BordersDraw {
			fmt.Println(app.Borders)
		} else {
			////app.CancelFunc()
			//select {
			//case _, _ = <-app.NeedNext:
			//default:
			//	app.NeedNext <- struct{}{}
			//}

			//viewModel.StartApp(app)
		}

		app.BordersDraw = !app.BordersDraw
	}

	if app.BordersDraw {
		app.UpdateCursor()
		return nil
	}

	app.UpdateDragAndZoom()

	if app.Pause {
		return nil
	}

	flagGetNext := true
	for _, d := range app.Dogs {
		flagGetNext = flagGetNext && d.Update()
	}
	if flagGetNext {
		//app.NeedNext <- struct{}{}
	}
	return nil
}

func (app *Application) DrawLineBorder(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	moveToCenterX := float64(borderImage.Bounds().Dx()) / 2 * BorderImgScale
	moveToCenterY := float64(borderImage.Bounds().Dy()) / 2 * BorderImgScale

	if app.cursor.Pressed {
		op.GeoM.Reset()
		op.GeoM.Translate(-moveToCenterX, -moveToCenterY)
		vector.StrokeLine(screen, float32(app.cursor.StartX), float32(app.cursor.StartY),
			float32(app.cursor.EndX), float32(app.cursor.EndY), BorderLineWidth, color.RGBA{R: 255, A: 64}, false)
	}
}

func (app *Application) DrawBorder(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scale := app.AppConfig.CamScale

	for _, line := range app.Borders {

		dx := line.EndX - line.StartX
		dy := line.EndY - line.StartY
		length := math.Hypot(dx, dy)
		dirX, dirY := dx/length, dy/length
		r := float64(borderImage.Bounds().Dx()) * BorderImgScale / 2
		step := r * 1.5

		for i := 0.0; i < length; i += step {
			op.GeoM.Reset()

			op.GeoM.Translate(float64(-borderImage.Bounds().Dx()/2), float64(-borderImage.Bounds().Dx()/2))
			op.GeoM.Scale(BorderImgScale, BorderImgScale)

			centerX := line.StartX + dirX*i
			centerY := line.StartY + dirY*i
			op.GeoM.Translate(centerX, centerY)

			op.GeoM.Translate(app.AppConfig.CamX, -app.AppConfig.CamY)
			op.GeoM.Scale(scale, scale)

			screen.DrawImage(borderImage, op)
		}
	}
}

func (app *Application) DrawBackground(screen *ebiten.Image) {
	scale := app.AppConfig.CamScale
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(app.AppConfig.CamX, -app.AppConfig.CamY)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(app.bg, op)
}

func (app *Application) Draw(screen *ebiten.Image) {
	if app.InMainMenu {
		app.DrawMenu(screen)
		return
	}

	if app.BordersDraw {
		app.ResetDragAndZoom()
		app.DrawBackground(screen)
		app.DrawLineBorder(screen)
		app.DrawBorder(screen)
		return
	}

	// draw bg
	app.DrawBackground(screen)

	// draw borders
	app.DrawBorder(screen)

	op := &ebiten.DrawImageOptions{}
	scale := app.AppConfig.CamScale
	//target := screen
	//
	//op.GeoM.Translate(app.AppConfig.CamX, -app.AppConfig.CamY)
	//op.GeoM.Scale(scale, scale)
	//target.DrawImage(app.bg, op)

	// draw dogs
	dogImgWidth, dogImgHeight := DogImageRun.Bounds().Dx(), DogImageRun.Bounds().Dy()
	for _, d := range app.Dogs {
		op.GeoM.Reset()

		op.GeoM.Translate(float64(-dogImgWidth/2), float64(-dogImgHeight/2))
		op.GeoM.Scale(DogImgScale, DogImgScale)

		op.GeoM.Translate(d.X, d.Y)
		op.GeoM.Translate(app.AppConfig.CamX, -app.AppConfig.CamY)
		op.GeoM.Scale(scale, scale)

		d.Draw(screen, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("ACTUAL FPS %f\n", ebiten.ActualFPS()))
}

func (app *Application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if app.InMainMenu {
		return outsideWidth, outsideHeight
	} else {
		return app.AppConfig.WindowX, app.AppConfig.WindowY
	}
}
